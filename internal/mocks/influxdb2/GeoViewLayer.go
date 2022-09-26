// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// GeoViewLayer is an autogenerated mock type for the GeoViewLayer type
type GeoViewLayer struct {
	mock.Mock
}

type GeoViewLayer_Expecter struct {
	mock *mock.Mock
}

func (_m *GeoViewLayer) EXPECT() *GeoViewLayer_Expecter {
	return &GeoViewLayer_Expecter{mock: &_m.Mock}
}

type mockConstructorTestingTNewGeoViewLayer interface {
	mock.TestingT
	Cleanup(func())
}

// NewGeoViewLayer creates a new instance of GeoViewLayer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGeoViewLayer(t mockConstructorTestingTNewGeoViewLayer) *GeoViewLayer {
	mock := &GeoViewLayer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}