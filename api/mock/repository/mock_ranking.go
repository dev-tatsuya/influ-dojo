// Code generated by MockGen. DO NOT EDIT.
// Source: ranking.go

// Package repository is a generated GoMock package.
package repository

import (
	gomock "github.com/golang/mock/gomock"
	dto "influ-dojo/api/usecase/dto"
	reflect "reflect"
)

// MockRanking is a mock of Ranking interface.
type MockRanking struct {
	ctrl     *gomock.Controller
	recorder *MockRankingMockRecorder
}

// MockRankingMockRecorder is the mock recorder for MockRanking.
type MockRankingMockRecorder struct {
	mock *MockRanking
}

// NewMockRanking creates a new mock instance.
func NewMockRanking(ctrl *gomock.Controller) *MockRanking {
	mock := &MockRanking{ctrl: ctrl}
	mock.recorder = &MockRankingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRanking) EXPECT() *MockRankingMockRecorder {
	return m.recorder
}

// LoadAll mocks base method.
func (m *MockRanking) LoadAll() (*dto.RankingAll, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadAll")
	ret0, _ := ret[0].(*dto.RankingAll)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadAll indicates an expected call of LoadAll.
func (mr *MockRankingMockRecorder) LoadAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadAll", reflect.TypeOf((*MockRanking)(nil).LoadAll))
}

// Store mocks base method.
func (m *MockRanking) Store(all *dto.RankingAll) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", all)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockRankingMockRecorder) Store(all interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockRanking)(nil).Store), all)
}
