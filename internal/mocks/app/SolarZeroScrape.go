// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	currentdata "github.com/mwinters-stuff/solar-zero-scrape-golang/app/currentdata"
	daydata "github.com/mwinters-stuff/solar-zero-scrape-golang/app/daydata"

	mock "github.com/stretchr/testify/mock"

	monthdata "github.com/mwinters-stuff/solar-zero-scrape-golang/app/monthdata"

	yeardata "github.com/mwinters-stuff/solar-zero-scrape-golang/app/yeardata"
)

// SolarZeroScrape is an autogenerated mock type for the SolarZeroScrape type
type SolarZeroScrape struct {
	mock.Mock
}

type SolarZeroScrape_Expecter struct {
	mock *mock.Mock
}

func (_m *SolarZeroScrape) EXPECT() *SolarZeroScrape_Expecter {
	return &SolarZeroScrape_Expecter{mock: &_m.Mock}
}

// AuthenticateFully provides a mock function with given fields:
func (_m *SolarZeroScrape) AuthenticateFully() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SolarZeroScrape_AuthenticateFully_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AuthenticateFully'
type SolarZeroScrape_AuthenticateFully_Call struct {
	*mock.Call
}

// AuthenticateFully is a helper method to define mock.On call
func (_e *SolarZeroScrape_Expecter) AuthenticateFully() *SolarZeroScrape_AuthenticateFully_Call {
	return &SolarZeroScrape_AuthenticateFully_Call{Call: _e.mock.On("AuthenticateFully")}
}

func (_c *SolarZeroScrape_AuthenticateFully_Call) Run(run func()) *SolarZeroScrape_AuthenticateFully_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SolarZeroScrape_AuthenticateFully_Call) Return(_a0 bool) *SolarZeroScrape_AuthenticateFully_Call {
	_c.Call.Return(_a0)
	return _c
}

// CurrentData provides a mock function with given fields:
func (_m *SolarZeroScrape) CurrentData() currentdata.CurrentData {
	ret := _m.Called()

	var r0 currentdata.CurrentData
	if rf, ok := ret.Get(0).(func() currentdata.CurrentData); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(currentdata.CurrentData)
	}

	return r0
}

// SolarZeroScrape_CurrentData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CurrentData'
type SolarZeroScrape_CurrentData_Call struct {
	*mock.Call
}

// CurrentData is a helper method to define mock.On call
func (_e *SolarZeroScrape_Expecter) CurrentData() *SolarZeroScrape_CurrentData_Call {
	return &SolarZeroScrape_CurrentData_Call{Call: _e.mock.On("CurrentData")}
}

func (_c *SolarZeroScrape_CurrentData_Call) Run(run func()) *SolarZeroScrape_CurrentData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SolarZeroScrape_CurrentData_Call) Return(_a0 currentdata.CurrentData) *SolarZeroScrape_CurrentData_Call {
	_c.Call.Return(_a0)
	return _c
}

// DayData provides a mock function with given fields:
func (_m *SolarZeroScrape) DayData() daydata.DayData {
	ret := _m.Called()

	var r0 daydata.DayData
	if rf, ok := ret.Get(0).(func() daydata.DayData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(daydata.DayData)
		}
	}

	return r0
}

// SolarZeroScrape_DayData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DayData'
type SolarZeroScrape_DayData_Call struct {
	*mock.Call
}

// DayData is a helper method to define mock.On call
func (_e *SolarZeroScrape_Expecter) DayData() *SolarZeroScrape_DayData_Call {
	return &SolarZeroScrape_DayData_Call{Call: _e.mock.On("DayData")}
}

func (_c *SolarZeroScrape_DayData_Call) Run(run func()) *SolarZeroScrape_DayData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SolarZeroScrape_DayData_Call) Return(_a0 daydata.DayData) *SolarZeroScrape_DayData_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetData provides a mock function with given fields:
func (_m *SolarZeroScrape) GetData() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SolarZeroScrape_GetData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetData'
type SolarZeroScrape_GetData_Call struct {
	*mock.Call
}

// GetData is a helper method to define mock.On call
func (_e *SolarZeroScrape_Expecter) GetData() *SolarZeroScrape_GetData_Call {
	return &SolarZeroScrape_GetData_Call{Call: _e.mock.On("GetData")}
}

func (_c *SolarZeroScrape_GetData_Call) Run(run func()) *SolarZeroScrape_GetData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SolarZeroScrape_GetData_Call) Return(_a0 bool) *SolarZeroScrape_GetData_Call {
	_c.Call.Return(_a0)
	return _c
}

// MonthData provides a mock function with given fields:
func (_m *SolarZeroScrape) MonthData() monthdata.MonthData {
	ret := _m.Called()

	var r0 monthdata.MonthData
	if rf, ok := ret.Get(0).(func() monthdata.MonthData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(monthdata.MonthData)
		}
	}

	return r0
}

// SolarZeroScrape_MonthData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MonthData'
type SolarZeroScrape_MonthData_Call struct {
	*mock.Call
}

// MonthData is a helper method to define mock.On call
func (_e *SolarZeroScrape_Expecter) MonthData() *SolarZeroScrape_MonthData_Call {
	return &SolarZeroScrape_MonthData_Call{Call: _e.mock.On("MonthData")}
}

func (_c *SolarZeroScrape_MonthData_Call) Run(run func()) *SolarZeroScrape_MonthData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SolarZeroScrape_MonthData_Call) Return(_a0 monthdata.MonthData) *SolarZeroScrape_MonthData_Call {
	_c.Call.Return(_a0)
	return _c
}

// YearData provides a mock function with given fields:
func (_m *SolarZeroScrape) YearData() yeardata.YearData {
	ret := _m.Called()

	var r0 yeardata.YearData
	if rf, ok := ret.Get(0).(func() yeardata.YearData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(yeardata.YearData)
		}
	}

	return r0
}

// SolarZeroScrape_YearData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'YearData'
type SolarZeroScrape_YearData_Call struct {
	*mock.Call
}

// YearData is a helper method to define mock.On call
func (_e *SolarZeroScrape_Expecter) YearData() *SolarZeroScrape_YearData_Call {
	return &SolarZeroScrape_YearData_Call{Call: _e.mock.On("YearData")}
}

func (_c *SolarZeroScrape_YearData_Call) Run(run func()) *SolarZeroScrape_YearData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SolarZeroScrape_YearData_Call) Return(_a0 yeardata.YearData) *SolarZeroScrape_YearData_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewSolarZeroScrape interface {
	mock.TestingT
	Cleanup(func())
}

// NewSolarZeroScrape creates a new instance of SolarZeroScrape. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSolarZeroScrape(t mockConstructorTestingTNewSolarZeroScrape) *SolarZeroScrape {
	mock := &SolarZeroScrape{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
