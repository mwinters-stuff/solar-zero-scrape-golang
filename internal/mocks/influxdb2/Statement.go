// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Statement is an autogenerated mock type for the Statement type
type Statement struct {
	mock.Mock
}

type Statement_Expecter struct {
	mock *mock.Mock
}

func (_m *Statement) EXPECT() *Statement_Expecter {
	return &Statement_Expecter{mock: &_m.Mock}
}

type mockConstructorTestingTNewStatement interface {
	mock.TestingT
	Cleanup(func())
}

// NewStatement creates a new instance of Statement. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStatement(t mockConstructorTestingTNewStatement) *Statement {
	mock := &Statement{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}