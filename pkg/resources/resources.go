package resources

import (
	"time"

	"github.com/google/uuid"
)

// Message represents the message in the service layer.
type Message struct {
	ID        string
	Name      string
	Email     string
	Text      string
	CreatedAt time.Time
}

// CreateMessageInput represents the input in the request body that our server receives
type CreateMessageInput struct {
	Name  string
	Email string
	Text  string
}

func (i *CreateMessageInput) NewMsg() *Message {
	id := uuid.New().String()
	now := time.Now().UTC()
	return &Message{
		ID:        id,
		Name:      i.Name,
		Email:     i.Email,
		Text:      i.Text,
		CreatedAt: now,
	}
}

type MessageResponse struct {
	Message *Message
}

type MessagesListResponse struct {
	Messages []*Message
}
