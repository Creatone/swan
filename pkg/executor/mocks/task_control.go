package mocks

import (
	"time"

	"github.com/stretchr/testify/mock"
)

// TaskControl is an autogenerated mock type for the TaskControl type
type TaskControl struct {
	mock.Mock
}

// Clean provides a mock function with given fields:
func (_m *TaskControl) Clean() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EraseOutput provides a mock function with given fields:
func (_m *TaskControl) EraseOutput() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *TaskControl) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Wait provides a mock function with given fields: timeout
func (_m *TaskControl) Wait(timeout time.Duration) bool {
	ret := _m.Called(timeout)

	var r0 bool
	if rf, ok := ret.Get(0).(func(time.Duration) bool); ok {
		r0 = rf(timeout)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
