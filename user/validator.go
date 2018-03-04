package user

import (
	"errors"
)

//go:generate mockgen -destination=mock/mock_validator.go github.com/wilianto/spy-tracker-backend/user Validator

var (
	// ErrInvalidUsername represents error when username is not valid
	// Username will be not valid when less than 4 chars
	ErrInvalidUsername = errors.New("Username is required at least consist of 4 chars")
	// ErrInvalidPassword represents error when password is not valid
	// Password will be not valid when less than 8 chars
	ErrInvalidPassword = errors.New("Password is required at least consist of 8 chars")
	// ErrInvalidName represents error when name is not valid
	// Name will be not valid when less than 4 chars
	ErrInvalidName = errors.New("Name is required at least consist of 4 chars")
)

//Validator validate data before save to repository
type Validator interface {
	Validate(user *User) map[string]error
}

type userValidator struct{}

func (v *userValidator) Validate(user *User) map[string]error {
	errors := make(map[string]error)

	if len(user.Username) < 4 {
		errors["username"] = ErrInvalidUsername
	}

	if len(user.Name) < 4 {
		errors["name"] = ErrInvalidName
	}

	if len(user.Password) < 8 {
		errors["password"] = ErrInvalidPassword
	}

	return errors
}

//NewValidator initialize new user validator
func NewValidator() Validator {
	return &userValidator{}
}
