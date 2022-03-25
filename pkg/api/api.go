package api

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

const (
	ErrInternalServer = "internal server error"
	ErrBadRequest     = "bad request"
)

type APIErr struct {
	Error string
	Code  int
}

func WriteAPIError(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	resp := APIErr{
		Error: msg,
		Code:  code,
	}

	json, err := json.Marshal(resp)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling checkout response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(json)
	if err != nil {
		log.Error().Err(err).Msg("error writing to http reply")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
}

func InternalServerError(w http.ResponseWriter) {
	WriteAPIError(w, ErrInternalServer, http.StatusInternalServerError)
}

func BadRequest(w http.ResponseWriter) {
	WriteAPIError(w, ErrBadRequest, http.StatusInternalServerError)
}

func NotFound(w http.ResponseWriter) {
	WriteAPIError(w, "not found", http.StatusNotFound)
}
