package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilianto/spy-tracker-backend/user"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		user   user.User
		errors map[string]error
	}{
		{
			user.User{
				Username: "",
				Name:     "",
				Password: "",
			}, map[string]error{
				"username": user.ErrInvalidUsername,
				"name":     user.ErrInvalidName,
				"password": user.ErrInvalidPassword,
			},
		},
		{
			user.User{
				Username: "no",
				Name:     "Valid Name",
				Password: "valid_password",
			}, map[string]error{"username": user.ErrInvalidUsername},
		},
		{
			user.User{
				Username: "valid",
				Name:     "no",
				Password: "valid_password",
			}, map[string]error{"name": user.ErrInvalidName},
		},
		{
			user.User{
				Username: "valid",
				Name:     "Valid Name",
				Password: "short",
			}, map[string]error{"password": user.ErrInvalidPassword},
		},
	}

	validator := user.NewValidator()
	for _, testCase := range testCases {
		assert.Equal(t, testCase.errors, validator.Validate(&testCase.user))
	}
}
