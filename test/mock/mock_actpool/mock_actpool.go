// Code generated by MockGen. DO NOT EDIT.
// Source: ./actpool/actpool.go

// Package mock_actpool is a generated GoMock package.
package mock_actpool

import (
	gomock "github.com/golang/mock/gomock"
	action "github.com/iotexproject/iotex-core/blockchain/action"
	proto "github.com/iotexproject/iotex-core/proto"
	reflect "reflect"
)

// MockActPool is a mock of ActPool interface
type MockActPool struct {
	ctrl     *gomock.Controller
	recorder *MockActPoolMockRecorder
}

// MockActPoolMockRecorder is the mock recorder for MockActPool
type MockActPoolMockRecorder struct {
	mock *MockActPool
}

// NewMockActPool creates a new mock instance
func NewMockActPool(ctrl *gomock.Controller) *MockActPool {
	mock := &MockActPool{ctrl: ctrl}
	mock.recorder = &MockActPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockActPool) EXPECT() *MockActPoolMockRecorder {
	return m.recorder
}

// Reset mocks base method
func (m *MockActPool) Reset() {
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset
func (mr *MockActPoolMockRecorder) Reset() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockActPool)(nil).Reset))
}

// PickActs mocks base method
func (m *MockActPool) PickActs() ([]*action.Transfer, []*action.Vote, []*action.Execution) {
	ret := m.ctrl.Call(m, "PickActs")
	ret0, _ := ret[0].([]*action.Transfer)
	ret1, _ := ret[1].([]*action.Vote)
	ret2, _ := ret[2].([]*action.Execution)
	return ret0, ret1, ret2
}

// PickActs indicates an expected call of PickActs
func (mr *MockActPoolMockRecorder) PickActs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PickActs", reflect.TypeOf((*MockActPool)(nil).PickActs))
}

// AddTsf mocks base method
func (m *MockActPool) AddTsf(tsf *action.Transfer) error {
	ret := m.ctrl.Call(m, "AddTsf", tsf)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTsf indicates an expected call of AddTsf
func (mr *MockActPoolMockRecorder) AddTsf(tsf interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTsf", reflect.TypeOf((*MockActPool)(nil).AddTsf), tsf)
}

// AddVote mocks base method
func (m *MockActPool) AddVote(vote *action.Vote) error {
	ret := m.ctrl.Call(m, "AddVote", vote)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddVote indicates an expected call of AddVote
func (mr *MockActPoolMockRecorder) AddVote(vote interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVote", reflect.TypeOf((*MockActPool)(nil).AddVote), vote)
}

// AddExecution mocks base method
func (m *MockActPool) AddExecution(execution *action.Execution) error {
	ret := m.ctrl.Call(m, "AddExecution", execution)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddExecution indicates an expected call of AddExecution
func (mr *MockActPoolMockRecorder) AddExecution(execution interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddExecution", reflect.TypeOf((*MockActPool)(nil).AddExecution), execution)
}

// GetPendingNonce mocks base method
func (m *MockActPool) GetPendingNonce(addr string) (uint64, error) {
	ret := m.ctrl.Call(m, "GetPendingNonce", addr)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPendingNonce indicates an expected call of GetPendingNonce
func (mr *MockActPoolMockRecorder) GetPendingNonce(addr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingNonce", reflect.TypeOf((*MockActPool)(nil).GetPendingNonce), addr)
}

// GetUnconfirmedActs mocks base method
func (m *MockActPool) GetUnconfirmedActs(addr string) []*proto.ActionPb {
	ret := m.ctrl.Call(m, "GetUnconfirmedActs", addr)
	ret0, _ := ret[0].([]*proto.ActionPb)
	return ret0
}

// GetUnconfirmedActs indicates an expected call of GetUnconfirmedActs
func (mr *MockActPoolMockRecorder) GetUnconfirmedActs(addr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnconfirmedActs", reflect.TypeOf((*MockActPool)(nil).GetUnconfirmedActs), addr)
}
