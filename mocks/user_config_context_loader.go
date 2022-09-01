// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/yolo-sh/hetzner-cloud-provider/userconfig (interfaces: ContextLoader)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	userconfig "github.com/yolo-sh/hetzner-cloud-provider/userconfig"
)

// UserConfigContextLoader is a mock of ContextLoader interface.
type UserConfigContextLoader struct {
	ctrl     *gomock.Controller
	recorder *UserConfigContextLoaderMockRecorder
}

// UserConfigContextLoaderMockRecorder is the mock recorder for UserConfigContextLoader.
type UserConfigContextLoaderMockRecorder struct {
	mock *UserConfigContextLoader
}

// NewUserConfigContextLoader creates a new mock instance.
func NewUserConfigContextLoader(ctrl *gomock.Controller) *UserConfigContextLoader {
	mock := &UserConfigContextLoader{ctrl: ctrl}
	mock.recorder = &UserConfigContextLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *UserConfigContextLoader) EXPECT() *UserConfigContextLoaderMockRecorder {
	return m.recorder
}

// Load mocks base method.
func (m *UserConfigContextLoader) Load(arg0, arg1 string) (*userconfig.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", arg0, arg1)
	ret0, _ := ret[0].(*userconfig.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Load indicates an expected call of Load.
func (mr *UserConfigContextLoaderMockRecorder) Load(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*UserConfigContextLoader)(nil).Load), arg0, arg1)
}
