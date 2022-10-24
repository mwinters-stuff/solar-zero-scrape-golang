// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	middleware "github.com/go-openapi/runtime/middleware"
	http_api "github.com/mwinters-stuff/solar-zero-scrape-golang/restapi/operations/http_api"

	mock "github.com/stretchr/testify/mock"
)

// GetHandler is an autogenerated mock type for the GetHandler type
type GetHandler struct {
	mock.Mock
}

type GetHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *GetHandler) EXPECT() *GetHandler_Expecter {
	return &GetHandler_Expecter{mock: &_m.Mock}
}

// Handle provides a mock function with given fields: _a0
func (_m *GetHandler) Handle(_a0 http_api.GetParams) middleware.Responder {
	ret := _m.Called(_a0)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(http_api.GetParams) middleware.Responder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// GetHandler_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type GetHandler_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//  - _a0 http_api.GetParams
func (_e *GetHandler_Expecter) Handle(_a0 interface{}) *GetHandler_Handle_Call {
	return &GetHandler_Handle_Call{Call: _e.mock.On("Handle", _a0)}
}

func (_c *GetHandler_Handle_Call) Run(run func(_a0 http_api.GetParams)) *GetHandler_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http_api.GetParams))
	})
	return _c
}

func (_c *GetHandler_Handle_Call) Return(_a0 middleware.Responder) *GetHandler_Handle_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewGetHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewGetHandler creates a new instance of GetHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGetHandler(t mockConstructorTestingTNewGetHandler) *GetHandler {
	mock := &GetHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
