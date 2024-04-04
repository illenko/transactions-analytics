package mock

import (
	reflect "reflect"

	analytic "github.com/illenko/transactions-service/internal/analytic"
	model "github.com/illenko/transactions-service/pkg/model"
	gomock "go.uber.org/mock/gomock"
)

// MockMapper is a mock of Mapper interface.
type MockMapper struct {
	ctrl     *gomock.Controller
	recorder *MockMapperMockRecorder
}

// MockMapperMockRecorder is the mock recorder for MockMapper.
type MockMapperMockRecorder struct {
	mock *MockMapper
}

// NewMockMapper creates a new mock instance.
func NewMockMapper(ctrl *gomock.Controller) *MockMapper {
	mock := &MockMapper{ctrl: ctrl}
	mock.recorder = &MockMapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMapper) EXPECT() *MockMapperMockRecorder {
	return m.recorder
}

// ToDayResponse mocks base method.
func (m *MockMapper) ToDayResponse(items []analytic.DateItem) model.AnalyticResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToDayResponse", items)
	ret0, _ := ret[0].(model.AnalyticResponse)
	return ret0
}

// ToDayResponse indicates an expected call of ToDayResponse.
func (mr *MockMapperMockRecorder) ToDayResponse(items any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToDayResponse", reflect.TypeOf((*MockMapper)(nil).ToDayResponse), items)
}

// ToMonthResponse mocks base method.
func (m *MockMapper) ToMonthResponse(items []analytic.DateItem) model.AnalyticResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToMonthResponse", items)
	ret0, _ := ret[0].(model.AnalyticResponse)
	return ret0
}

// ToMonthResponse indicates an expected call of ToMonthResponse.
func (mr *MockMapperMockRecorder) ToMonthResponse(items any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToMonthResponse", reflect.TypeOf((*MockMapper)(nil).ToMonthResponse), items)
}

// ToResponse mocks base method.
func (m *MockMapper) ToResponse(items []analytic.Item) model.AnalyticResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToResponse", items)
	ret0, _ := ret[0].(model.AnalyticResponse)
	return ret0
}

// ToResponse indicates an expected call of ToResponse.
func (mr *MockMapperMockRecorder) ToResponse(items any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToResponse", reflect.TypeOf((*MockMapper)(nil).ToResponse), items)
}
