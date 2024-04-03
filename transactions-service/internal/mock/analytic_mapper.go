package mock

import (
	reflect "reflect"

	model "github.com/illenko/transactions-service/internal/model"
	model0 "github.com/illenko/transactions-service/pkg/model"
	gomock "go.uber.org/mock/gomock"
)

// MockAnalyticMapper is a mock of AnalyticMapper interface.
type MockAnalyticMapper struct {
	ctrl     *gomock.Controller
	recorder *MockAnalyticMapperMockRecorder
}

// MockAnalyticMapperMockRecorder is the mock recorder for MockAnalyticMapper.
type MockAnalyticMapperMockRecorder struct {
	mock *MockAnalyticMapper
}

// NewMockAnalyticMapper creates a new mock instance.
func NewMockAnalyticMapper(ctrl *gomock.Controller) *MockAnalyticMapper {
	mock := &MockAnalyticMapper{ctrl: ctrl}
	mock.recorder = &MockAnalyticMapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnalyticMapper) EXPECT() *MockAnalyticMapperMockRecorder {
	return m.recorder
}

// ToDayResponse mocks base method.
func (m *MockAnalyticMapper) ToDayResponse(items []model.DateAnalyticItem) model0.AnalyticResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToDayResponse", items)
	ret0, _ := ret[0].(model0.AnalyticResponse)
	return ret0
}

// ToDayResponse indicates an expected call of ToDayResponse.
func (mr *MockAnalyticMapperMockRecorder) ToDayResponse(items any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToDayResponse", reflect.TypeOf((*MockAnalyticMapper)(nil).ToDayResponse), items)
}

// ToMonthResponse mocks base method.
func (m *MockAnalyticMapper) ToMonthResponse(items []model.DateAnalyticItem) model0.AnalyticResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToMonthResponse", items)
	ret0, _ := ret[0].(model0.AnalyticResponse)
	return ret0
}

// ToMonthResponse indicates an expected call of ToMonthResponse.
func (mr *MockAnalyticMapperMockRecorder) ToMonthResponse(items any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToMonthResponse", reflect.TypeOf((*MockAnalyticMapper)(nil).ToMonthResponse), items)
}

// ToResponse mocks base method.
func (m *MockAnalyticMapper) ToResponse(items []model.AnalyticItem) model0.AnalyticResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToResponse", items)
	ret0, _ := ret[0].(model0.AnalyticResponse)
	return ret0
}

// ToResponse indicates an expected call of ToResponse.
func (mr *MockAnalyticMapperMockRecorder) ToResponse(items any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToResponse", reflect.TypeOf((*MockAnalyticMapper)(nil).ToResponse), items)
}
