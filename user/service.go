package user

import "errors"

//go:generate mockgen -destination=mock/mock_service.go github.com/wilianto/spy-tracker-backend/user Service

//Service interface for communicating with user use cases
type Service interface {
	Register(user *User) (userID int64, err error)
}

type service struct {
	userRepo      Repository
	userValidator Validator
}

func (s *service) Register(user *User) (int64, error) {
	if len(s.userValidator.Validate(user)) == 0 {
		return s.userRepo.Store(user)
	}
	//TODO: send validator error message
	return 0, errors.New("data not valid!")
}

//NewService instatiane new user service with user repository in param
func NewService(userRepo Repository, userValidator Validator) Service {
	return &service{userRepo, userValidator}
}
