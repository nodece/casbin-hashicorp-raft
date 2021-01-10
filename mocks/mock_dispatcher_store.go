// Code generated by MockGen. DO NOT EDIT.
// Source: dispatcher_store.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDispatcherStore is a mock of DispatcherStore interface
type MockDispatcherStore struct {
	ctrl     *gomock.Controller
	recorder *MockDispatcherStoreMockRecorder
}

// MockDispatcherStoreMockRecorder is the mock recorder for MockDispatcherStore
type MockDispatcherStoreMockRecorder struct {
	mock *MockDispatcherStore
}

// NewMockDispatcherStore creates a new mock instance
func NewMockDispatcherStore(ctrl *gomock.Controller) *MockDispatcherStore {
	mock := &MockDispatcherStore{ctrl: ctrl}
	mock.recorder = &MockDispatcherStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDispatcherStore) EXPECT() *MockDispatcherStoreMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockDispatcherStore) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockDispatcherStoreMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockDispatcherStore)(nil).Start))
}

// Stop mocks base method
func (m *MockDispatcherStore) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockDispatcherStoreMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockDispatcherStore)(nil).Stop))
}

// AddNode mocks base method
func (m *MockDispatcherStore) AddNode(serverID, address string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNode", serverID, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNode indicates an expected call of AddNode
func (mr *MockDispatcherStoreMockRecorder) AddNode(serverID, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNode", reflect.TypeOf((*MockDispatcherStore)(nil).AddNode), serverID, address)
}

// RemoveNode mocks base method
func (m *MockDispatcherStore) RemoveNode(serverID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveNode", serverID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveNode indicates an expected call of RemoveNode
func (mr *MockDispatcherStoreMockRecorder) RemoveNode(serverID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNode", reflect.TypeOf((*MockDispatcherStore)(nil).RemoveNode), serverID)
}

// Apply mocks base method
func (m *MockDispatcherStore) Apply(buf []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", buf)
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply
func (mr *MockDispatcherStoreMockRecorder) Apply(buf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockDispatcherStore)(nil).Apply), buf)
}

// Leader mocks base method
func (m *MockDispatcherStore) Leader() (bool, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Leader")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	return ret0, ret1
}

// Leader indicates an expected call of Leader
func (mr *MockDispatcherStoreMockRecorder) Leader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leader", reflect.TypeOf((*MockDispatcherStore)(nil).Leader))
}