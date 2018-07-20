// Code generated by MockGen. DO NOT EDIT.
// Source: /home/tim/git/pmgo/session.go

// Package pmgomock is a generated GoMock package.
package pmgomock

import (
	x "."
	gomock "github.com/golang/mock/gomock"
	. "github.com/percona/pmgo"
	mgo_v2 "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
	reflect "reflect"
	time "time"
)

// MockSessionManager is a mock of SessionManager interface
type MockSessionManager struct {
	ctrl     *gomock.Controller
	recorder *MockSessionManagerMockRecorder
}

// MockSessionManagerMockRecorder is the mock recorder for MockSessionManager
type MockSessionManagerMockRecorder struct {
	mock *MockSessionManager
}

// NewMockSessionManager creates a new mock instance
func NewMockSessionManager(ctrl *gomock.Controller) *MockSessionManager {
	mock := &MockSessionManager{ctrl: ctrl}
	mock.recorder = &MockSessionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessionManager) EXPECT() *MockSessionManagerMockRecorder {
	return m.recorder
}

// BuildInfo mocks base method
func (m *MockSessionManager) BuildInfo() (mgo_v2.BuildInfo, error) {
	ret := m.ctrl.Call(m, "BuildInfo")
	ret0, _ := ret[0].(mgo_v2.BuildInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildInfo indicates an expected call of BuildInfo
func (mr *MockSessionManagerMockRecorder) BuildInfo() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildInfo", reflect.TypeOf((*MockSessionManager)(nil).BuildInfo))
}

// Clone mocks base method
func (m *MockSessionManager) Clone() x.SessionManager {
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(x.SessionManager)
	return ret0
}

// Clone indicates an expected call of Clone
func (mr *MockSessionManagerMockRecorder) Clone() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockSessionManager)(nil).Clone))
}

// Close mocks base method
func (m *MockSessionManager) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockSessionManagerMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSessionManager)(nil).Close))
}

// Copy mocks base method
func (m *MockSessionManager) Copy() x.SessionManager {
	ret := m.ctrl.Call(m, "Copy")
	ret0, _ := ret[0].(x.SessionManager)
	return ret0
}

// Copy indicates an expected call of Copy
func (mr *MockSessionManagerMockRecorder) Copy() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockSessionManager)(nil).Copy))
}

// DB mocks base method
func (m *MockSessionManager) DB(name string) x.DatabaseManager {
	ret := m.ctrl.Call(m, "DB", name)
	ret0, _ := ret[0].(x.DatabaseManager)
	return ret0
}

// DB indicates an expected call of DB
func (mr *MockSessionManagerMockRecorder) DB(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockSessionManager)(nil).DB), name)
}

// DatabaseNames mocks base method
func (m *MockSessionManager) DatabaseNames() ([]string, error) {
	ret := m.ctrl.Call(m, "DatabaseNames")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatabaseNames indicates an expected call of DatabaseNames
func (mr *MockSessionManagerMockRecorder) DatabaseNames() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseNames", reflect.TypeOf((*MockSessionManager)(nil).DatabaseNames))
}

// EnsureSafe mocks base method
func (m *MockSessionManager) EnsureSafe(safe *mgo_v2.Safe) {
	m.ctrl.Call(m, "EnsureSafe", safe)
}

// EnsureSafe indicates an expected call of EnsureSafe
func (mr *MockSessionManagerMockRecorder) EnsureSafe(safe interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureSafe", reflect.TypeOf((*MockSessionManager)(nil).EnsureSafe), safe)
}

// FindRef mocks base method
func (m *MockSessionManager) FindRef(ref *mgo_v2.DBRef) x.QueryManager {
	ret := m.ctrl.Call(m, "FindRef", ref)
	ret0, _ := ret[0].(x.QueryManager)
	return ret0
}

// FindRef indicates an expected call of FindRef
func (mr *MockSessionManagerMockRecorder) FindRef(ref interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRef", reflect.TypeOf((*MockSessionManager)(nil).FindRef), ref)
}

// Fsync mocks base method
func (m *MockSessionManager) Fsync(async bool) error {
	ret := m.ctrl.Call(m, "Fsync", async)
	ret0, _ := ret[0].(error)
	return ret0
}

// Fsync indicates an expected call of Fsync
func (mr *MockSessionManagerMockRecorder) Fsync(async interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fsync", reflect.TypeOf((*MockSessionManager)(nil).Fsync), async)
}

// FsyncLock mocks base method
func (m *MockSessionManager) FsyncLock() error {
	ret := m.ctrl.Call(m, "FsyncLock")
	ret0, _ := ret[0].(error)
	return ret0
}

// FsyncLock indicates an expected call of FsyncLock
func (mr *MockSessionManagerMockRecorder) FsyncLock() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FsyncLock", reflect.TypeOf((*MockSessionManager)(nil).FsyncLock))
}

// FsyncUnlock mocks base method
func (m *MockSessionManager) FsyncUnlock() error {
	ret := m.ctrl.Call(m, "FsyncUnlock")
	ret0, _ := ret[0].(error)
	return ret0
}

// FsyncUnlock indicates an expected call of FsyncUnlock
func (mr *MockSessionManagerMockRecorder) FsyncUnlock() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FsyncUnlock", reflect.TypeOf((*MockSessionManager)(nil).FsyncUnlock))
}

// LiveServers mocks base method
func (m *MockSessionManager) LiveServers() []string {
	ret := m.ctrl.Call(m, "LiveServers")
	ret0, _ := ret[0].([]string)
	return ret0
}

// LiveServers indicates an expected call of LiveServers
func (mr *MockSessionManagerMockRecorder) LiveServers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LiveServers", reflect.TypeOf((*MockSessionManager)(nil).LiveServers))
}

// Login mocks base method
func (m *MockSessionManager) Login(cred *mgo_v2.Credential) error {
	ret := m.ctrl.Call(m, "Login", cred)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login
func (mr *MockSessionManagerMockRecorder) Login(cred interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockSessionManager)(nil).Login), cred)
}

// LogoutAll mocks base method
func (m *MockSessionManager) LogoutAll() {
	m.ctrl.Call(m, "LogoutAll")
}

// LogoutAll indicates an expected call of LogoutAll
func (mr *MockSessionManagerMockRecorder) LogoutAll() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogoutAll", reflect.TypeOf((*MockSessionManager)(nil).LogoutAll))
}

// Mode mocks base method
func (m *MockSessionManager) Mode() mgo_v2.Mode {
	ret := m.ctrl.Call(m, "Mode")
	ret0, _ := ret[0].(mgo_v2.Mode)
	return ret0
}

// Mode indicates an expected call of Mode
func (mr *MockSessionManagerMockRecorder) Mode() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mode", reflect.TypeOf((*MockSessionManager)(nil).Mode))
}

// New mocks base method
func (m *MockSessionManager) New() x.SessionManager {
	ret := m.ctrl.Call(m, "New")
	ret0, _ := ret[0].(x.SessionManager)
	return ret0
}

// New indicates an expected call of New
func (mr *MockSessionManagerMockRecorder) New() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockSessionManager)(nil).New))
}

// Ping mocks base method
func (m *MockSessionManager) Ping() error {
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
func (mr *MockSessionManagerMockRecorder) Ping() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockSessionManager)(nil).Ping))
}

// Refresh mocks base method
func (m *MockSessionManager) Refresh() {
	m.ctrl.Call(m, "Refresh")
}

// Refresh indicates an expected call of Refresh
func (mr *MockSessionManagerMockRecorder) Refresh() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockSessionManager)(nil).Refresh))
}

// RemoveUser mocks base method
func (m *MockSessionManager) RemoveUser(user string) error {
	ret := m.ctrl.Call(m, "RemoveUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUser indicates an expected call of RemoveUser
func (mr *MockSessionManagerMockRecorder) RemoveUser(user interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockSessionManager)(nil).RemoveUser), user)
}

// ResetIndexCache mocks base method
func (m *MockSessionManager) ResetIndexCache() {
	m.ctrl.Call(m, "ResetIndexCache")
}

// ResetIndexCache indicates an expected call of ResetIndexCache
func (mr *MockSessionManagerMockRecorder) ResetIndexCache() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetIndexCache", reflect.TypeOf((*MockSessionManager)(nil).ResetIndexCache))
}

// Run mocks base method
func (m *MockSessionManager) Run(cmd, result interface{}) error {
	ret := m.ctrl.Call(m, "Run", cmd, result)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockSessionManagerMockRecorder) Run(cmd, result interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockSessionManager)(nil).Run), cmd, result)
}

// Safe mocks base method
func (m *MockSessionManager) Safe() *mgo_v2.Safe {
	ret := m.ctrl.Call(m, "Safe")
	ret0, _ := ret[0].(*mgo_v2.Safe)
	return ret0
}

// Safe indicates an expected call of Safe
func (mr *MockSessionManagerMockRecorder) Safe() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Safe", reflect.TypeOf((*MockSessionManager)(nil).Safe))
}

// SelectServers mocks base method
func (m *MockSessionManager) SelectServers(tags ...bson.D) {
	varargs := []interface{}{}
	for _, a := range tags {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "SelectServers", varargs...)
}

// SelectServers indicates an expected call of SelectServers
func (mr *MockSessionManagerMockRecorder) SelectServers(tags ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectServers", reflect.TypeOf((*MockSessionManager)(nil).SelectServers), tags...)
}

// SetBatch mocks base method
func (m *MockSessionManager) SetBatch(n int) {
	m.ctrl.Call(m, "SetBatch", n)
}

// SetBatch indicates an expected call of SetBatch
func (mr *MockSessionManagerMockRecorder) SetBatch(n interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBatch", reflect.TypeOf((*MockSessionManager)(nil).SetBatch), n)
}

// SetBypassValidation mocks base method
func (m *MockSessionManager) SetBypassValidation(bypass bool) {
	m.ctrl.Call(m, "SetBypassValidation", bypass)
}

// SetBypassValidation indicates an expected call of SetBypassValidation
func (mr *MockSessionManagerMockRecorder) SetBypassValidation(bypass interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBypassValidation", reflect.TypeOf((*MockSessionManager)(nil).SetBypassValidation), bypass)
}

// SetCursorTimeout mocks base method
func (m *MockSessionManager) SetCursorTimeout(d time.Duration) {
	m.ctrl.Call(m, "SetCursorTimeout", d)
}

// SetCursorTimeout indicates an expected call of SetCursorTimeout
func (mr *MockSessionManagerMockRecorder) SetCursorTimeout(d interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCursorTimeout", reflect.TypeOf((*MockSessionManager)(nil).SetCursorTimeout), d)
}

// SetMode mocks base method
func (m *MockSessionManager) SetMode(consistency mgo_v2.Mode, refresh bool) {
	m.ctrl.Call(m, "SetMode", consistency, refresh)
}

// SetMode indicates an expected call of SetMode
func (mr *MockSessionManagerMockRecorder) SetMode(consistency, refresh interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMode", reflect.TypeOf((*MockSessionManager)(nil).SetMode), consistency, refresh)
}

// SetPoolLimit mocks base method
func (m *MockSessionManager) SetPoolLimit(limit int) {
	m.ctrl.Call(m, "SetPoolLimit", limit)
}

// SetPoolLimit indicates an expected call of SetPoolLimit
func (mr *MockSessionManagerMockRecorder) SetPoolLimit(limit interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPoolLimit", reflect.TypeOf((*MockSessionManager)(nil).SetPoolLimit), limit)
}

// SetPrefetch mocks base method
func (m *MockSessionManager) SetPrefetch(p float64) {
	m.ctrl.Call(m, "SetPrefetch", p)
}

// SetPrefetch indicates an expected call of SetPrefetch
func (mr *MockSessionManagerMockRecorder) SetPrefetch(p interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPrefetch", reflect.TypeOf((*MockSessionManager)(nil).SetPrefetch), p)
}

// SetSafe mocks base method
func (m *MockSessionManager) SetSafe(safe *mgo_v2.Safe) {
	m.ctrl.Call(m, "SetSafe", safe)
}

// SetSafe indicates an expected call of SetSafe
func (mr *MockSessionManagerMockRecorder) SetSafe(safe interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSafe", reflect.TypeOf((*MockSessionManager)(nil).SetSafe), safe)
}

// SetSocketTimeout mocks base method
func (m *MockSessionManager) SetSocketTimeout(d time.Duration) {
	m.ctrl.Call(m, "SetSocketTimeout", d)
}

// SetSocketTimeout indicates an expected call of SetSocketTimeout
func (mr *MockSessionManagerMockRecorder) SetSocketTimeout(d interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSocketTimeout", reflect.TypeOf((*MockSessionManager)(nil).SetSocketTimeout), d)
}

// SetSyncTimeout mocks base method
func (m *MockSessionManager) SetSyncTimeout(d time.Duration) {
	m.ctrl.Call(m, "SetSyncTimeout", d)
}

// SetSyncTimeout indicates an expected call of SetSyncTimeout
func (mr *MockSessionManagerMockRecorder) SetSyncTimeout(d interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSyncTimeout", reflect.TypeOf((*MockSessionManager)(nil).SetSyncTimeout), d)
}

// UpsertUser mocks base method
func (m *MockSessionManager) UpsertUser(user *mgo_v2.User) error {
	ret := m.ctrl.Call(m, "UpsertUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertUser indicates an expected call of UpsertUser
func (mr *MockSessionManagerMockRecorder) UpsertUser(user interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertUser", reflect.TypeOf((*MockSessionManager)(nil).UpsertUser), user)
}
