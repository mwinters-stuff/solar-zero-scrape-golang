// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	middleware "github.com/go-openapi/runtime/middleware"
	kubernetes "github.com/mwinters-stuff/solar-zero-scrape-golang/api/restapi/operations/kubernetes"

	mock "github.com/stretchr/testify/mock"
)

// GetReadyzHandler is an autogenerated mock type for the GetReadyzHandler type
type GetReadyzHandler struct {
	mock.Mock
}

type GetReadyzHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *GetReadyzHandler) EXPECT() *GetReadyzHandler_Expecter {
	return &GetReadyzHandler_Expecter{mock: &_m.Mock}
}

// Handle provides a mock function with given fields: _a0
func (_m *GetReadyzHandler) Handle(_a0 kubernetes.GetReadyzParams) middleware.Responder {
	ret := _m.Called(_a0)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(kubernetes.GetReadyzParams) middleware.Responder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// GetReadyzHandler_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type GetReadyzHandler_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//  - _a0 kubernetes.GetReadyzParams
func (_e *GetReadyzHandler_Expecter) Handle(_a0 interface{}) *GetReadyzHandler_Handle_Call {
	return &GetReadyzHandler_Handle_Call{Call: _e.mock.On("Handle", _a0)}
}

func (_c *GetReadyzHandler_Handle_Call) Run(run func(_a0 kubernetes.GetReadyzParams)) *GetReadyzHandler_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(kubernetes.GetReadyzParams))
	})
	return _c
}

func (_c *GetReadyzHandler_Handle_Call) Return(_a0 middleware.Responder) *GetReadyzHandler_Handle_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewGetReadyzHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewGetReadyzHandler creates a new instance of GetReadyzHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGetReadyzHandler(t mockConstructorTestingTNewGetReadyzHandler) *GetReadyzHandler {
	mock := &GetReadyzHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
