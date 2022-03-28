package main

import (
	"compress/flate"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dalot/go-skeleton-mid/cmd/app/config"
	"github.com/dalot/go-skeleton-mid/internal/contracts"
	"github.com/dalot/go-skeleton-mid/internal/database"
	"github.com/dalot/go-skeleton-mid/internal/http/adminapi"
	"github.com/dalot/go-skeleton-mid/internal/http/publicapi"
	"github.com/dalot/go-skeleton-mid/internal/middlewares"
	"github.com/dalot/go-skeleton-mid/internal/services/admin"
	"github.com/dalot/go-skeleton-mid/internal/services/public"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	go func() {
		signalChan := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier is not blocked
		signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

		<-signalChan
		cancel()
	}()

	cfg, err := config.Parse()
	if err != nil {
		panic("could not parse environment variables: " + err.Error())
	}

	logger := cfg.Logger()

	db := database.New(true)
	publicAPIService := public.Service{
		DB: db,
	}
	publicAPIHandler := publicapi.Handler{
		Logger:  logger,
		Service: publicAPIService,
	}
	adminAPIService := admin.Service{
		DB: db,
	}
	adminAPIHandler := adminapi.Handler{
		Logger:  logger,
		Service: adminAPIService,
	}

	router := initRouter(cfg, logger)
	setupRoutes(router, publicAPIHandler, adminAPIHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.ServerPort),
		Handler: router,

		IdleTimeout:       cfg.IdleTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
	}

	g.Go(func() error {
		logger.Info().Msgf("staring server on :%d", cfg.ServerPort)
		return server.ListenAndServe()
	})

	// handle graceful shutdown in another goroutine
	g.Go(func() error {
		<-ctx.Done()
		gracefullShutdown(ctx.Err(), cfg, server, logger)
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
}

func setupRoutes(r *chi.Mux, handlers ...contracts.Handler) {
	for _, h := range handlers {
		h.Routes(r)
	}
}

func initRouter(cfg config.Config, logger zerolog.Logger) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middlewares.RequestIDHandler)
	router.Use(hlog.NewHandler(logger))
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Compress(flate.BestSpeed))
	router.Use(middlewares.JsonResponse)
	router.Use(middlewares.RequestLogWrapper)
	router.Use(middlewares.TimeoutWrapper(cfg.RequestTimeout))

	return router
}

func gracefullShutdown(
	ctxErr error,
	cfg config.Config,
	server *http.Server,
	logger zerolog.Logger) error {

	logger.Info().Msgf("received signal: %s, starting graceful shutdown...", ctxErr)

	ctx, done := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer done() // avoid a context leak

	err := server.Shutdown(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("error during gracefully shutdown")
	}

	logger.Info().Msg("Gracefully shutdown finished")

	return err
}
