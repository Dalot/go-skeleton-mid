package adminapi

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dalot/go-skeleton-mid/internal/contracts"
	"github.com/dalot/go-skeleton-mid/pkg/api"
	"github.com/dalot/go-skeleton-mid/pkg/resources"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const messageContextKey = "message"

type Handler struct {
	Service contracts.AdminAPIService
	Logger  zerolog.Logger
}

func (h Handler) Routes(r *chi.Mux) {
	r.Group(func(r chi.Router) {
		r.Use(h.Admin)
		r.Route("/admin", func(r chi.Router) {
			r.Route("/messages", func(r chi.Router) {
				r.With(paginate).Get("/", h.Messages)
				r.Get("/", h.Messages)

				r.Route("/{id}", func(r chi.Router) {
					r.Use(h.loadMessage) // Load the *Message on the request context
					r.Get("/", h.GetMessage)
					r.Patch("/", h.UpdateText)
				})
			})

		})
	})
}

func (h *Handler) Messages(w http.ResponseWriter, r *http.Request) {
	msgs := h.Service.Messages()
	resp := resources.MessagesListResponse{
		Messages: msgs,
	}

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling checkout response")
		api.InternalServerError(w)
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Error().Err(err).Msg("error writing to http reply")
		api.InternalServerError(w)
		return

	}
}

func (h *Handler) GetMessage(w http.ResponseWriter, r *http.Request) {
	value := r.Context().Value(messageContextKey)
	msg, ok := value.(*resources.Message)
	if !ok {
		log.Error().Msg("could not convert interface to message from context")
		api.InternalServerError(w)
		return
	}

	resp := resources.MessageResponse{
		Message: msg,
	}
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling checkout response")
		api.InternalServerError(w)
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Error().Err(err).Msg("error writing to http reply")
		api.InternalServerError(w)
		return

	}

}

func (h *Handler) UpdateText(w http.ResponseWriter, r *http.Request) {
	value := r.Context().Value(messageContextKey)
	msg, ok := value.(*resources.Message)
	if !ok {
		log.Error().Msg("could not convert interface to message from context")
		api.InternalServerError(w)
		return
	}

	text, err := textFromRequest(r)
	if err != nil {
		log.Error().Msg("could not parse text from request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newMsg, err := h.Service.UpdateText(msg.ID, text)
	if err != nil {
		log.Error().Err(err).Msg("could not update text body")
		api.NotFound(w)
		return
	}

	resp := resources.MessageResponse{
		Message: newMsg,
	}
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling checkout response")
		api.InternalServerError(w)
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Error().Err(err).Msg("error writing to http reply")
		api.InternalServerError(w)
		return

	}
}

type updateTextBody struct {
	Text string `json:"text"`
}

func textFromRequest(r *http.Request) (string, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	text := updateTextBody{}
	if jsonErr := json.Unmarshal(body, &text); jsonErr != nil {
		return "", jsonErr
	}
	return text.Text, nil
}

func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
}

// loadMessage middleware is used to load a Message object from
// the URL parameters passed through as the request. In case
// the Article could not be found, we stop here and return a 404.
func (h *Handler) loadMessage(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		uuid, err := uuid.Parse(id)
		if err != nil {
			msg := "message id is not valid, received: " + id
			log.Debug().Msg(msg)
			api.BadRequest(w)
			return
		}

		msg, err := h.Service.GetMessage(uuid.String())
		if err != nil {
			msg := "not found, message id : " + id
			log.Debug().Msg(msg)
			api.NotFound(w)
			return
		}

		ctx := context.WithValue(r.Context(), messageContextKey, msg)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *Handler) Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			api.WriteAPIError(w, "not authorized", http.StatusUnauthorized)
			return
		}

		_, err := h.Service.Auth(username, password)
		if err != nil {
			api.WriteAPIError(w, "wrong credentials", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
