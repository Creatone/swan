package mocks

import "github.com/stretchr/testify/mock"

import "github.com/intelsdi-x/swan/pkg/executor"

// Launcher is an autogenerated mock type for the Launcher type
type Launcher struct {
	mock.Mock
}

// Launch provides a mock function with given fields:
func (_m *Launcher) Launch() (executor.TaskHandle, error) {
	ret := _m.Called()

	var r0 executor.TaskHandle
	if rf, ok := ret.Get(0).(func() executor.TaskHandle); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(executor.TaskHandle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
