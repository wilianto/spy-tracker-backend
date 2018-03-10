package user

//go:generate mockgen -destination=mock/mock_service.go github.com/wilianto/spy-tracker-backend/user Service

//Service interface for communicating with user use cases
type Service interface {
	Register(user *User) (userID int64, errs []error)
}

type service struct {
	userRepo      Repository
	userValidator Validator
}

func (s *service) Register(user *User) (int64, []error) {
	errs := s.userValidator.Validate(user)
	if len(errs) != 0 {
		return 0, errs
	}
	id, err := s.userRepo.Store(user)
	if err != nil {
		return id, []error{err}
	}
	return id, []error{}
}

//NewService instatiane new user service with user repository in param
func NewService(userRepo Repository, userValidator Validator) Service {
	return &service{userRepo, userValidator}
}
