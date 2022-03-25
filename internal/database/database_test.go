package database

import (
	"testing"
	"time"

	"github.com/dalot/go-skeleton-mid/pkg/resources"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestStore_SetUser(t *testing.T) {
	is := require.New(t)
	s := New(false)
	tests := []struct {
		name     string
		username string
		password string
		len      int
	}{
		{
			name:     "success",
			username: "john",
			password: "doe",
			len:      2,
		},
		{
			name:     "update password of existent user",
			username: "admin",
			password: "different_password",
			len:      2,
		},
		{
			name:     "it does not validate empty strings",
			username: "",
			password: "I accept it willingfully",
			len:      3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s.SetUser(tt.username, []byte(tt.password))
			is.EqualValues(tt.password, s.users[tt.username])
			is.Len(s.users, tt.len)
		})
	}
}

func TestStore_GetHash(t *testing.T) {
	is := require.New(t)
	s := New(false)
	tests := []struct {
		name     string
		username string
		want     []byte
		err      error
	}{
		{
			name:     "success",
			username: "admin",
			want:     []byte("go-skeleton-mid"),
		},
		{
			name:     "unexistent user",
			username: "john",
			err:      ErrNotExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetHash(tt.username)
			if tt.err != nil {
				is.ErrorIs(err, tt.err)
				return
			}
			is.Nil(err)
			is.Nil(bcrypt.CompareHashAndPassword(got, tt.want))
		})
	}
}

func TestStore_UserExists(t *testing.T) {
	is := require.New(t)
	s := New(false)
	tests := []struct {
		name     string
		username string
		want     bool
	}{
		{
			name:     "success",
			username: "admin",
			want:     true,
		},
		{
			name:     "unexistent user",
			username: "john",
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.UserExists(tt.username)
			is.EqualValues(tt.want, got)
		})
	}
}

func TestStore_SetMessage(t *testing.T) {
	is := require.New(t)
	s := New(false)
	now := time.Now()
	input := &resources.Message{
		ID:        "ID",
		Name:      "NAME",
		Email:     "email@email.com",
		Text:      "text",
		CreatedAt: now,
	}
	s.SetMessage(input)
	is.EqualValues(input, s.messages[input.ID])
}

func TestStore_GetMessage(t *testing.T) {
	is := require.New(t)
	s := New(false)
	now := time.Now()
	want := &resources.Message{
		ID:        "ID",
		Name:      "NAME",
		Email:     "email@email.com",
		Text:      "text",
		CreatedAt: now,
	}
	s.SetMessage(want)

	tests := []struct {
		name string
		id   string
		want *resources.Message
		err  error
	}{
		{
			name: "success",
			id:   want.ID,
			want: want,
		},
		{
			name: "not found",
			id:   "NOT_FOUND",
			err:  ErrNotExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetMessage(tt.id)
			if tt.err != nil {
				is.ErrorIs(err, tt.err)
				return
			}
			is.Nil(err)
			is.EqualValues(tt.want, got)
			is.True(tt.want != got)
		})
	}
}

func TestStore_Messages(t *testing.T) {
	is := require.New(t)
	s := New(false)
	now := time.Now()
	oneHourLater := now.Add(time.Hour * 1)
	twoHoursLater := now.Add(time.Hour * 2)

	seederMsgs := []*resources.Message{
		{
			ID:        "1",
			CreatedAt: now,
		},
		{
			ID:        "2",
			CreatedAt: oneHourLater,
		},
		{
			ID:        "3",
			CreatedAt: twoHoursLater,
		},
	}
	for _, m := range seederMsgs {
		s.SetMessage(m)
	}

	got := s.Messages()
	is.Len(got, len(seederMsgs))
	for i := 1; i < len(got); i++ {
		is.True(got[i].CreatedAt.Before(got[i-1].CreatedAt))
	}
}

func TestStore_MessageExists(t *testing.T) {
	is := require.New(t)
	s := New(false)
	want := &resources.Message{
		ID: "ID",
	}
	s.SetMessage(want)

	tests := []struct {
		name string
		id   string
		want bool
	}{
		{
			name: "success",
			id:   "ID",
			want: true,
		},
		{
			name: "unexistent message",
			id:   "not found",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.MessageExists(tt.id)
			is.EqualValues(tt.want, got)
		})
	}
}

func TestStore_UpdateText(t *testing.T) {
	is := require.New(t)
	s := New(false)
	want := &resources.Message{
		ID: "ID",
	}
	s.SetMessage(want)

	tests := []struct {
		name string
		id   string
		text string
		err  error
	}{
		{
			name: "success",
			id:   "ID",
			text: "NEW TEXT",
		},
		{
			name: "unexistent message",
			id:   "NOT_FOUND",
			err:  ErrNotExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.UpdateText(tt.id, tt.text)
			if tt.err != nil {
				is.ErrorIs(err, tt.err)
				return
			}
			is.Nil(err)
			is.EqualValues(tt.text, got.Text)
			is.True(want != got)
		})
	}
}
