package user

//go:generate mockgen -destination=mock/mock_repository.go github.com/wilianto/spy-tracker-backend/user Repository

//Repository interface for interacting with data source``
type Repository interface {
	Store(user *User) (ID int64, err error)
}
