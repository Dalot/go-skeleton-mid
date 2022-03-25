package public

import (
	"github.com/dalot/go-skeleton-mid/pkg/resources"
)

type Database interface {
	UserExists(key string) bool
	GetHash(username string) ([]byte, error)
	SetUser(username string, password []byte)
	SetMessage(input *resources.Message)
}

type Service struct {
	DB Database
}

func (s Service) CreateMessage(input *resources.CreateMessageInput) *resources.Message {
	msg := input.NewMsg()
	s.DB.SetMessage(msg)

	return msg
}
