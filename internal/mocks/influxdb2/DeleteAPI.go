// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/influxdata/influxdb-client-go/v2/domain"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// DeleteAPI is an autogenerated mock type for the DeleteAPI type
type DeleteAPI struct {
	mock.Mock
}

type DeleteAPI_Expecter struct {
	mock *mock.Mock
}

func (_m *DeleteAPI) EXPECT() *DeleteAPI_Expecter {
	return &DeleteAPI_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, org, bucket, start, stop, predicate
func (_m *DeleteAPI) Delete(ctx context.Context, org *domain.Organization, bucket *domain.Bucket, start time.Time, stop time.Time, predicate string) error {
	ret := _m.Called(ctx, org, bucket, start, stop, predicate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Organization, *domain.Bucket, time.Time, time.Time, string) error); ok {
		r0 = rf(ctx, org, bucket, start, stop, predicate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAPI_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type DeleteAPI_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//  - ctx context.Context
//  - org *domain.Organization
//  - bucket *domain.Bucket
//  - start time.Time
//  - stop time.Time
//  - predicate string
func (_e *DeleteAPI_Expecter) Delete(ctx interface{}, org interface{}, bucket interface{}, start interface{}, stop interface{}, predicate interface{}) *DeleteAPI_Delete_Call {
	return &DeleteAPI_Delete_Call{Call: _e.mock.On("Delete", ctx, org, bucket, start, stop, predicate)}
}

func (_c *DeleteAPI_Delete_Call) Run(run func(ctx context.Context, org *domain.Organization, bucket *domain.Bucket, start time.Time, stop time.Time, predicate string)) *DeleteAPI_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Organization), args[2].(*domain.Bucket), args[3].(time.Time), args[4].(time.Time), args[5].(string))
	})
	return _c
}

func (_c *DeleteAPI_Delete_Call) Return(_a0 error) *DeleteAPI_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// DeleteWithID provides a mock function with given fields: ctx, orgID, bucketID, start, stop, predicate
func (_m *DeleteAPI) DeleteWithID(ctx context.Context, orgID string, bucketID string, start time.Time, stop time.Time, predicate string) error {
	ret := _m.Called(ctx, orgID, bucketID, start, stop, predicate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, time.Time, time.Time, string) error); ok {
		r0 = rf(ctx, orgID, bucketID, start, stop, predicate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAPI_DeleteWithID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteWithID'
type DeleteAPI_DeleteWithID_Call struct {
	*mock.Call
}

// DeleteWithID is a helper method to define mock.On call
//  - ctx context.Context
//  - orgID string
//  - bucketID string
//  - start time.Time
//  - stop time.Time
//  - predicate string
func (_e *DeleteAPI_Expecter) DeleteWithID(ctx interface{}, orgID interface{}, bucketID interface{}, start interface{}, stop interface{}, predicate interface{}) *DeleteAPI_DeleteWithID_Call {
	return &DeleteAPI_DeleteWithID_Call{Call: _e.mock.On("DeleteWithID", ctx, orgID, bucketID, start, stop, predicate)}
}

func (_c *DeleteAPI_DeleteWithID_Call) Run(run func(ctx context.Context, orgID string, bucketID string, start time.Time, stop time.Time, predicate string)) *DeleteAPI_DeleteWithID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(time.Time), args[4].(time.Time), args[5].(string))
	})
	return _c
}

func (_c *DeleteAPI_DeleteWithID_Call) Return(_a0 error) *DeleteAPI_DeleteWithID_Call {
	_c.Call.Return(_a0)
	return _c
}

// DeleteWithName provides a mock function with given fields: ctx, orgName, bucketName, start, stop, predicate
func (_m *DeleteAPI) DeleteWithName(ctx context.Context, orgName string, bucketName string, start time.Time, stop time.Time, predicate string) error {
	ret := _m.Called(ctx, orgName, bucketName, start, stop, predicate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, time.Time, time.Time, string) error); ok {
		r0 = rf(ctx, orgName, bucketName, start, stop, predicate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAPI_DeleteWithName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteWithName'
type DeleteAPI_DeleteWithName_Call struct {
	*mock.Call
}

// DeleteWithName is a helper method to define mock.On call
//  - ctx context.Context
//  - orgName string
//  - bucketName string
//  - start time.Time
//  - stop time.Time
//  - predicate string
func (_e *DeleteAPI_Expecter) DeleteWithName(ctx interface{}, orgName interface{}, bucketName interface{}, start interface{}, stop interface{}, predicate interface{}) *DeleteAPI_DeleteWithName_Call {
	return &DeleteAPI_DeleteWithName_Call{Call: _e.mock.On("DeleteWithName", ctx, orgName, bucketName, start, stop, predicate)}
}

func (_c *DeleteAPI_DeleteWithName_Call) Run(run func(ctx context.Context, orgName string, bucketName string, start time.Time, stop time.Time, predicate string)) *DeleteAPI_DeleteWithName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(time.Time), args[4].(time.Time), args[5].(string))
	})
	return _c
}

func (_c *DeleteAPI_DeleteWithName_Call) Return(_a0 error) *DeleteAPI_DeleteWithName_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewDeleteAPI interface {
	mock.TestingT
	Cleanup(func())
}

// NewDeleteAPI creates a new instance of DeleteAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDeleteAPI(t mockConstructorTestingTNewDeleteAPI) *DeleteAPI {
	mock := &DeleteAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}