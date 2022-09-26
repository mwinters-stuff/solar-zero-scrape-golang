// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	api "github.com/influxdata/influxdb-client-go/v2/api"

	domain "github.com/influxdata/influxdb-client-go/v2/domain"

	http "github.com/influxdata/influxdb-client-go/v2/api/http"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// AuthorizationsAPI provides a mock function with given fields:
func (_m *Client) AuthorizationsAPI() api.AuthorizationsAPI {
	ret := _m.Called()

	var r0 api.AuthorizationsAPI
	if rf, ok := ret.Get(0).(func() api.AuthorizationsAPI); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.AuthorizationsAPI)
		}
	}

	return r0
}

// Client_AuthorizationsAPI_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AuthorizationsAPI'
type Client_AuthorizationsAPI_Call struct {
	*mock.Call
}

// AuthorizationsAPI is a helper method to define mock.On call
func (_e *Client_Expecter) AuthorizationsAPI() *Client_AuthorizationsAPI_Call {
	return &Client_AuthorizationsAPI_Call{Call: _e.mock.On("AuthorizationsAPI")}
}

func (_c *Client_AuthorizationsAPI_Call) Run(run func()) *Client_AuthorizationsAPI_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_AuthorizationsAPI_Call) Return(_a0 api.AuthorizationsAPI) *Client_AuthorizationsAPI_Call {
	_c.Call.Return(_a0)
	return _c
}

// BucketsAPI provides a mock function with given fields:
func (_m *Client) BucketsAPI() api.BucketsAPI {
	ret := _m.Called()

	var r0 api.BucketsAPI
	if rf, ok := ret.Get(0).(func() api.BucketsAPI); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.BucketsAPI)
		}
	}

	return r0
}

// Client_BucketsAPI_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BucketsAPI'
type Client_BucketsAPI_Call struct {
	*mock.Call
}

// BucketsAPI is a helper method to define mock.On call
func (_e *Client_Expecter) BucketsAPI() *Client_BucketsAPI_Call {
	return &Client_BucketsAPI_Call{Call: _e.mock.On("BucketsAPI")}
}

func (_c *Client_BucketsAPI_Call) Run(run func()) *Client_BucketsAPI_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_BucketsAPI_Call) Return(_a0 api.BucketsAPI) *Client_BucketsAPI_Call {
	_c.Call.Return(_a0)
	return _c
}

// Close provides a mock function with given fields:
func (_m *Client) Close() {
	_m.Called()
}

// Client_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type Client_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *Client_Expecter) Close() *Client_Close_Call {
	return &Client_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *Client_Close_Call) Run(run func()) *Client_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_Close_Call) Return() *Client_Close_Call {
	_c.Call.Return()
	return _c
}

// DeleteAPI provides a mock function with given fields:
func (_m *Client) DeleteAPI() api.DeleteAPI {
	ret := _m.Called()

	var r0 api.DeleteAPI
	if rf, ok := ret.Get(0).(func() api.DeleteAPI); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.DeleteAPI)
		}
	}

	return r0
}

// Client_DeleteAPI_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAPI'
type Client_DeleteAPI_Call struct {
	*mock.Call
}

// DeleteAPI is a helper method to define mock.On call
func (_e *Client_Expecter) DeleteAPI() *Client_DeleteAPI_Call {
	return &Client_DeleteAPI_Call{Call: _e.mock.On("DeleteAPI")}
}

func (_c *Client_DeleteAPI_Call) Run(run func()) *Client_DeleteAPI_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_DeleteAPI_Call) Return(_a0 api.DeleteAPI) *Client_DeleteAPI_Call {
	_c.Call.Return(_a0)
	return _c
}

// HTTPService provides a mock function with given fields:
func (_m *Client) HTTPService() http.Service {
	ret := _m.Called()

	var r0 http.Service
	if rf, ok := ret.Get(0).(func() http.Service); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Service)
		}
	}

	return r0
}

// Client_HTTPService_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HTTPService'
type Client_HTTPService_Call struct {
	*mock.Call
}

// HTTPService is a helper method to define mock.On call
func (_e *Client_Expecter) HTTPService() *Client_HTTPService_Call {
	return &Client_HTTPService_Call{Call: _e.mock.On("HTTPService")}
}

func (_c *Client_HTTPService_Call) Run(run func()) *Client_HTTPService_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_HTTPService_Call) Return(_a0 http.Service) *Client_HTTPService_Call {
	_c.Call.Return(_a0)
	return _c
}

// Health provides a mock function with given fields: ctx
func (_m *Client) Health(ctx context.Context) (*domain.HealthCheck, error) {
	ret := _m.Called(ctx)

	var r0 *domain.HealthCheck
	if rf, ok := ret.Get(0).(func(context.Context) *domain.HealthCheck); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.HealthCheck)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_Health_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Health'
type Client_Health_Call struct {
	*mock.Call
}

// Health is a helper method to define mock.On call
//  - ctx context.Context
func (_e *Client_Expecter) Health(ctx interface{}) *Client_Health_Call {
	return &Client_Health_Call{Call: _e.mock.On("Health", ctx)}
}

func (_c *Client_Health_Call) Run(run func(ctx context.Context)) *Client_Health_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Client_Health_Call) Return(_a0 *domain.HealthCheck, _a1 error) *Client_Health_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// LabelsAPI provides a mock function with given fields:
func (_m *Client) LabelsAPI() api.LabelsAPI {
	ret := _m.Called()

	var r0 api.LabelsAPI
	if rf, ok := ret.Get(0).(func() api.LabelsAPI); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.LabelsAPI)
		}
	}

	return r0
}

// Client_LabelsAPI_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LabelsAPI'
type Client_LabelsAPI_Call struct {
	*mock.Call
}

// LabelsAPI is a helper method to define mock.On call
func (_e *Client_Expecter) LabelsAPI() *Client_LabelsAPI_Call {
	return &Client_LabelsAPI_Call{Call: _e.mock.On("LabelsAPI")}
}

func (_c *Client_LabelsAPI_Call) Run(run func()) *Client_LabelsAPI_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_LabelsAPI_Call) Return(_a0 api.LabelsAPI) *Client_LabelsAPI_Call {
	_c.Call.Return(_a0)
	return _c
}

// Options provides a mock function with given fields:
func (_m *Client) Options() *influxdb2.Options {
	ret := _m.Called()

	var r0 *influxdb2.Options
	if rf, ok := ret.Get(0).(func() *influxdb2.Options); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*influxdb2.Options)
		}
	}

	return r0
}

// Client_Options_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Options'
type Client_Options_Call struct {
	*mock.Call
}

// Options is a helper method to define mock.On call
func (_e *Client_Expecter) Options() *Client_Options_Call {
	return &Client_Options_Call{Call: _e.mock.On("Options")}
}

func (_c *Client_Options_Call) Run(run func()) *Client_Options_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_Options_Call) Return(_a0 *influxdb2.Options) *Client_Options_Call {
	_c.Call.Return(_a0)
	return _c
}

// OrganizationsAPI provides a mock function with given fields:
func (_m *Client) OrganizationsAPI() api.OrganizationsAPI {
	ret := _m.Called()

	var r0 api.OrganizationsAPI
	if rf, ok := ret.Get(0).(func() api.OrganizationsAPI); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.OrganizationsAPI)
		}
	}

	return r0
}

// Client_OrganizationsAPI_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OrganizationsAPI'
type Client_OrganizationsAPI_Call struct {
	*mock.Call
}

// OrganizationsAPI is a helper method to define mock.On call
func (_e *Client_Expecter) OrganizationsAPI() *Client_OrganizationsAPI_Call {
	return &Client_OrganizationsAPI_Call{Call: _e.mock.On("OrganizationsAPI")}
}

func (_c *Client_OrganizationsAPI_Call) Run(run func()) *Client_OrganizationsAPI_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_OrganizationsAPI_Call) Return(_a0 api.OrganizationsAPI) *Client_OrganizationsAPI_Call {
	_c.Call.Return(_a0)
	return _c
}

// Ping provides a mock function with given fields: ctx
func (_m *Client) Ping(ctx context.Context) (bool, error) {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_Ping_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ping'
type Client_Ping_Call struct {
	*mock.Call
}

// Ping is a helper method to define mock.On call
//  - ctx context.Context
func (_e *Client_Expecter) Ping(ctx interface{}) *Client_Ping_Call {
	return &Client_Ping_Call{Call: _e.mock.On("Ping", ctx)}
}

func (_c *Client_Ping_Call) Run(run func(ctx context.Context)) *Client_Ping_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Client_Ping_Call) Return(_a0 bool, _a1 error) *Client_Ping_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// QueryAPI provides a mock function with given fields: org
func (_m *Client) QueryAPI(org string) api.QueryAPI {
	ret := _m.Called(org)

	var r0 api.QueryAPI
	if rf, ok := ret.Get(0).(func(string) api.QueryAPI); ok {
		r0 = rf(org)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.QueryAPI)
		}
	}

	return r0
}

// Client_QueryAPI_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryAPI'
type Client_QueryAPI_Call struct {
	*mock.Call
}

// QueryAPI is a helper method to define mock.On call
//  - org string
func (_e *Client_Expecter) QueryAPI(org interface{}) *Client_QueryAPI_Call {
	return &Client_QueryAPI_Call{Call: _e.mock.On("QueryAPI", org)}
}

func (_c *Client_QueryAPI_Call) Run(run func(org string)) *Client_QueryAPI_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Client_QueryAPI_Call) Return(_a0 api.QueryAPI) *Client_QueryAPI_Call {
	_c.Call.Return(_a0)
	return _c
}

// Ready provides a mock function with given fields: ctx
func (_m *Client) Ready(ctx context.Context) (*domain.Ready, error) {
	ret := _m.Called(ctx)

	var r0 *domain.Ready
	if rf, ok := ret.Get(0).(func(context.Context) *domain.Ready); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Ready)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_Ready_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ready'
type Client_Ready_Call struct {
	*mock.Call
}

// Ready is a helper method to define mock.On call
//  - ctx context.Context
func (_e *Client_Expecter) Ready(ctx interface{}) *Client_Ready_Call {
	return &Client_Ready_Call{Call: _e.mock.On("Ready", ctx)}
}

func (_c *Client_Ready_Call) Run(run func(ctx context.Context)) *Client_Ready_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Client_Ready_Call) Return(_a0 *domain.Ready, _a1 error) *Client_Ready_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ServerURL provides a mock function with given fields:
func (_m *Client) ServerURL() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Client_ServerURL_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ServerURL'
type Client_ServerURL_Call struct {
	*mock.Call
}

// ServerURL is a helper method to define mock.On call
func (_e *Client_Expecter) ServerURL() *Client_ServerURL_Call {
	return &Client_ServerURL_Call{Call: _e.mock.On("ServerURL")}
}

func (_c *Client_ServerURL_Call) Run(run func()) *Client_ServerURL_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_ServerURL_Call) Return(_a0 string) *Client_ServerURL_Call {
	_c.Call.Return(_a0)
	return _c
}

// Setup provides a mock function with given fields: ctx, username, password, org, bucket, retentionPeriodHours
func (_m *Client) Setup(ctx context.Context, username string, password string, org string, bucket string, retentionPeriodHours int) (*domain.OnboardingResponse, error) {
	ret := _m.Called(ctx, username, password, org, bucket, retentionPeriodHours)

	var r0 *domain.OnboardingResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, int) *domain.OnboardingResponse); ok {
		r0 = rf(ctx, username, password, org, bucket, retentionPeriodHours)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.OnboardingResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, int) error); ok {
		r1 = rf(ctx, username, password, org, bucket, retentionPeriodHours)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_Setup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Setup'
type Client_Setup_Call struct {
	*mock.Call
}

// Setup is a helper method to define mock.On call
//  - ctx context.Context
//  - username string
//  - password string
//  - org string
//  - bucket string
//  - retentionPeriodHours int
func (_e *Client_Expecter) Setup(ctx interface{}, username interface{}, password interface{}, org interface{}, bucket interface{}, retentionPeriodHours interface{}) *Client_Setup_Call {
	return &Client_Setup_Call{Call: _e.mock.On("Setup", ctx, username, password, org, bucket, retentionPeriodHours)}
}

func (_c *Client_Setup_Call) Run(run func(ctx context.Context, username string, password string, org string, bucket string, retentionPeriodHours int)) *Client_Setup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(string), args[5].(int))
	})
	return _c
}

func (_c *Client_Setup_Call) Return(_a0 *domain.OnboardingResponse, _a1 error) *Client_Setup_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// SetupWithToken provides a mock function with given fields: ctx, username, password, org, bucket, retentionPeriodHours, token
func (_m *Client) SetupWithToken(ctx context.Context, username string, password string, org string, bucket string, retentionPeriodHours int, token string) (*domain.OnboardingResponse, error) {
	ret := _m.Called(ctx, username, password, org, bucket, retentionPeriodHours, token)

	var r0 *domain.OnboardingResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, int, string) *domain.OnboardingResponse); ok {
		r0 = rf(ctx, username, password, org, bucket, retentionPeriodHours, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.OnboardingResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, int, string) error); ok {
		r1 = rf(ctx, username, password, org, bucket, retentionPeriodHours, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_SetupWithToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetupWithToken'
type Client_SetupWithToken_Call struct {
	*mock.Call
}

// SetupWithToken is a helper method to define mock.On call
//  - ctx context.Context
//  - username string
//  - password string
//  - org string
//  - bucket string
//  - retentionPeriodHours int
//  - token string
func (_e *Client_Expecter) SetupWithToken(ctx interface{}, username interface{}, password interface{}, org interface{}, bucket interface{}, retentionPeriodHours interface{}, token interface{}) *Client_SetupWithToken_Call {
	return &Client_SetupWithToken_Call{Call: _e.mock.On("SetupWithToken", ctx, username, password, org, bucket, retentionPeriodHours, token)}
}

func (_c *Client_SetupWithToken_Call) Run(run func(ctx context.Context, username string, password string, org string, bucket string, retentionPeriodHours int, token string)) *Client_SetupWithToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(string), args[5].(int), args[6].(string))
	})
	return _c
}

func (_c *Client_SetupWithToken_Call) Return(_a0 *domain.OnboardingResponse, _a1 error) *Client_SetupWithToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// TasksAPI provides a mock function with given fields:
func (_m *Client) TasksAPI() api.TasksAPI {
	ret := _m.Called()

	var r0 api.TasksAPI
	if rf, ok := ret.Get(0).(func() api.TasksAPI); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.TasksAPI)
		}
	}

	return r0
}

// Client_TasksAPI_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TasksAPI'
type Client_TasksAPI_Call struct {
	*mock.Call
}

// TasksAPI is a helper method to define mock.On call
func (_e *Client_Expecter) TasksAPI() *Client_TasksAPI_Call {
	return &Client_TasksAPI_Call{Call: _e.mock.On("TasksAPI")}
}

func (_c *Client_TasksAPI_Call) Run(run func()) *Client_TasksAPI_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_TasksAPI_Call) Return(_a0 api.TasksAPI) *Client_TasksAPI_Call {
	_c.Call.Return(_a0)
	return _c
}

// UsersAPI provides a mock function with given fields:
func (_m *Client) UsersAPI() api.UsersAPI {
	ret := _m.Called()

	var r0 api.UsersAPI
	if rf, ok := ret.Get(0).(func() api.UsersAPI); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.UsersAPI)
		}
	}

	return r0
}

// Client_UsersAPI_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UsersAPI'
type Client_UsersAPI_Call struct {
	*mock.Call
}

// UsersAPI is a helper method to define mock.On call
func (_e *Client_Expecter) UsersAPI() *Client_UsersAPI_Call {
	return &Client_UsersAPI_Call{Call: _e.mock.On("UsersAPI")}
}

func (_c *Client_UsersAPI_Call) Run(run func()) *Client_UsersAPI_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_UsersAPI_Call) Return(_a0 api.UsersAPI) *Client_UsersAPI_Call {
	_c.Call.Return(_a0)
	return _c
}

// WriteAPI provides a mock function with given fields: org, bucket
func (_m *Client) WriteAPI(org string, bucket string) api.WriteAPI {
	ret := _m.Called(org, bucket)

	var r0 api.WriteAPI
	if rf, ok := ret.Get(0).(func(string, string) api.WriteAPI); ok {
		r0 = rf(org, bucket)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.WriteAPI)
		}
	}

	return r0
}

// Client_WriteAPI_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteAPI'
type Client_WriteAPI_Call struct {
	*mock.Call
}

// WriteAPI is a helper method to define mock.On call
//  - org string
//  - bucket string
func (_e *Client_Expecter) WriteAPI(org interface{}, bucket interface{}) *Client_WriteAPI_Call {
	return &Client_WriteAPI_Call{Call: _e.mock.On("WriteAPI", org, bucket)}
}

func (_c *Client_WriteAPI_Call) Run(run func(org string, bucket string)) *Client_WriteAPI_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Client_WriteAPI_Call) Return(_a0 api.WriteAPI) *Client_WriteAPI_Call {
	_c.Call.Return(_a0)
	return _c
}

// WriteAPIBlocking provides a mock function with given fields: org, bucket
func (_m *Client) WriteAPIBlocking(org string, bucket string) api.WriteAPIBlocking {
	ret := _m.Called(org, bucket)

	var r0 api.WriteAPIBlocking
	if rf, ok := ret.Get(0).(func(string, string) api.WriteAPIBlocking); ok {
		r0 = rf(org, bucket)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.WriteAPIBlocking)
		}
	}

	return r0
}

// Client_WriteAPIBlocking_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteAPIBlocking'
type Client_WriteAPIBlocking_Call struct {
	*mock.Call
}

// WriteAPIBlocking is a helper method to define mock.On call
//  - org string
//  - bucket string
func (_e *Client_Expecter) WriteAPIBlocking(org interface{}, bucket interface{}) *Client_WriteAPIBlocking_Call {
	return &Client_WriteAPIBlocking_Call{Call: _e.mock.On("WriteAPIBlocking", org, bucket)}
}

func (_c *Client_WriteAPIBlocking_Call) Run(run func(org string, bucket string)) *Client_WriteAPIBlocking_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Client_WriteAPIBlocking_Call) Return(_a0 api.WriteAPIBlocking) *Client_WriteAPIBlocking_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
