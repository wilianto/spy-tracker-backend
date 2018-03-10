package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilianto/spy-tracker-backend/user"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		user   user.User
		errors []error
	}{
		{
			user.User{
				Username: "",
				Name:     "",
				Password: "",
			}, []error{
				user.ErrInvalidUsername,
				user.ErrInvalidName,
				user.ErrInvalidPassword,
			},
		},
		{
			user.User{
				Username: "no",
				Name:     "Valid Name",
				Password: "valid_password",
			}, []error{user.ErrInvalidUsername},
		},
		{
			user.User{
				Username: "valid",
				Name:     "no",
				Password: "valid_password",
			}, []error{user.ErrInvalidName},
		},
		{
			user.User{
				Username: "valid",
				Name:     "Valid Name",
				Password: "short",
			}, []error{user.ErrInvalidPassword},
		},
	}

	validator := user.NewValidator()
	for _, testCase := range testCases {
		assert.Equal(t, testCase.errors, validator.Validate(&testCase.user))
	}
}
