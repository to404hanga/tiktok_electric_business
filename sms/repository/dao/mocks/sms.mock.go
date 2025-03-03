// Code generated by MockGen. DO NOT EDIT.
// Source: ./types.go

// Package daomocks is a generated GoMock package.
package daomocks

import (
	context "context"
	reflect "reflect"
	dao "tiktok_electric_business/sms/repository/dao"

	gomock "github.com/golang/mock/gomock"
)

// MockAsyncSmsDAO is a mock of AsyncSmsDAO interface.
type MockAsyncSmsDAO struct {
	ctrl     *gomock.Controller
	recorder *MockAsyncSmsDAOMockRecorder
}

// MockAsyncSmsDAOMockRecorder is the mock recorder for MockAsyncSmsDAO.
type MockAsyncSmsDAOMockRecorder struct {
	mock *MockAsyncSmsDAO
}

// NewMockAsyncSmsDAO creates a new mock instance.
func NewMockAsyncSmsDAO(ctrl *gomock.Controller) *MockAsyncSmsDAO {
	mock := &MockAsyncSmsDAO{ctrl: ctrl}
	mock.recorder = &MockAsyncSmsDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAsyncSmsDAO) EXPECT() *MockAsyncSmsDAOMockRecorder {
	return m.recorder
}

// GetWaitingSMS mocks base method.
func (m *MockAsyncSmsDAO) GetWaitingSMS(ctx context.Context) (dao.AsyncSms, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWaitingSMS", ctx)
	ret0, _ := ret[0].(dao.AsyncSms)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWaitingSMS indicates an expected call of GetWaitingSMS.
func (mr *MockAsyncSmsDAOMockRecorder) GetWaitingSMS(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWaitingSMS", reflect.TypeOf((*MockAsyncSmsDAO)(nil).GetWaitingSMS), ctx)
}

// Insert mocks base method.
func (m *MockAsyncSmsDAO) Insert(ctx context.Context, sms dao.AsyncSms) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, sms)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockAsyncSmsDAOMockRecorder) Insert(ctx, sms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockAsyncSmsDAO)(nil).Insert), ctx, sms)
}

// MarkFailed mocks base method.
func (m *MockAsyncSmsDAO) MarkFailed(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkFailed", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkFailed indicates an expected call of MarkFailed.
func (mr *MockAsyncSmsDAOMockRecorder) MarkFailed(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkFailed", reflect.TypeOf((*MockAsyncSmsDAO)(nil).MarkFailed), ctx, id)
}

// MarkSuccess mocks base method.
func (m *MockAsyncSmsDAO) MarkSuccess(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkSuccess", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkSuccess indicates an expected call of MarkSuccess.
func (mr *MockAsyncSmsDAOMockRecorder) MarkSuccess(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkSuccess", reflect.TypeOf((*MockAsyncSmsDAO)(nil).MarkSuccess), ctx, id)
}
