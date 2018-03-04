package user_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/wilianto/spy-tracker-backend/user"
	mock_user "github.com/wilianto/spy-tracker-backend/user/mock"
)

func TestRegister_WhenSuccess(t *testing.T) {
	usr := &user.User{
		Username: "wilianto",
		Password: "hash_password",
		Name:     "Wilianto Indrawan",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockValidator := mock_user.NewMockValidator(ctrl)
	mockValidator.EXPECT().Validate(gomock.Any()).Return(map[string]error{})

	mockRepo := mock_user.NewMockRepository(ctrl)
	mockRepo.EXPECT().Store(usr).Return(int64(2), nil)

	service := user.NewService(mockRepo, mockValidator)
	userID, err := service.Register(usr)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), userID)
}
func TestRegister_WhenDataNotValid(t *testing.T) {
	usr := &user.User{
		Username: "wilianto",
		Password: "hash_password",
		Name:     "Wilianto Indrawan",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	validatorErrs := map[string]error{
		"error": errors.New("Some error"),
	}
	mockValidator := mock_user.NewMockValidator(ctrl)
	mockValidator.EXPECT().Validate(gomock.Any()).Return(validatorErrs)

	mockRepo := mock_user.NewMockRepository(ctrl)
	mockRepo.EXPECT().Store(usr).Times(0)

	service := user.NewService(mockRepo, mockValidator)
	userID, err := service.Register(usr)
	assert.Error(t, err)
	assert.Equal(t, int64(0), userID)
}
