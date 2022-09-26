// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	// http "github.com/influxdata/influxdb-client-go/v2/api/http"
	mock "github.com/stretchr/testify/mock"

	// write "github.com/influxdata/influxdb-client-go/v2/internal/write"
)

// BatchErrorCallback is an autogenerated mock type for the BatchErrorCallback type
type BatchErrorCallback struct {
	mock.Mock
}

type BatchErrorCallback_Expecter struct {
	mock *mock.Mock
}

func (_m *BatchErrorCallback) EXPECT() *BatchErrorCallback_Expecter {
	return &BatchErrorCallback_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: batch, error2
// func (_m *BatchErrorCallback) Execute(batch *write.Batch, error2 http.Error) bool {
// 	ret := _m.Called(batch, error2)

// 	var r0 bool
// 	if rf, ok := ret.Get(0).(func(*write.Batch, http.Error) bool); ok {
// 		r0 = rf(batch, error2)
// 	} else {
// 		r0 = ret.Get(0).(bool)
// 	}

// 	return r0
// }

// BatchErrorCallback_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type BatchErrorCallback_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//  - batch *write.Batch
//  - error2 http.Error
func (_e *BatchErrorCallback_Expecter) Execute(batch interface{}, error2 interface{}) *BatchErrorCallback_Execute_Call {
	return &BatchErrorCallback_Execute_Call{Call: _e.mock.On("Execute", batch, error2)}
}

// func (_c *BatchErrorCallback_Execute_Call) Run(run func(batch *write.Batch, error2 http.Error)) *BatchErrorCallback_Execute_Call {
// 	_c.Call.Run(func(args mock.Arguments) {
// 		run(args[0].(*write.Batch), args[1].(http.Error))
// 	})
// 	return _c
// }

func (_c *BatchErrorCallback_Execute_Call) Return(_a0 bool) *BatchErrorCallback_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewBatchErrorCallback interface {
	mock.TestingT
	Cleanup(func())
}

// NewBatchErrorCallback creates a new instance of BatchErrorCallback. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBatchErrorCallback(t mockConstructorTestingTNewBatchErrorCallback) *BatchErrorCallback {
	mock := &BatchErrorCallback{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
