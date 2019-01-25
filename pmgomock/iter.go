// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/luoji/gopath/src/github.com/gunsluo/pmgo/iter.go

// Package pmgomock is a generated GoMock package.
package pmgomock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIterManager is a mock of IterManager interface
type MockIterManager struct {
	ctrl     *gomock.Controller
	recorder *MockIterManagerMockRecorder
}

// MockIterManagerMockRecorder is the mock recorder for MockIterManager
type MockIterManagerMockRecorder struct {
	mock *MockIterManager
}

// NewMockIterManager creates a new mock instance
func NewMockIterManager(ctrl *gomock.Controller) *MockIterManager {
	mock := &MockIterManager{ctrl: ctrl}
	mock.recorder = &MockIterManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIterManager) EXPECT() *MockIterManagerMockRecorder {
	return m.recorder
}

// All mocks base method
func (m *MockIterManager) All(result interface{}) error {
	ret := m.ctrl.Call(m, "All", result)
	ret0, _ := ret[0].(error)
	return ret0
}

// All indicates an expected call of All
func (mr *MockIterManagerMockRecorder) All(result interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockIterManager)(nil).All), result)
}

// Close mocks base method
func (m *MockIterManager) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockIterManagerMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIterManager)(nil).Close))
}

// Done mocks base method
func (m *MockIterManager) Done() bool {
	ret := m.ctrl.Call(m, "Done")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Done indicates an expected call of Done
func (mr *MockIterManagerMockRecorder) Done() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockIterManager)(nil).Done))
}

// Err mocks base method
func (m *MockIterManager) Err() error {
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockIterManagerMockRecorder) Err() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockIterManager)(nil).Err))
}

// For mocks base method
func (m *MockIterManager) For(result interface{}, f func() error) error {
	ret := m.ctrl.Call(m, "For", result, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// For indicates an expected call of For
func (mr *MockIterManagerMockRecorder) For(result, f interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "For", reflect.TypeOf((*MockIterManager)(nil).For), result, f)
}

// Next mocks base method
func (m *MockIterManager) Next(result interface{}) bool {
	ret := m.ctrl.Call(m, "Next", result)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockIterManagerMockRecorder) Next(result interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockIterManager)(nil).Next), result)
}

// Timeout mocks base method
func (m *MockIterManager) Timeout() bool {
	ret := m.ctrl.Call(m, "Timeout")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Timeout indicates an expected call of Timeout
func (mr *MockIterManagerMockRecorder) Timeout() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timeout", reflect.TypeOf((*MockIterManager)(nil).Timeout))
}
