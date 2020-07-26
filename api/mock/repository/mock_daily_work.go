// Code generated by MockGen. DO NOT EDIT.
// Source: daily_work.go

// Package repository is a generated GoMock package.
package repository

import (
	gomock "github.com/golang/mock/gomock"
	model "influ-dojo/api/domain/model"
	reflect "reflect"
)

// MockDailyWork is a mock of DailyWork interface.
type MockDailyWork struct {
	ctrl     *gomock.Controller
	recorder *MockDailyWorkMockRecorder
}

// MockDailyWorkMockRecorder is the mock recorder for MockDailyWork.
type MockDailyWorkMockRecorder struct {
	mock *MockDailyWork
}

// NewMockDailyWork creates a new mock instance.
func NewMockDailyWork(ctrl *gomock.Controller) *MockDailyWork {
	mock := &MockDailyWork{ctrl: ctrl}
	mock.recorder = &MockDailyWorkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDailyWork) EXPECT() *MockDailyWorkMockRecorder {
	return m.recorder
}

// LoadTop3 mocks base method.
func (m *MockDailyWork) LoadTop3() ([]*model.Work, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadTop3")
	ret0, _ := ret[0].([]*model.Work)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadTop3 indicates an expected call of LoadTop3.
func (mr *MockDailyWorkMockRecorder) LoadTop3() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadTop3", reflect.TypeOf((*MockDailyWork)(nil).LoadTop3))
}

// LoadByScreenName mocks base method.
func (m *MockDailyWork) LoadByScreenName(screenName string) (*model.Work, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadByScreenName", screenName)
	ret0, _ := ret[0].(*model.Work)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadByScreenName indicates an expected call of LoadByScreenName.
func (mr *MockDailyWorkMockRecorder) LoadByScreenName(screenName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadByScreenName", reflect.TypeOf((*MockDailyWork)(nil).LoadByScreenName), screenName)
}

// Save mocks base method.
func (m *MockDailyWork) Save(work *model.Work) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", work)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockDailyWorkMockRecorder) Save(work interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockDailyWork)(nil).Save), work)
}