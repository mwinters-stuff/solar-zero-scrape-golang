// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	write "github.com/influxdata/influxdb-client-go/v2/api/write"
	mock "github.com/stretchr/testify/mock"
)

// WriteAPIBlocking is an autogenerated mock type for the WriteAPIBlocking type
type WriteAPIBlocking struct {
	mock.Mock
}

type WriteAPIBlocking_Expecter struct {
	mock *mock.Mock
}

func (_m *WriteAPIBlocking) EXPECT() *WriteAPIBlocking_Expecter {
	return &WriteAPIBlocking_Expecter{mock: &_m.Mock}
}

// EnableBatching provides a mock function with given fields:
func (_m *WriteAPIBlocking) EnableBatching() {
	_m.Called()
}

// WriteAPIBlocking_EnableBatching_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnableBatching'
type WriteAPIBlocking_EnableBatching_Call struct {
	*mock.Call
}

// EnableBatching is a helper method to define mock.On call
func (_e *WriteAPIBlocking_Expecter) EnableBatching() *WriteAPIBlocking_EnableBatching_Call {
	return &WriteAPIBlocking_EnableBatching_Call{Call: _e.mock.On("EnableBatching")}
}

func (_c *WriteAPIBlocking_EnableBatching_Call) Run(run func()) *WriteAPIBlocking_EnableBatching_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *WriteAPIBlocking_EnableBatching_Call) Return() *WriteAPIBlocking_EnableBatching_Call {
	_c.Call.Return()
	return _c
}

// Flush provides a mock function with given fields: ctx
func (_m *WriteAPIBlocking) Flush(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteAPIBlocking_Flush_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Flush'
type WriteAPIBlocking_Flush_Call struct {
	*mock.Call
}

// Flush is a helper method to define mock.On call
//  - ctx context.Context
func (_e *WriteAPIBlocking_Expecter) Flush(ctx interface{}) *WriteAPIBlocking_Flush_Call {
	return &WriteAPIBlocking_Flush_Call{Call: _e.mock.On("Flush", ctx)}
}

func (_c *WriteAPIBlocking_Flush_Call) Run(run func(ctx context.Context)) *WriteAPIBlocking_Flush_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *WriteAPIBlocking_Flush_Call) Return(_a0 error) *WriteAPIBlocking_Flush_Call {
	_c.Call.Return(_a0)
	return _c
}

// WritePoint provides a mock function with given fields: ctx, point
func (_m *WriteAPIBlocking) WritePoint(ctx context.Context, point ...*write.Point) error {
	_va := make([]interface{}, len(point))
	for _i := range point {
		_va[_i] = point[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...*write.Point) error); ok {
		r0 = rf(ctx, point...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteAPIBlocking_WritePoint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WritePoint'
type WriteAPIBlocking_WritePoint_Call struct {
	*mock.Call
}

// WritePoint is a helper method to define mock.On call
//  - ctx context.Context
//  - point ...*write.Point
func (_e *WriteAPIBlocking_Expecter) WritePoint(ctx interface{}, point ...interface{}) *WriteAPIBlocking_WritePoint_Call {
	return &WriteAPIBlocking_WritePoint_Call{Call: _e.mock.On("WritePoint",
		append([]interface{}{ctx}, point...)...)}
}

func (_c *WriteAPIBlocking_WritePoint_Call) Run(run func(ctx context.Context, point ...*write.Point)) *WriteAPIBlocking_WritePoint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*write.Point, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(*write.Point)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *WriteAPIBlocking_WritePoint_Call) Return(_a0 error) *WriteAPIBlocking_WritePoint_Call {
	_c.Call.Return(_a0)
	return _c
}

// WriteRecord provides a mock function with given fields: ctx, line
func (_m *WriteAPIBlocking) WriteRecord(ctx context.Context, line ...string) error {
	_va := make([]interface{}, len(line))
	for _i := range line {
		_va[_i] = line[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...string) error); ok {
		r0 = rf(ctx, line...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteAPIBlocking_WriteRecord_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteRecord'
type WriteAPIBlocking_WriteRecord_Call struct {
	*mock.Call
}

// WriteRecord is a helper method to define mock.On call
//  - ctx context.Context
//  - line ...string
func (_e *WriteAPIBlocking_Expecter) WriteRecord(ctx interface{}, line ...interface{}) *WriteAPIBlocking_WriteRecord_Call {
	return &WriteAPIBlocking_WriteRecord_Call{Call: _e.mock.On("WriteRecord",
		append([]interface{}{ctx}, line...)...)}
}

func (_c *WriteAPIBlocking_WriteRecord_Call) Run(run func(ctx context.Context, line ...string)) *WriteAPIBlocking_WriteRecord_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *WriteAPIBlocking_WriteRecord_Call) Return(_a0 error) *WriteAPIBlocking_WriteRecord_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewWriteAPIBlocking interface {
	mock.TestingT
	Cleanup(func())
}

// NewWriteAPIBlocking creates a new instance of WriteAPIBlocking. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWriteAPIBlocking(t mockConstructorTestingTNewWriteAPIBlocking) *WriteAPIBlocking {
	mock := &WriteAPIBlocking{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
