// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import rabbitmq "github.com/tomochain/tomox-sdk/rabbitmq"

// Engine is an autogenerated mock type for the Engine type
type Engine struct {
	mock.Mock
}

// HandleOrders provides a mock function with given fields: msg
func (_m *Engine) HandleOrders(msg *rabbitmq.Message) error {
	ret := _m.Called(msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(*rabbitmq.Message) error); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
