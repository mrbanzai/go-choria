// Code generated by MockGen. DO NOT EDIT.
// Source: notifications.go

// Package machine is a generated GoMock package.
package machine

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockInfoSource is a mock of InfoSource interface
type MockInfoSource struct {
	ctrl     *gomock.Controller
	recorder *MockInfoSourceMockRecorder
}

// MockInfoSourceMockRecorder is the mock recorder for MockInfoSource
type MockInfoSourceMockRecorder struct {
	mock *MockInfoSource
}

// NewMockInfoSource creates a new mock instance
func NewMockInfoSource(ctrl *gomock.Controller) *MockInfoSource {
	mock := &MockInfoSource{ctrl: ctrl}
	mock.recorder = &MockInfoSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInfoSource) EXPECT() *MockInfoSourceMockRecorder {
	return m.recorder
}

// Identity mocks base method
func (m *MockInfoSource) Identity() string {
	ret := m.ctrl.Call(m, "Identity")
	ret0, _ := ret[0].(string)
	return ret0
}

// Identity indicates an expected call of Identity
func (mr *MockInfoSourceMockRecorder) Identity() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identity", reflect.TypeOf((*MockInfoSource)(nil).Identity))
}

// Version mocks base method
func (m *MockInfoSource) Version() string {
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

// Version indicates an expected call of Version
func (mr *MockInfoSourceMockRecorder) Version() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockInfoSource)(nil).Version))
}

// Name mocks base method
func (m *MockInfoSource) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockInfoSourceMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockInfoSource)(nil).Name))
}

// State mocks base method
func (m *MockInfoSource) State() string {
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(string)
	return ret0
}

// State indicates an expected call of State
func (mr *MockInfoSourceMockRecorder) State() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockInfoSource)(nil).State))
}

// MockWatcherStateNotification is a mock of WatcherStateNotification interface
type MockWatcherStateNotification struct {
	ctrl     *gomock.Controller
	recorder *MockWatcherStateNotificationMockRecorder
}

// MockWatcherStateNotificationMockRecorder is the mock recorder for MockWatcherStateNotification
type MockWatcherStateNotificationMockRecorder struct {
	mock *MockWatcherStateNotification
}

// NewMockWatcherStateNotification creates a new mock instance
func NewMockWatcherStateNotification(ctrl *gomock.Controller) *MockWatcherStateNotification {
	mock := &MockWatcherStateNotification{ctrl: ctrl}
	mock.recorder = &MockWatcherStateNotificationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWatcherStateNotification) EXPECT() *MockWatcherStateNotificationMockRecorder {
	return m.recorder
}

// JSON mocks base method
func (m *MockWatcherStateNotification) JSON() ([]byte, error) {
	ret := m.ctrl.Call(m, "JSON")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JSON indicates an expected call of JSON
func (mr *MockWatcherStateNotificationMockRecorder) JSON() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSON", reflect.TypeOf((*MockWatcherStateNotification)(nil).JSON))
}

// String mocks base method
func (m *MockWatcherStateNotification) String() string {
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockWatcherStateNotificationMockRecorder) String() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockWatcherStateNotification)(nil).String))
}

// MockNotificationService is a mock of NotificationService interface
type MockNotificationService struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationServiceMockRecorder
}

// MockNotificationServiceMockRecorder is the mock recorder for MockNotificationService
type MockNotificationServiceMockRecorder struct {
	mock *MockNotificationService
}

// NewMockNotificationService creates a new mock instance
func NewMockNotificationService(ctrl *gomock.Controller) *MockNotificationService {
	mock := &MockNotificationService{ctrl: ctrl}
	mock.recorder = &MockNotificationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNotificationService) EXPECT() *MockNotificationServiceMockRecorder {
	return m.recorder
}

// NotifyPostTransition mocks base method
func (m *MockNotificationService) NotifyPostTransition(t *TransitionNotification) error {
	ret := m.ctrl.Call(m, "NotifyPostTransition", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyPostTransition indicates an expected call of NotifyPostTransition
func (mr *MockNotificationServiceMockRecorder) NotifyPostTransition(t interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyPostTransition", reflect.TypeOf((*MockNotificationService)(nil).NotifyPostTransition), t)
}

// NotifyWatcherState mocks base method
func (m *MockNotificationService) NotifyWatcherState(watcher string, state WatcherStateNotification) error {
	ret := m.ctrl.Call(m, "NotifyWatcherState", watcher, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyWatcherState indicates an expected call of NotifyWatcherState
func (mr *MockNotificationServiceMockRecorder) NotifyWatcherState(watcher, state interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyWatcherState", reflect.TypeOf((*MockNotificationService)(nil).NotifyWatcherState), watcher, state)
}

// Debugf mocks base method
func (m *MockNotificationService) Debugf(machine InfoSource, watcher, format string, args ...interface{}) {
	varargs := []interface{}{machine, watcher, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf
func (mr *MockNotificationServiceMockRecorder) Debugf(machine, watcher, format interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{machine, watcher, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockNotificationService)(nil).Debugf), varargs...)
}

// Infof mocks base method
func (m *MockNotificationService) Infof(machine InfoSource, watcher, format string, args ...interface{}) {
	varargs := []interface{}{machine, watcher, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof
func (mr *MockNotificationServiceMockRecorder) Infof(machine, watcher, format interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{machine, watcher, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockNotificationService)(nil).Infof), varargs...)
}

// Warnf mocks base method
func (m *MockNotificationService) Warnf(machine InfoSource, watcher, format string, args ...interface{}) {
	varargs := []interface{}{machine, watcher, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf
func (mr *MockNotificationServiceMockRecorder) Warnf(machine, watcher, format interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{machine, watcher, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*MockNotificationService)(nil).Warnf), varargs...)
}

// Errorf mocks base method
func (m *MockNotificationService) Errorf(machine InfoSource, watcher, format string, args ...interface{}) {
	varargs := []interface{}{machine, watcher, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf
func (mr *MockNotificationServiceMockRecorder) Errorf(machine, watcher, format interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{machine, watcher, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockNotificationService)(nil).Errorf), varargs...)
}