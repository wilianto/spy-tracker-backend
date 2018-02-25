package user_test

import (
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

	mockRepo := mock_user.NewMockRepository(ctrl)
	mockRepo.EXPECT().Store(usr).Return(int64(2), nil)

	service := user.NewService(mockRepo)
	userID, err := service.Register(usr)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), userID)
}
