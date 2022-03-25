package admin

import (
	"errors"
	"testing"

	"github.com/dalot/go-skeleton-mid/internal/database"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestService_Auth(t *testing.T) {
	is := require.New(t)
	tests := []struct {
		name     string
		username string
		password string
		dbMock   Database
		err      string
	}{
		{
			name:     "success",
			username: "username",
			password: "password",
			dbMock: &DatabaseMock{
				GetHashFunc: func(username string) ([]byte, error) {
					hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
					is.Nil(err)
					return hashedPassword, nil
				},
			},
		},
		{
			name:     "not found hash",
			username: "username",
			password: "password",
			dbMock: &DatabaseMock{
				GetHashFunc: func(username string) ([]byte, error) {
					return []byte(""), database.ErrNotExist
				},
			},
			err: NotFoundResourceError.Error(),
		},
		{
			name:     "error",
			username: "username",
			password: "password",
			dbMock: &DatabaseMock{
				GetHashFunc: func(username string) ([]byte, error) {
					return []byte(""), errors.New("new error")
				},
			},
			err: "could not get hash: new error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				DB: tt.dbMock,
			}
			got, err := s.Auth(tt.username, tt.password)
			if tt.err != "" {
				is.EqualError(err, tt.err)
				return
			}
			is.Nil(err)
			is.EqualValues(tt.username, got)
		})
	}
}
