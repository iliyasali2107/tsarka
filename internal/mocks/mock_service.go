// Code generated by MockGen. DO NOT EDIT.
// Source: tsarka/internal/service (interfaces: SelfService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSelfService is a mock of SelfService interface.
type MockSelfService struct {
	ctrl     *gomock.Controller
	recorder *MockSelfServiceMockRecorder
}

// MockSelfServiceMockRecorder is the mock recorder for MockSelfService.
type MockSelfServiceMockRecorder struct {
	mock *MockSelfService
}

// NewMockSelfService creates a new mock instance.
func NewMockSelfService(ctrl *gomock.Controller) *MockSelfService {
	mock := &MockSelfService{ctrl: ctrl}
	mock.recorder = &MockSelfServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSelfService) EXPECT() *MockSelfServiceMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockSelfService) Find(arg0 context.Context, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockSelfServiceMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockSelfService)(nil).Find), arg0, arg1)
}
