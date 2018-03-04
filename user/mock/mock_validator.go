// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/wilianto/spy-tracker-backend/user (interfaces: Validator)

// Package mock_user is a generated GoMock package.
package mock_user

import (
	gomock "github.com/golang/mock/gomock"
	user "github.com/wilianto/spy-tracker-backend/user"
	reflect "reflect"
)

// MockValidator is a mock of Validator interface
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockValidator) Validate(arg0 *user.User) map[string]error {
	ret := m.ctrl.Call(m, "Validate", arg0)
	ret0, _ := ret[0].(map[string]error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockValidatorMockRecorder) Validate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockValidator)(nil).Validate), arg0)
}