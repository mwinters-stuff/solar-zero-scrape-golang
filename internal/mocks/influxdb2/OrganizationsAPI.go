// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	api "github.com/influxdata/influxdb-client-go/v2/api"

	domain "github.com/influxdata/influxdb-client-go/v2/domain"

	mock "github.com/stretchr/testify/mock"
)

// OrganizationsAPI is an autogenerated mock type for the OrganizationsAPI type
type OrganizationsAPI struct {
	mock.Mock
}

type OrganizationsAPI_Expecter struct {
	mock *mock.Mock
}

func (_m *OrganizationsAPI) EXPECT() *OrganizationsAPI_Expecter {
	return &OrganizationsAPI_Expecter{mock: &_m.Mock}
}

// AddMember provides a mock function with given fields: ctx, org, user
func (_m *OrganizationsAPI) AddMember(ctx context.Context, org *domain.Organization, user *domain.User) (*domain.ResourceMember, error) {
	ret := _m.Called(ctx, org, user)

	var r0 *domain.ResourceMember
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Organization, *domain.User) *domain.ResourceMember); ok {
		r0 = rf(ctx, org, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ResourceMember)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Organization, *domain.User) error); ok {
		r1 = rf(ctx, org, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrganizationsAPI_AddMember_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddMember'
type OrganizationsAPI_AddMember_Call struct {
	*mock.Call
}

// AddMember is a helper method to define mock.On call
//  - ctx context.Context
//  - org *domain.Organization
//  - user *domain.User
func (_e *OrganizationsAPI_Expecter) AddMember(ctx interface{}, org interface{}, user interface{}) *OrganizationsAPI_AddMember_Call {
	return &OrganizationsAPI_AddMember_Call{Call: _e.mock.On("AddMember", ctx, org, user)}
}

func (_c *OrganizationsAPI_AddMember_Call) Run(run func(ctx context.Context, org *domain.Organization, user *domain.User)) *OrganizationsAPI_AddMember_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Organization), args[2].(*domain.User))
	})
	return _c
}

func (_c *OrganizationsAPI_AddMember_Call) Return(_a0 *domain.ResourceMember, _a1 error) *OrganizationsAPI_AddMember_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// AddMemberWithID provides a mock function with given fields: ctx, orgID, memberID
func (_m *OrganizationsAPI) AddMemberWithID(ctx context.Context, orgID string, memberID string) (*domain.ResourceMember, error) {
	ret := _m.Called(ctx, orgID, memberID)

	var r0 *domain.ResourceMember
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *domain.ResourceMember); ok {
		r0 = rf(ctx, orgID, memberID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ResourceMember)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, orgID, memberID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrganizationsAPI_AddMemberWithID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddMemberWithID'
type OrganizationsAPI_AddMemberWithID_Call struct {
	*mock.Call
}

// AddMemberWithID is a helper method to define mock.On call
//  - ctx context.Context
//  - orgID string
//  - memberID string
func (_e *OrganizationsAPI_Expecter) AddMemberWithID(ctx interface{}, orgID interface{}, memberID interface{}) *OrganizationsAPI_AddMemberWithID_Call {
	return &OrganizationsAPI_AddMemberWithID_Call{Call: _e.mock.On("AddMemberWithID", ctx, orgID, memberID)}
}

func (_c *OrganizationsAPI_AddMemberWithID_Call) Run(run func(ctx context.Context, orgID string, memberID string)) *OrganizationsAPI_AddMemberWithID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *OrganizationsAPI_AddMemberWithID_Call) Return(_a0 *domain.ResourceMember, _a1 error) *OrganizationsAPI_AddMemberWithID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// AddOwner provides a mock function with given fields: ctx, org, user
func (_m *OrganizationsAPI) AddOwner(ctx context.Context, org *domain.Organization, user *domain.User) (*domain.ResourceOwner, error) {
	ret := _m.Called(ctx, org, user)

	var r0 *domain.ResourceOwner
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Organization, *domain.User) *domain.ResourceOwner); ok {
		r0 = rf(ctx, org, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ResourceOwner)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Organization, *domain.User) error); ok {
		r1 = rf(ctx, org, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrganizationsAPI_AddOwner_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddOwner'
type OrganizationsAPI_AddOwner_Call struct {
	*mock.Call
}

// AddOwner is a helper method to define mock.On call
//  - ctx context.Context
//  - org *domain.Organization
//  - user *domain.User
func (_e *OrganizationsAPI_Expecter) AddOwner(ctx interface{}, org interface{}, user interface{}) *OrganizationsAPI_AddOwner_Call {
	return &OrganizationsAPI_AddOwner_Call{Call: _e.mock.On("AddOwner", ctx, org, user)}
}

func (_c *OrganizationsAPI_AddOwner_Call) Run(run func(ctx context.Context, org *domain.Organization, user *domain.User)) *OrganizationsAPI_AddOwner_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Organization), args[2].(*domain.User))
	})
	return _c
}

func (_c *OrganizationsAPI_AddOwner_Call) Return(_a0 *domain.ResourceOwner, _a1 error) *OrganizationsAPI_AddOwner_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// AddOwnerWithID provides a mock function with given fields: ctx, orgID, memberID
func (_m *OrganizationsAPI) AddOwnerWithID(ctx context.Context, orgID string, memberID string) (*domain.ResourceOwner, error) {
	ret := _m.Called(ctx, orgID, memberID)

	var r0 *domain.ResourceOwner
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *domain.ResourceOwner); ok {
		r0 = rf(ctx, orgID, memberID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ResourceOwner)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, orgID, memberID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrganizationsAPI_AddOwnerWithID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddOwnerWithID'
type OrganizationsAPI_AddOwnerWithID_Call struct {
	*mock.Call
}

// AddOwnerWithID is a helper method to define mock.On call
//  - ctx context.Context
//  - orgID string
//  - memberID string
func (_e *OrganizationsAPI_Expecter) AddOwnerWithID(ctx interface{}, orgID interface{}, memberID interface{}) *OrganizationsAPI_AddOwnerWithID_Call {
	return &OrganizationsAPI_AddOwnerWithID_Call{Call: _e.mock.On("AddOwnerWithID", ctx, orgID, memberID)}
}

func (_c *OrganizationsAPI_AddOwnerWithID_Call) Run(run func(ctx context.Context, orgID string, memberID string)) *OrganizationsAPI_AddOwnerWithID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *OrganizationsAPI_AddOwnerWithID_Call) Return(_a0 *domain.ResourceOwner, _a1 error) *OrganizationsAPI_AddOwnerWithID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// CreateOrganization provides a mock function with given fields: ctx, org
func (_m *OrganizationsAPI) CreateOrganization(ctx context.Context, org *domain.Organization) (*domain.Organization, error) {
	ret := _m.Called(ctx, org)

	var r0 *domain.Organization
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Organization) *domain.Organization); ok {
		r0 = rf(ctx, org)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Organization) error); ok {
		r1 = rf(ctx, org)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrganizationsAPI_CreateOrganization_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrganization'
type OrganizationsAPI_CreateOrganization_Call struct {
	*mock.Call
}

// CreateOrganization is a helper method to define mock.On call
//  - ctx context.Context
//  - org *domain.Organization
func (_e *OrganizationsAPI_Expecter) CreateOrganization(ctx interface{}, org interface{}) *OrganizationsAPI_CreateOrganization_Call {
	return &OrganizationsAPI_CreateOrganization_Call{Call: _e.mock.On("CreateOrganization", ctx, org)}
}

func (_c *OrganizationsAPI_CreateOrganization_Call) Run(run func(ctx context.Context, org *domain.Organization)) *OrganizationsAPI_CreateOrganization_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Organization))
	})
	return _c
}

func (_c *OrganizationsAPI_CreateOrganization_Call) Return(_a0 *domain.Organization, _a1 error) *OrganizationsAPI_CreateOrganization_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// CreateOrganizationWithName provides a mock function with given fields: ctx, orgName
func (_m *OrganizationsAPI) CreateOrganizationWithName(ctx context.Context, orgName string) (*domain.Organization, error) {
	ret := _m.Called(ctx, orgName)

	var r0 *domain.Organization
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Organization); ok {
		r0 = rf(ctx, orgName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, orgName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrganizationsAPI_CreateOrganizationWithName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrganizationWithName'
type OrganizationsAPI_CreateOrganizationWithName_Call struct {
	*mock.Call
}

// CreateOrganizationWithName is a helper method to define mock.On call
//  - ctx context.Context
//  - orgName string
func (_e *OrganizationsAPI_Expecter) CreateOrganizationWithName(ctx interface{}, orgName interface{}) *OrganizationsAPI_CreateOrganizationWithName_Call {
	return &OrganizationsAPI_CreateOrganizationWithName_Call{Call: _e.mock.On("CreateOrganizationWithName", ctx, orgName)}
}

func (_c *OrganizationsAPI_CreateOrganizationWithName_Call) Run(run func(ctx context.Context, orgName string)) *OrganizationsAPI_CreateOrganizationWithName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *OrganizationsAPI_CreateOrganizationWithName_Call) Return(_a0 *domain.Organization, _a1 error) *OrganizationsAPI_CreateOrganizationWithName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// DeleteOrganization provides a mock function with given fields: ctx, org
func (_m *OrganizationsAPI) DeleteOrganization(ctx context.Context, org *domain.Organization) error {
	ret := _m.Called(ctx, org)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Organization) error); ok {
		r0 = rf(ctx, org)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrganizationsAPI_DeleteOrganization_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteOrganization'
type OrganizationsAPI_DeleteOrganization_Call struct {
	*mock.Call
}

// DeleteOrganization is a helper method to define mock.On call
//  - ctx context.Context
//  - org *domain.Organization
func (_e *OrganizationsAPI_Expecter) DeleteOrganization(ctx interface{}, org interface{}) *OrganizationsAPI_DeleteOrganization_Call {
	return &OrganizationsAPI_DeleteOrganization_Call{Call: _e.mock.On("DeleteOrganization", ctx, org)}
}

func (_c *OrganizationsAPI_DeleteOrganization_Call) Run(run func(ctx context.Context, org *domain.Organization)) *OrganizationsAPI_DeleteOrganization_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Organization))
	})
	return _c
}

func (_c *OrganizationsAPI_DeleteOrganization_Call) Return(_a0 error) *OrganizationsAPI_DeleteOrganization_Call {
	_c.Call.Return(_a0)
	return _c
}

// DeleteOrganizationWithID provides a mock function with given fields: ctx, orgID
func (_m *OrganizationsAPI) DeleteOrganizationWithID(ctx context.Context, orgID string) error {
	ret := _m.Called(ctx, orgID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, orgID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrganizationsAPI_DeleteOrganizationWithID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteOrganizationWithID'
type OrganizationsAPI_DeleteOrganizationWithID_Call struct {
	*mock.Call
}

// DeleteOrganizationWithID is a helper method to define mock.On call
//  - ctx context.Context
//  - orgID string
func (_e *OrganizationsAPI_Expecter) DeleteOrganizationWithID(ctx interface{}, orgID interface{}) *OrganizationsAPI_DeleteOrganizationWithID_Call {
	return &OrganizationsAPI_DeleteOrganizationWithID_Call{Call: _e.mock.On("DeleteOrganizationWithID", ctx, orgID)}
}

func (_c *OrganizationsAPI_DeleteOrganizationWithID_Call) Run(run func(ctx context.Context, orgID string)) *OrganizationsAPI_DeleteOrganizationWithID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *OrganizationsAPI_DeleteOrganizationWithID_Call) Return(_a0 error) *OrganizationsAPI_DeleteOrganizationWithID_Call {
	_c.Call.Return(_a0)
	return _c
}

// FindOrganizationByID provides a mock function with given fields: ctx, orgID
func (_m *OrganizationsAPI) FindOrganizationByID(ctx context.Context, orgID string) (*domain.Organization, error) {
	ret := _m.Called(ctx, orgID)

	var r0 *domain.Organization
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Organization); ok {
		r0 = rf(ctx, orgID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, orgID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrganizationsAPI_FindOrganizationByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOrganizationByID'
type OrganizationsAPI_FindOrganizationByID_Call struct {
	*mock.Call
}

// FindOrganizationByID is a helper method to define mock.On call
//  - ctx context.Context
//  - orgID string
func (_e *OrganizationsAPI_Expecter) FindOrganizationByID(ctx interface{}, orgID interface{}) *OrganizationsAPI_FindOrganizationByID_Call {
	return &OrganizationsAPI_FindOrganizationByID_Call{Call: _e.mock.On("FindOrganizationByID", ctx, orgID)}
}

func (_c *OrganizationsAPI_FindOrganizationByID_Call) Run(run func(ctx context.Context, orgID string)) *OrganizationsAPI_FindOrganizationByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *OrganizationsAPI_FindOrganizationByID_Call) Return(_a0 *domain.Organization, _a1 error) *OrganizationsAPI_FindOrganizationByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// FindOrganizationByName provides a mock function with given fields: ctx, orgName
func (_m *OrganizationsAPI) FindOrganizationByName(ctx context.Context, orgName string) (*domain.Organization, error) {
	ret := _m.Called(ctx, orgName)

	var r0 *domain.Organization
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Organization); ok {
		r0 = rf(ctx, orgName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, orgName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrganizationsAPI_FindOrganizationByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOrganizationByName'
type OrganizationsAPI_FindOrganizationByName_Call struct {
	*mock.Call
}

// FindOrganizationByName is a helper method to define mock.On call
//  - ctx context.Context
//  - orgName string
func (_e *OrganizationsAPI_Expecter) FindOrganizationByName(ctx interface{}, orgName interface{}) *OrganizationsAPI_FindOrganizationByName_Call {
	return &OrganizationsAPI_FindOrganizationByName_Call{Call: _e.mock.On("FindOrganizationByName", ctx, orgName)}
}

func (_c *OrganizationsAPI_FindOrganizationByName_Call) Run(run func(ctx context.Context, orgName string)) *OrganizationsAPI_FindOrganizationByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *OrganizationsAPI_FindOrganizationByName_Call) Return(_a0 *domain.Organization, _a1 error) *OrganizationsAPI_FindOrganizationByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// FindOrganizationsByUserID provides a mock function with given fields: ctx, userID, pagingOptions
func (_m *OrganizationsAPI) FindOrganizationsByUserID(ctx context.Context, userID string, pagingOptions ...api.PagingOption) (*[]domain.Organization, error) {
	_va := make([]interface{}, len(pagingOptions))
	for _i := range pagingOptions {
		_va[_i] = pagingOptions[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, userID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *[]domain.Organization
	if rf, ok := ret.Get(0).(func(context.Context, string, ...api.PagingOption) *[]domain.Organization); ok {
		r0 = rf(ctx, userID, pagingOptions...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...api.PagingOption) error); ok {
		r1 = rf(ctx, userID, pagingOptions...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrganizationsAPI_FindOrganizationsByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOrganizationsByUserID'
type OrganizationsAPI_FindOrganizationsByUserID_Call struct {
	*mock.Call
}

// FindOrganizationsByUserID is a helper method to define mock.On call
//  - ctx context.Context
//  - userID string
//  - pagingOptions ...api.PagingOption
func (_e *OrganizationsAPI_Expecter) FindOrganizationsByUserID(ctx interface{}, userID interface{}, pagingOptions ...interface{}) *OrganizationsAPI_FindOrganizationsByUserID_Call {
	return &OrganizationsAPI_FindOrganizationsByUserID_Call{Call: _e.mock.On("FindOrganizationsByUserID",
		append([]interface{}{ctx, userID}, pagingOptions...)...)}
}

func (_c *OrganizationsAPI_FindOrganizationsByUserID_Call) Run(run func(ctx context.Context, userID string, pagingOptions ...api.PagingOption)) *OrganizationsAPI_FindOrganizationsByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]api.PagingOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(api.PagingOption)
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *OrganizationsAPI_FindOrganizationsByUserID_Call) Return(_a0 *[]domain.Organization, _a1 error) *OrganizationsAPI_FindOrganizationsByUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetMembers provides a mock function with given fields: ctx, org
func (_m *OrganizationsAPI) GetMembers(ctx context.Context, org *domain.Organization) (*[]domain.ResourceMember, error) {
	ret := _m.Called(ctx, org)

	var r0 *[]domain.ResourceMember
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Organization) *[]domain.ResourceMember); ok {
		r0 = rf(ctx, org)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.ResourceMember)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Organization) error); ok {
		r1 = rf(ctx, org)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrganizationsAPI_GetMembers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMembers'
type OrganizationsAPI_GetMembers_Call struct {
	*mock.Call
}

// GetMembers is a helper method to define mock.On call
//  - ctx context.Context
//  - org *domain.Organization
func (_e *OrganizationsAPI_Expecter) GetMembers(ctx interface{}, org interface{}) *OrganizationsAPI_GetMembers_Call {
	return &OrganizationsAPI_GetMembers_Call{Call: _e.mock.On("GetMembers", ctx, org)}
}

func (_c *OrganizationsAPI_GetMembers_Call) Run(run func(ctx context.Context, org *domain.Organization)) *OrganizationsAPI_GetMembers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Organization))
	})
	return _c
}

func (_c *OrganizationsAPI_GetMembers_Call) Return(_a0 *[]domain.ResourceMember, _a1 error) *OrganizationsAPI_GetMembers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetMembersWithID provides a mock function with given fields: ctx, orgID
func (_m *OrganizationsAPI) GetMembersWithID(ctx context.Context, orgID string) (*[]domain.ResourceMember, error) {
	ret := _m.Called(ctx, orgID)

	var r0 *[]domain.ResourceMember
	if rf, ok := ret.Get(0).(func(context.Context, string) *[]domain.ResourceMember); ok {
		r0 = rf(ctx, orgID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.ResourceMember)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, orgID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrganizationsAPI_GetMembersWithID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMembersWithID'
type OrganizationsAPI_GetMembersWithID_Call struct {
	*mock.Call
}

// GetMembersWithID is a helper method to define mock.On call
//  - ctx context.Context
//  - orgID string
func (_e *OrganizationsAPI_Expecter) GetMembersWithID(ctx interface{}, orgID interface{}) *OrganizationsAPI_GetMembersWithID_Call {
	return &OrganizationsAPI_GetMembersWithID_Call{Call: _e.mock.On("GetMembersWithID", ctx, orgID)}
}

func (_c *OrganizationsAPI_GetMembersWithID_Call) Run(run func(ctx context.Context, orgID string)) *OrganizationsAPI_GetMembersWithID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *OrganizationsAPI_GetMembersWithID_Call) Return(_a0 *[]domain.ResourceMember, _a1 error) *OrganizationsAPI_GetMembersWithID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetOrganizations provides a mock function with given fields: ctx, pagingOptions
func (_m *OrganizationsAPI) GetOrganizations(ctx context.Context, pagingOptions ...api.PagingOption) (*[]domain.Organization, error) {
	_va := make([]interface{}, len(pagingOptions))
	for _i := range pagingOptions {
		_va[_i] = pagingOptions[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *[]domain.Organization
	if rf, ok := ret.Get(0).(func(context.Context, ...api.PagingOption) *[]domain.Organization); ok {
		r0 = rf(ctx, pagingOptions...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...api.PagingOption) error); ok {
		r1 = rf(ctx, pagingOptions...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrganizationsAPI_GetOrganizations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrganizations'
type OrganizationsAPI_GetOrganizations_Call struct {
	*mock.Call
}

// GetOrganizations is a helper method to define mock.On call
//  - ctx context.Context
//  - pagingOptions ...api.PagingOption
func (_e *OrganizationsAPI_Expecter) GetOrganizations(ctx interface{}, pagingOptions ...interface{}) *OrganizationsAPI_GetOrganizations_Call {
	return &OrganizationsAPI_GetOrganizations_Call{Call: _e.mock.On("GetOrganizations",
		append([]interface{}{ctx}, pagingOptions...)...)}
}

func (_c *OrganizationsAPI_GetOrganizations_Call) Run(run func(ctx context.Context, pagingOptions ...api.PagingOption)) *OrganizationsAPI_GetOrganizations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]api.PagingOption, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(api.PagingOption)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *OrganizationsAPI_GetOrganizations_Call) Return(_a0 *[]domain.Organization, _a1 error) *OrganizationsAPI_GetOrganizations_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetOwners provides a mock function with given fields: ctx, org
func (_m *OrganizationsAPI) GetOwners(ctx context.Context, org *domain.Organization) (*[]domain.ResourceOwner, error) {
	ret := _m.Called(ctx, org)

	var r0 *[]domain.ResourceOwner
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Organization) *[]domain.ResourceOwner); ok {
		r0 = rf(ctx, org)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.ResourceOwner)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Organization) error); ok {
		r1 = rf(ctx, org)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrganizationsAPI_GetOwners_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOwners'
type OrganizationsAPI_GetOwners_Call struct {
	*mock.Call
}

// GetOwners is a helper method to define mock.On call
//  - ctx context.Context
//  - org *domain.Organization
func (_e *OrganizationsAPI_Expecter) GetOwners(ctx interface{}, org interface{}) *OrganizationsAPI_GetOwners_Call {
	return &OrganizationsAPI_GetOwners_Call{Call: _e.mock.On("GetOwners", ctx, org)}
}

func (_c *OrganizationsAPI_GetOwners_Call) Run(run func(ctx context.Context, org *domain.Organization)) *OrganizationsAPI_GetOwners_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Organization))
	})
	return _c
}

func (_c *OrganizationsAPI_GetOwners_Call) Return(_a0 *[]domain.ResourceOwner, _a1 error) *OrganizationsAPI_GetOwners_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetOwnersWithID provides a mock function with given fields: ctx, orgID
func (_m *OrganizationsAPI) GetOwnersWithID(ctx context.Context, orgID string) (*[]domain.ResourceOwner, error) {
	ret := _m.Called(ctx, orgID)

	var r0 *[]domain.ResourceOwner
	if rf, ok := ret.Get(0).(func(context.Context, string) *[]domain.ResourceOwner); ok {
		r0 = rf(ctx, orgID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.ResourceOwner)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, orgID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrganizationsAPI_GetOwnersWithID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOwnersWithID'
type OrganizationsAPI_GetOwnersWithID_Call struct {
	*mock.Call
}

// GetOwnersWithID is a helper method to define mock.On call
//  - ctx context.Context
//  - orgID string
func (_e *OrganizationsAPI_Expecter) GetOwnersWithID(ctx interface{}, orgID interface{}) *OrganizationsAPI_GetOwnersWithID_Call {
	return &OrganizationsAPI_GetOwnersWithID_Call{Call: _e.mock.On("GetOwnersWithID", ctx, orgID)}
}

func (_c *OrganizationsAPI_GetOwnersWithID_Call) Run(run func(ctx context.Context, orgID string)) *OrganizationsAPI_GetOwnersWithID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *OrganizationsAPI_GetOwnersWithID_Call) Return(_a0 *[]domain.ResourceOwner, _a1 error) *OrganizationsAPI_GetOwnersWithID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// RemoveMember provides a mock function with given fields: ctx, org, user
func (_m *OrganizationsAPI) RemoveMember(ctx context.Context, org *domain.Organization, user *domain.User) error {
	ret := _m.Called(ctx, org, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Organization, *domain.User) error); ok {
		r0 = rf(ctx, org, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrganizationsAPI_RemoveMember_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveMember'
type OrganizationsAPI_RemoveMember_Call struct {
	*mock.Call
}

// RemoveMember is a helper method to define mock.On call
//  - ctx context.Context
//  - org *domain.Organization
//  - user *domain.User
func (_e *OrganizationsAPI_Expecter) RemoveMember(ctx interface{}, org interface{}, user interface{}) *OrganizationsAPI_RemoveMember_Call {
	return &OrganizationsAPI_RemoveMember_Call{Call: _e.mock.On("RemoveMember", ctx, org, user)}
}

func (_c *OrganizationsAPI_RemoveMember_Call) Run(run func(ctx context.Context, org *domain.Organization, user *domain.User)) *OrganizationsAPI_RemoveMember_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Organization), args[2].(*domain.User))
	})
	return _c
}

func (_c *OrganizationsAPI_RemoveMember_Call) Return(_a0 error) *OrganizationsAPI_RemoveMember_Call {
	_c.Call.Return(_a0)
	return _c
}

// RemoveMemberWithID provides a mock function with given fields: ctx, orgID, memberID
func (_m *OrganizationsAPI) RemoveMemberWithID(ctx context.Context, orgID string, memberID string) error {
	ret := _m.Called(ctx, orgID, memberID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, orgID, memberID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrganizationsAPI_RemoveMemberWithID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveMemberWithID'
type OrganizationsAPI_RemoveMemberWithID_Call struct {
	*mock.Call
}

// RemoveMemberWithID is a helper method to define mock.On call
//  - ctx context.Context
//  - orgID string
//  - memberID string
func (_e *OrganizationsAPI_Expecter) RemoveMemberWithID(ctx interface{}, orgID interface{}, memberID interface{}) *OrganizationsAPI_RemoveMemberWithID_Call {
	return &OrganizationsAPI_RemoveMemberWithID_Call{Call: _e.mock.On("RemoveMemberWithID", ctx, orgID, memberID)}
}

func (_c *OrganizationsAPI_RemoveMemberWithID_Call) Run(run func(ctx context.Context, orgID string, memberID string)) *OrganizationsAPI_RemoveMemberWithID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *OrganizationsAPI_RemoveMemberWithID_Call) Return(_a0 error) *OrganizationsAPI_RemoveMemberWithID_Call {
	_c.Call.Return(_a0)
	return _c
}

// RemoveOwner provides a mock function with given fields: ctx, org, user
func (_m *OrganizationsAPI) RemoveOwner(ctx context.Context, org *domain.Organization, user *domain.User) error {
	ret := _m.Called(ctx, org, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Organization, *domain.User) error); ok {
		r0 = rf(ctx, org, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrganizationsAPI_RemoveOwner_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveOwner'
type OrganizationsAPI_RemoveOwner_Call struct {
	*mock.Call
}

// RemoveOwner is a helper method to define mock.On call
//  - ctx context.Context
//  - org *domain.Organization
//  - user *domain.User
func (_e *OrganizationsAPI_Expecter) RemoveOwner(ctx interface{}, org interface{}, user interface{}) *OrganizationsAPI_RemoveOwner_Call {
	return &OrganizationsAPI_RemoveOwner_Call{Call: _e.mock.On("RemoveOwner", ctx, org, user)}
}

func (_c *OrganizationsAPI_RemoveOwner_Call) Run(run func(ctx context.Context, org *domain.Organization, user *domain.User)) *OrganizationsAPI_RemoveOwner_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Organization), args[2].(*domain.User))
	})
	return _c
}

func (_c *OrganizationsAPI_RemoveOwner_Call) Return(_a0 error) *OrganizationsAPI_RemoveOwner_Call {
	_c.Call.Return(_a0)
	return _c
}

// RemoveOwnerWithID provides a mock function with given fields: ctx, orgID, memberID
func (_m *OrganizationsAPI) RemoveOwnerWithID(ctx context.Context, orgID string, memberID string) error {
	ret := _m.Called(ctx, orgID, memberID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, orgID, memberID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrganizationsAPI_RemoveOwnerWithID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveOwnerWithID'
type OrganizationsAPI_RemoveOwnerWithID_Call struct {
	*mock.Call
}

// RemoveOwnerWithID is a helper method to define mock.On call
//  - ctx context.Context
//  - orgID string
//  - memberID string
func (_e *OrganizationsAPI_Expecter) RemoveOwnerWithID(ctx interface{}, orgID interface{}, memberID interface{}) *OrganizationsAPI_RemoveOwnerWithID_Call {
	return &OrganizationsAPI_RemoveOwnerWithID_Call{Call: _e.mock.On("RemoveOwnerWithID", ctx, orgID, memberID)}
}

func (_c *OrganizationsAPI_RemoveOwnerWithID_Call) Run(run func(ctx context.Context, orgID string, memberID string)) *OrganizationsAPI_RemoveOwnerWithID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *OrganizationsAPI_RemoveOwnerWithID_Call) Return(_a0 error) *OrganizationsAPI_RemoveOwnerWithID_Call {
	_c.Call.Return(_a0)
	return _c
}

// UpdateOrganization provides a mock function with given fields: ctx, org
func (_m *OrganizationsAPI) UpdateOrganization(ctx context.Context, org *domain.Organization) (*domain.Organization, error) {
	ret := _m.Called(ctx, org)

	var r0 *domain.Organization
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Organization) *domain.Organization); ok {
		r0 = rf(ctx, org)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Organization) error); ok {
		r1 = rf(ctx, org)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrganizationsAPI_UpdateOrganization_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateOrganization'
type OrganizationsAPI_UpdateOrganization_Call struct {
	*mock.Call
}

// UpdateOrganization is a helper method to define mock.On call
//  - ctx context.Context
//  - org *domain.Organization
func (_e *OrganizationsAPI_Expecter) UpdateOrganization(ctx interface{}, org interface{}) *OrganizationsAPI_UpdateOrganization_Call {
	return &OrganizationsAPI_UpdateOrganization_Call{Call: _e.mock.On("UpdateOrganization", ctx, org)}
}

func (_c *OrganizationsAPI_UpdateOrganization_Call) Run(run func(ctx context.Context, org *domain.Organization)) *OrganizationsAPI_UpdateOrganization_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Organization))
	})
	return _c
}

func (_c *OrganizationsAPI_UpdateOrganization_Call) Return(_a0 *domain.Organization, _a1 error) *OrganizationsAPI_UpdateOrganization_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewOrganizationsAPI interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrganizationsAPI creates a new instance of OrganizationsAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrganizationsAPI(t mockConstructorTestingTNewOrganizationsAPI) *OrganizationsAPI {
	mock := &OrganizationsAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
