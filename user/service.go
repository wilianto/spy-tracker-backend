package user

//go:generate mockgen -destination=mock/mock_service.go github.com/wilianto/spy-tracker-backend/user Service

//Service interface for communicating with user use cases
type Service interface {
	Register(user *User) (userID int64, err error)
}

type service struct {
	userRepo Repository
}

func (s *service) Register(user *User) (int64, error) {
	//TODO: add validation
	return s.userRepo.Store(user)
}

//NewService instatiane new user service with user repository in param
func NewService(userRepo Repository) Service {
	return &service{userRepo}
}
