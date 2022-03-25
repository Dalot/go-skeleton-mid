package publicapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/mail"

	"github.com/dalot/go-skeleton-mid/internal/contracts"
	"github.com/dalot/go-skeleton-mid/pkg/api"
	"github.com/dalot/go-skeleton-mid/pkg/resources"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	Service contracts.PublicAPIService
	Logger  zerolog.Logger
}

func (h Handler) Routes(r *chi.Mux) {
	r.Post("/message", h.CreateMessage)
}

func (h *Handler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	createMessageInput, err := parseCreateMessageInput(r)
	if err != nil {
		msg := fmt.Errorf("could not parse request input: %s", err).Error()
		log.Error().Err(err).Msg(msg)
		api.WriteAPIError(w, "wrong input", http.StatusBadRequest)
		return
	}

	if err := validateCreateMessageInput(createMessageInput); err != nil {
		msg := fmt.Errorf("validation error: %s", err).Error()
		log.Error().Msg(msg)
		api.WriteAPIError(w, msg, http.StatusBadRequest)
		return
	}

	msg := h.Service.CreateMessage(createMessageInput)

	createMessageResponse := resources.MessageResponse{
		Message: msg,
	}
	jsonResponse, err := json.Marshal(createMessageResponse)
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

func validateCreateMessageInput(input *resources.CreateMessageInput) error {
	if len(input.Name) < 3 {
		return errors.New("short name")
	}

	if _, err := mail.ParseAddress(input.Email); err != nil {
		return fmt.Errorf("invalid email: %s", err)
	}

	if len(input.Text) == 0 {
		return errors.New("empty text")
	}

	return nil
}

func parseCreateMessageInput(r *http.Request) (*resources.CreateMessageInput, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	createMessageInput := &resources.CreateMessageInput{}
	if err := json.Unmarshal(body, createMessageInput); err != nil {
		return nil, err
	}

	return createMessageInput, nil
}
