// Code generated by mockery v1.0.0. DO NOT EDIT.
package z_mocks

import exec "github.com/oscarzhao/launch/exec"

import mock "github.com/stretchr/testify/mock"

// SessionManager is an autogenerated mock type for the SessionManager type
type SessionManager struct {
	mock.Mock
}

// RegisterProc provides a mock function with given fields: name, execer
func (_m *SessionManager) RegisterProc(name string, execer exec.Execer) error {
	ret := _m.Called(name, execer)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, exec.Execer) error); ok {
		r0 = rf(name, execer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartProc provides a mock function with given fields: name
func (_m *SessionManager) StartProc(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StopProc provides a mock function with given fields: name
func (_m *SessionManager) StopProc(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
