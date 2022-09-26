// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

type Logger_Expecter struct {
	mock *mock.Mock
}

func (_m *Logger) EXPECT() *Logger_Expecter {
	return &Logger_Expecter{mock: &_m.Mock}
}

// Debug provides a mock function with given fields: msg
func (_m *Logger) Debug(msg string) {
	_m.Called(msg)
}

// Logger_Debug_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Debug'
type Logger_Debug_Call struct {
	*mock.Call
}

// Debug is a helper method to define mock.On call
//  - msg string
func (_e *Logger_Expecter) Debug(msg interface{}) *Logger_Debug_Call {
	return &Logger_Debug_Call{Call: _e.mock.On("Debug", msg)}
}

func (_c *Logger_Debug_Call) Run(run func(msg string)) *Logger_Debug_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Logger_Debug_Call) Return() *Logger_Debug_Call {
	_c.Call.Return()
	return _c
}

// Debugf provides a mock function with given fields: format, v
func (_m *Logger) Debugf(format string, v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Logger_Debugf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Debugf'
type Logger_Debugf_Call struct {
	*mock.Call
}

// Debugf is a helper method to define mock.On call
//  - format string
//  - v ...interface{}
func (_e *Logger_Expecter) Debugf(format interface{}, v ...interface{}) *Logger_Debugf_Call {
	return &Logger_Debugf_Call{Call: _e.mock.On("Debugf",
		append([]interface{}{format}, v...)...)}
}

func (_c *Logger_Debugf_Call) Run(run func(format string, v ...interface{})) *Logger_Debugf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Debugf_Call) Return() *Logger_Debugf_Call {
	_c.Call.Return()
	return _c
}

// Error provides a mock function with given fields: msg
func (_m *Logger) Error(msg string) {
	_m.Called(msg)
}

// Logger_Error_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Error'
type Logger_Error_Call struct {
	*mock.Call
}

// Error is a helper method to define mock.On call
//  - msg string
func (_e *Logger_Expecter) Error(msg interface{}) *Logger_Error_Call {
	return &Logger_Error_Call{Call: _e.mock.On("Error", msg)}
}

func (_c *Logger_Error_Call) Run(run func(msg string)) *Logger_Error_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Logger_Error_Call) Return() *Logger_Error_Call {
	_c.Call.Return()
	return _c
}

// Errorf provides a mock function with given fields: format, v
func (_m *Logger) Errorf(format string, v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Logger_Errorf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Errorf'
type Logger_Errorf_Call struct {
	*mock.Call
}

// Errorf is a helper method to define mock.On call
//  - format string
//  - v ...interface{}
func (_e *Logger_Expecter) Errorf(format interface{}, v ...interface{}) *Logger_Errorf_Call {
	return &Logger_Errorf_Call{Call: _e.mock.On("Errorf",
		append([]interface{}{format}, v...)...)}
}

func (_c *Logger_Errorf_Call) Run(run func(format string, v ...interface{})) *Logger_Errorf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Errorf_Call) Return() *Logger_Errorf_Call {
	_c.Call.Return()
	return _c
}

// Info provides a mock function with given fields: msg
func (_m *Logger) Info(msg string) {
	_m.Called(msg)
}

// Logger_Info_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Info'
type Logger_Info_Call struct {
	*mock.Call
}

// Info is a helper method to define mock.On call
//  - msg string
func (_e *Logger_Expecter) Info(msg interface{}) *Logger_Info_Call {
	return &Logger_Info_Call{Call: _e.mock.On("Info", msg)}
}

func (_c *Logger_Info_Call) Run(run func(msg string)) *Logger_Info_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Logger_Info_Call) Return() *Logger_Info_Call {
	_c.Call.Return()
	return _c
}

// Infof provides a mock function with given fields: format, v
func (_m *Logger) Infof(format string, v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Logger_Infof_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Infof'
type Logger_Infof_Call struct {
	*mock.Call
}

// Infof is a helper method to define mock.On call
//  - format string
//  - v ...interface{}
func (_e *Logger_Expecter) Infof(format interface{}, v ...interface{}) *Logger_Infof_Call {
	return &Logger_Infof_Call{Call: _e.mock.On("Infof",
		append([]interface{}{format}, v...)...)}
}

func (_c *Logger_Infof_Call) Run(run func(format string, v ...interface{})) *Logger_Infof_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Infof_Call) Return() *Logger_Infof_Call {
	_c.Call.Return()
	return _c
}

// LogLevel provides a mock function with given fields:
func (_m *Logger) LogLevel() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// Logger_LogLevel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LogLevel'
type Logger_LogLevel_Call struct {
	*mock.Call
}

// LogLevel is a helper method to define mock.On call
func (_e *Logger_Expecter) LogLevel() *Logger_LogLevel_Call {
	return &Logger_LogLevel_Call{Call: _e.mock.On("LogLevel")}
}

func (_c *Logger_LogLevel_Call) Run(run func()) *Logger_LogLevel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Logger_LogLevel_Call) Return(_a0 uint) *Logger_LogLevel_Call {
	_c.Call.Return(_a0)
	return _c
}

// SetLogLevel provides a mock function with given fields: logLevel
func (_m *Logger) SetLogLevel(logLevel uint) {
	_m.Called(logLevel)
}

// Logger_SetLogLevel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetLogLevel'
type Logger_SetLogLevel_Call struct {
	*mock.Call
}

// SetLogLevel is a helper method to define mock.On call
//  - logLevel uint
func (_e *Logger_Expecter) SetLogLevel(logLevel interface{}) *Logger_SetLogLevel_Call {
	return &Logger_SetLogLevel_Call{Call: _e.mock.On("SetLogLevel", logLevel)}
}

func (_c *Logger_SetLogLevel_Call) Run(run func(logLevel uint)) *Logger_SetLogLevel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *Logger_SetLogLevel_Call) Return() *Logger_SetLogLevel_Call {
	_c.Call.Return()
	return _c
}

// SetPrefix provides a mock function with given fields: prefix
func (_m *Logger) SetPrefix(prefix string) {
	_m.Called(prefix)
}

// Logger_SetPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPrefix'
type Logger_SetPrefix_Call struct {
	*mock.Call
}

// SetPrefix is a helper method to define mock.On call
//  - prefix string
func (_e *Logger_Expecter) SetPrefix(prefix interface{}) *Logger_SetPrefix_Call {
	return &Logger_SetPrefix_Call{Call: _e.mock.On("SetPrefix", prefix)}
}

func (_c *Logger_SetPrefix_Call) Run(run func(prefix string)) *Logger_SetPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Logger_SetPrefix_Call) Return() *Logger_SetPrefix_Call {
	_c.Call.Return()
	return _c
}

// Warn provides a mock function with given fields: msg
func (_m *Logger) Warn(msg string) {
	_m.Called(msg)
}

// Logger_Warn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Warn'
type Logger_Warn_Call struct {
	*mock.Call
}

// Warn is a helper method to define mock.On call
//  - msg string
func (_e *Logger_Expecter) Warn(msg interface{}) *Logger_Warn_Call {
	return &Logger_Warn_Call{Call: _e.mock.On("Warn", msg)}
}

func (_c *Logger_Warn_Call) Run(run func(msg string)) *Logger_Warn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Logger_Warn_Call) Return() *Logger_Warn_Call {
	_c.Call.Return()
	return _c
}

// Warnf provides a mock function with given fields: format, v
func (_m *Logger) Warnf(format string, v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Logger_Warnf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Warnf'
type Logger_Warnf_Call struct {
	*mock.Call
}

// Warnf is a helper method to define mock.On call
//  - format string
//  - v ...interface{}
func (_e *Logger_Expecter) Warnf(format interface{}, v ...interface{}) *Logger_Warnf_Call {
	return &Logger_Warnf_Call{Call: _e.mock.On("Warnf",
		append([]interface{}{format}, v...)...)}
}

func (_c *Logger_Warnf_Call) Run(run func(format string, v ...interface{})) *Logger_Warnf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Warnf_Call) Return() *Logger_Warnf_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewLogger interface {
	mock.TestingT
	Cleanup(func())
}

// NewLogger creates a new instance of Logger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLogger(t mockConstructorTestingTNewLogger) *Logger {
	mock := &Logger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}