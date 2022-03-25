package publicapi

import (
	"testing"

	"github.com/dalot/go-skeleton-mid/pkg/resources"
	"github.com/stretchr/testify/require"
)

func Test_validateCreateMessageInput(t *testing.T) {
	is := require.New(t)
	tests := []struct {
		name  string
		input *resources.CreateMessageInput
		err   string
	}{
		{
			name: "short name",
			input: &resources.CreateMessageInput{
				Name: "ab",
			},
			err: "short name",
		},
		{
			// there is no point in adding more cases to the invalid email since the std lib code is already tested
			// this test is here only to validate that we call the validation from std lib
			name: "invalid email",
			input: &resources.CreateMessageInput{
				Name:  "name",
				Email: "emailexample.com",
			},
			err: "invalid email: mail: missing '@' or angle-addr",
		},
		{
			name: "empty email",
			input: &resources.CreateMessageInput{
				Name:  "name",
				Email: "",
			},
			err: "invalid email: mail: no address",
		},
		{
			name: "empty text",
			input: &resources.CreateMessageInput{
				Name:  "name",
				Email: "email@example.com",
				Text:  "",
			},
			err: "empty text",
		},
		{
			name: "success",
			input: &resources.CreateMessageInput{
				Name:  "name",
				Email: "email@example.com",
				Text:  "hello world",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCreateMessageInput(tt.input)
			if tt.err != "" {
				is.EqualError(err, tt.err)
				return
			}
			is.Nil(err)
		})
	}
}
