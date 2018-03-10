package user_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/wilianto/spy-tracker-backend/user"
	mock_user "github.com/wilianto/spy-tracker-backend/user/mock"
)

func TestPostUser_InvalidRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserService := mock_user.NewMockService(ctrl)

	params := `{"username": "wilianto",`
	req, _ := http.NewRequest(http.MethodPost, "/users", strings.NewReader(params))
	respRec := httptest.NewRecorder()

	router := mux.NewRouter()
	user.NewHTTPHandler(router, mockUserService)
	router.ServeHTTP(respRec, req)

	expectedBodyResponse := `{"errors":["unexpected EOF"]}`
	assert.Equal(t, http.StatusBadRequest, respRec.Code)
	assert.Equal(t, "application/json", respRec.Header().Get("Content-Type"))
	assert.Equal(t, expectedBodyResponse, respRec.Body.String())
}
func TestPostUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserService := mock_user.NewMockService(ctrl)
	mockUserService.EXPECT().Register(gomock.Any()).Return(int64(2), nil)

	params := `{
		"username": "wilianto",
		"password": "123456789",
		"name": "Wilianto Indrawan"
	}`
	req, _ := http.NewRequest(http.MethodPost, "/users", strings.NewReader(params))
	respRec := httptest.NewRecorder()

	router := mux.NewRouter()
	user.NewHTTPHandler(router, mockUserService)
	router.ServeHTTP(respRec, req)

	expectedBodyResponse := `{"id":2}`
	assert.Equal(t, http.StatusCreated, respRec.Code)
	assert.Equal(t, "application/json", respRec.Header().Get("Content-Type"))
	assert.Equal(t, expectedBodyResponse, respRec.Body.String())
}
func TestPostUser_Failed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserService := mock_user.NewMockService(ctrl)
	mockUserService.EXPECT().Register(gomock.Any()).Return(int64(0), []error{errors.New("Duplicate")})

	params := `{
		"username": "wilianto",
		"password": "123456789",
		"name": "Wilianto Indrawan"
	}`
	req, _ := http.NewRequest(http.MethodPost, "/users", strings.NewReader(params))
	respRec := httptest.NewRecorder()

	router := mux.NewRouter()
	user.NewHTTPHandler(router, mockUserService)
	router.ServeHTTP(respRec, req)

	expectedBodyResponse := `{"errors":["Duplicate"]}`
	assert.Equal(t, http.StatusUnprocessableEntity, respRec.Code)
	assert.Equal(t, "application/json", respRec.Header().Get("Content-Type"))
	assert.Equal(t, expectedBodyResponse, respRec.Body.String())
}
