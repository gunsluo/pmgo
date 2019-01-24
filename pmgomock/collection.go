// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/luoji/gopath/src/github.com/gunsluo/pmgo/collection.go

// Package pmgomock is a generated GoMock package.
package pmgomock

import (
	mgo "github.com/globalsign/mgo"
	gomock "github.com/golang/mock/gomock"
	pmgo "github.com/gunsluo/pmgo"
	. "github.com/percona/pmgo"
	reflect "reflect"
)

// MockCollectionManager is a mock of CollectionManager interface
type MockCollectionManager struct {
	ctrl     *gomock.Controller
	recorder *MockCollectionManagerMockRecorder
}

// MockCollectionManagerMockRecorder is the mock recorder for MockCollectionManager
type MockCollectionManagerMockRecorder struct {
	mock *MockCollectionManager
}

// NewMockCollectionManager creates a new mock instance
func NewMockCollectionManager(ctrl *gomock.Controller) *MockCollectionManager {
	mock := &MockCollectionManager{ctrl: ctrl}
	mock.recorder = &MockCollectionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCollectionManager) EXPECT() *MockCollectionManagerMockRecorder {
	return m.recorder
}

// Count mocks base method
func (m *MockCollectionManager) Count() (int, error) {
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockCollectionManagerMockRecorder) Count() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockCollectionManager)(nil).Count))
}

// Create mocks base method
func (m *MockCollectionManager) Create(arg0 *mgo.CollectionInfo) error {
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockCollectionManagerMockRecorder) Create(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCollectionManager)(nil).Create), arg0)
}

// DropCollection mocks base method
func (m *MockCollectionManager) DropCollection() error {
	ret := m.ctrl.Call(m, "DropCollection")
	ret0, _ := ret[0].(error)
	return ret0
}

// DropCollection indicates an expected call of DropCollection
func (mr *MockCollectionManagerMockRecorder) DropCollection() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropCollection", reflect.TypeOf((*MockCollectionManager)(nil).DropCollection))
}

// DropIndex mocks base method
func (m *MockCollectionManager) DropIndex(key ...string) error {
	varargs := []interface{}{}
	for _, a := range key {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DropIndex", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropIndex indicates an expected call of DropIndex
func (mr *MockCollectionManagerMockRecorder) DropIndex(key ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropIndex", reflect.TypeOf((*MockCollectionManager)(nil).DropIndex), key...)
}

// DropIndexName mocks base method
func (m *MockCollectionManager) DropIndexName(name string) error {
	ret := m.ctrl.Call(m, "DropIndexName", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropIndexName indicates an expected call of DropIndexName
func (mr *MockCollectionManagerMockRecorder) DropIndexName(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropIndexName", reflect.TypeOf((*MockCollectionManager)(nil).DropIndexName), name)
}

// EnsureIndex mocks base method
func (m *MockCollectionManager) EnsureIndex(index mgo.Index) error {
	ret := m.ctrl.Call(m, "EnsureIndex", index)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureIndex indicates an expected call of EnsureIndex
func (mr *MockCollectionManagerMockRecorder) EnsureIndex(index interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureIndex", reflect.TypeOf((*MockCollectionManager)(nil).EnsureIndex), index)
}

// EnsureIndexKey mocks base method
func (m *MockCollectionManager) EnsureIndexKey(key ...string) error {
	varargs := []interface{}{}
	for _, a := range key {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnsureIndexKey", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureIndexKey indicates an expected call of EnsureIndexKey
func (mr *MockCollectionManagerMockRecorder) EnsureIndexKey(key ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureIndexKey", reflect.TypeOf((*MockCollectionManager)(nil).EnsureIndexKey), key...)
}

// Find mocks base method
func (m *MockCollectionManager) Find(arg0 interface{}) pmgo.QueryManager {
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(pmgo.QueryManager)
	return ret0
}

// Find indicates an expected call of Find
func (mr *MockCollectionManagerMockRecorder) Find(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCollectionManager)(nil).Find), arg0)
}

// FindId mocks base method
func (m *MockCollectionManager) FindId(id interface{}) pmgo.QueryManager {
	ret := m.ctrl.Call(m, "FindId", id)
	ret0, _ := ret[0].(pmgo.QueryManager)
	return ret0
}

// FindId indicates an expected call of FindId
func (mr *MockCollectionManagerMockRecorder) FindId(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindId", reflect.TypeOf((*MockCollectionManager)(nil).FindId), id)
}

// Indexes mocks base method
func (m *MockCollectionManager) Indexes() ([]mgo.Index, error) {
	ret := m.ctrl.Call(m, "Indexes")
	ret0, _ := ret[0].([]mgo.Index)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Indexes indicates an expected call of Indexes
func (mr *MockCollectionManagerMockRecorder) Indexes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Indexes", reflect.TypeOf((*MockCollectionManager)(nil).Indexes))
}

// Insert mocks base method
func (m *MockCollectionManager) Insert(docs ...interface{}) error {
	varargs := []interface{}{}
	for _, a := range docs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Insert", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockCollectionManagerMockRecorder) Insert(docs ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockCollectionManager)(nil).Insert), docs...)
}

// Pipe mocks base method
func (m *MockCollectionManager) Pipe(arg0 interface{}) pmgo.PipeManager {
	ret := m.ctrl.Call(m, "Pipe", arg0)
	ret0, _ := ret[0].(pmgo.PipeManager)
	return ret0
}

// Pipe indicates an expected call of Pipe
func (mr *MockCollectionManagerMockRecorder) Pipe(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipe", reflect.TypeOf((*MockCollectionManager)(nil).Pipe), arg0)
}

// Remove mocks base method
func (m *MockCollectionManager) Remove(selector interface{}) error {
	ret := m.ctrl.Call(m, "Remove", selector)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockCollectionManagerMockRecorder) Remove(selector interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockCollectionManager)(nil).Remove), selector)
}

// RemoveAll mocks base method
func (m *MockCollectionManager) RemoveAll(selector interface{}) (*mgo.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "RemoveAll", selector)
	ret0, _ := ret[0].(*mgo.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveAll indicates an expected call of RemoveAll
func (mr *MockCollectionManagerMockRecorder) RemoveAll(selector interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*MockCollectionManager)(nil).RemoveAll), selector)
}

// RemoveId mocks base method
func (m *MockCollectionManager) RemoveId(id interface{}) error {
	ret := m.ctrl.Call(m, "RemoveId", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveId indicates an expected call of RemoveId
func (mr *MockCollectionManagerMockRecorder) RemoveId(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveId", reflect.TypeOf((*MockCollectionManager)(nil).RemoveId), id)
}

// Update mocks base method
func (m *MockCollectionManager) Update(selector, update interface{}) error {
	ret := m.ctrl.Call(m, "Update", selector, update)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockCollectionManagerMockRecorder) Update(selector, update interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCollectionManager)(nil).Update), selector, update)
}

// UpdateAll mocks base method
func (m *MockCollectionManager) UpdateAll(selector, update interface{}) (*mgo.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "UpdateAll", selector, update)
	ret0, _ := ret[0].(*mgo.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll
func (mr *MockCollectionManagerMockRecorder) UpdateAll(selector, update interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockCollectionManager)(nil).UpdateAll), selector, update)
}

// UpdateId mocks base method
func (m *MockCollectionManager) UpdateId(id, update interface{}) error {
	ret := m.ctrl.Call(m, "UpdateId", id, update)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateId indicates an expected call of UpdateId
func (mr *MockCollectionManagerMockRecorder) UpdateId(id, update interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateId", reflect.TypeOf((*MockCollectionManager)(nil).UpdateId), id, update)
}

// Upsert mocks base method
func (m *MockCollectionManager) Upsert(selector, update interface{}) (*mgo.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "Upsert", selector, update)
	ret0, _ := ret[0].(*mgo.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert
func (mr *MockCollectionManagerMockRecorder) Upsert(selector, update interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockCollectionManager)(nil).Upsert), selector, update)
}

// UpsertId mocks base method
func (m *MockCollectionManager) UpsertId(id, update interface{}) (*mgo.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "UpsertId", id, update)
	ret0, _ := ret[0].(*mgo.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertId indicates an expected call of UpsertId
func (mr *MockCollectionManagerMockRecorder) UpsertId(id, update interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertId", reflect.TypeOf((*MockCollectionManager)(nil).UpsertId), id, update)
}
