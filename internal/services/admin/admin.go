package admin

import (
	"errors"
	"fmt"

	"github.com/dalot/go-skeleton-mid/internal/database"
	"github.com/dalot/go-skeleton-mid/pkg/resources"
	"golang.org/x/crypto/bcrypt"
)

var NotFoundResourceError = errors.New("could not find resource")

//go:generate moq -out ./db_mock.go . Database
type Database interface {
	UserExists(key string) bool
	GetHash(username string) ([]byte, error)

	Messages() []*resources.Message
	GetMessage(id string) (*resources.Message, error)
	UpdateText(id, text string) (*resources.Message, error)
}

type Service struct {
	DB Database
}

func (s Service) Messages() []*resources.Message {
	return s.DB.Messages()
}

func (s Service) GetMessage(id string) (*resources.Message, error) {
	return s.DB.GetMessage(id)
}

func (s Service) UpdateText(id, text string) (*resources.Message, error) {
	return s.DB.UpdateText(id, text)
}
func (s Service) Auth(username, password string) (string, error) {
	hash, err := s.DB.GetHash(username)
	if err != nil {
		if errors.Is(err, database.ErrNotExist) {
			return "", NotFoundResourceError
		}
		return "", fmt.Errorf("could not get hash: %s", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return "", fmt.Errorf("could not compare hashes: %s", err)
	}

	return username, nil
}
