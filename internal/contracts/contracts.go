package contracts

import (
	"github.com/dalot/go-skeleton-mid/pkg/resources"
	"github.com/go-chi/chi"
)

type Handler interface {
	Routes(*chi.Mux)
}

type PublicAPIService interface {
	CreateMessage(*resources.CreateMessageInput) *resources.Message
}

type AdminAPIService interface {
	Auth(username, password string) (string, error)

	Messages() []*resources.Message
	GetMessage(id string) (*resources.Message, error)
	UpdateText(id, text string) (*resources.Message, error)
}
