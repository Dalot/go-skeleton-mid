package database

import (
	"errors"
	"sort"
	"sync"

	"github.com/dalot/go-skeleton-mid/pkg/resources"
	"golang.org/x/crypto/bcrypt"
)

var ErrNotExist = errors.New("key does not exist")
var ErrNoMatch = errors.New("no keys match")

type Store struct {
	messagesLock sync.RWMutex
	usersLock    sync.RWMutex
	messages     map[string]*resources.Message // the key is a username
	users        map[string][]byte             // the key is a username and the value is a password hash
}

// New creates a new Store for the database. It can panic when loading
// the csv messages from a file so make.
func New(loadCSV bool) *Store {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("go-skeleton-mid"), bcrypt.DefaultCost)
	if err != nil {
		// Let's assume we can panic if we are not able to setup the db properly
		panic(err)
	}
	users := map[string][]byte{
		"admin": hashedPassword,
	}
	store := &Store{
		messages: make(map[string]*resources.Message),
		users:    users,
	}
	if loadCSV {
		err := store.LoadMessages()
		if err != nil {
			// Let's assume we can panic if we are not able to setup the db properly
			panic(err)
		}
	}

	return store
}

// UserExists checks for the existence of a user in the store
func (s *Store) UserExists(key string) bool {
	_, err := s.GetHash(key)
	if err != nil {
		return false
	}
	return true
}

// GetHash gets the Hash associated with the username
// If the username does not exist, it returns ErrNotExist
func (s *Store) GetHash(username string) ([]byte, error) {
	s.usersLock.RLock()
	pw, ok := s.users[username]
	s.usersLock.RUnlock()
	if !ok {
		return pw, ErrNotExist
	}
	return pw, nil
}

// SetUser sets a new user in the store. It expects that the password
// is already hashed, leaving the hash protocol decision to the client.
func (s *Store) SetUser(username string, password []byte) {
	s.usersLock.Lock()
	s.users[username] = password
	s.usersLock.Unlock()
}

// SetMessage sets a new message in the store.
func (s *Store) SetMessage(message *resources.Message) {
	s.messagesLock.Lock()
	s.messages[message.ID] = message
	s.messagesLock.Unlock()
}

// GetMessage gets the Message associated with id
// If the message does not exist, it returns ErrNotExist
func (s *Store) GetMessage(id string) (*resources.Message, error) {
	s.messagesLock.RLock()
	defer s.messagesLock.RUnlock()

	msg, ok := s.messages[id]
	if !ok {
		return msg, ErrNotExist
	}
	res := copyMessage(msg)
	return &res, nil
}

// Messages gets the Message associated with id
// If the message does not exist, it returns ErrNotExist
func (s *Store) Messages() []*resources.Message {
	result := make([]*resources.Message, len(s.messages))
	s.messagesLock.RLock()
	i := 0
	for _, msg := range s.messages {
		result[i] = msg
		i++
	}
	s.messagesLock.RUnlock()
	sort.Slice(result, func(i, j int) bool {
		return result[i].CreatedAt.After(result[j].CreatedAt)
	})

	return result
}

// MessageExists checks for the existence of a message in the store
func (s *Store) MessageExists(key string) bool {
	_, err := s.GetMessage(key)
	if err != nil {
		return false
	}
	return true
}

func (s *Store) UpdateText(id, text string) (*resources.Message, error) {
	if exists := s.MessageExists(id); !exists {
		return nil, ErrNotExist
	}
	s.messagesLock.Lock()
	msg := s.messages[id]
	msg.Text = text
	res := copyMessage(msg)
	s.messagesLock.Unlock()

	return &res, nil
}

func copyMessage(msg *resources.Message) resources.Message {
	return resources.Message{
		ID:        msg.ID,
		Name:      msg.Name,
		Email:     msg.Email,
		Text:      msg.Text,
		CreatedAt: msg.CreatedAt,
	}
}
