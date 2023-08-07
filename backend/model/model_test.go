package model

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordCardValidate(t *testing.T) {
	testCases := []struct {
		name  string
		model PasswordCard
		err   error
	}{
		{
			name: "invalid id",
			model: PasswordCard{
				ID: "",
			},
			err: errors.New("invalid id"),
		},
		{
			name: "invalid name",
			model: PasswordCard{
				ID:   "card-id",
				Name: "",
			},
			err: errors.New("invalid name"),
		},
		{
			name: "username can't be empty",
			model: PasswordCard{
				ID:       "card-id",
				Name:     "AWS",
				Username: "",
			},
			err: errors.New("username can't be empty"),
		},
		{
			name: "password can't be empty",
			model: PasswordCard{
				ID:       "card-id",
				Name:     "AWS",
				Username: "username",
				Password: "",
			},
			err: errors.New("password can't be empty"),
		},
		{
			name: "URL can't be empty",
			model: PasswordCard{
				ID:       "card-id",
				Name:     "AWS",
				Username: "username",
				Password: "supersecret",
				URL:      "",
			},
			err: errors.New("invalid URL"),
		},
		{
			name: "invalid URL",
			model: PasswordCard{
				ID:       "card-id",
				Name:     "AWS",
				Username: "username",
				Password: "supersecret",
				URL:      "%invalid%",
			},
			err: errors.New(`invalid URL provided: parse "%invalid%": invalid URL escape "%in"`),
		},
		{
			name: "🎉 valid password card",
			model: PasswordCard{
				ID:       "card-id",
				Name:     "AWS",
				Username: "username",
				Password: "supersecret",
				URL:      `https://signin.aws.amazon.com/signin?redirect_uri=https%3A%2F%2Fconsole.aws.amazon.com%2Fconsole%2Fhome%3FhashArgs%3D%2523%26isauthcode%3Dtrue%26state%3DhashArgsFromTB_us-east-2_cd755aefcc13d319&client_id=arn%3Aaws%3Asignin%3A%3A%3Aconsole%2Fcanvas&forceMobileApp=0&code_challenge=wLFrJiiggHlLfoqWwEH9EQ5tPxPrxtow-RBQB4HeEZk&code_challenge_method=SHA-256`,
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.model.Validate()
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
