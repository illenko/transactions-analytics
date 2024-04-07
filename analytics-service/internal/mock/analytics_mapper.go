package mock

import (
	reflect "reflect"

	model "github.com/illenko/analytics-service/internal/model"
	model0 "github.com/illenko/analytics-service/pkg/model"
	gomock "go.uber.org/mock/gomock"
)

// MockAnalyticsMapper is a mock of AnalyticsMapper interface.
type MockAnalyticsMapper struct {
	ctrl     *gomock.Controller
	recorder *MockAnalyticsMapperMockRecorder
}

// MockAnalyticsMapperMockRecorder is the mock recorder for MockAnalyticsMapper.
type MockAnalyticsMapperMockRecorder struct {
	mock *MockAnalyticsMapper
}

// NewMockAnalyticsMapper creates a new mock instance.
func NewMockAnalyticsMapper(ctrl *gomock.Controller) *MockAnalyticsMapper {
	mock := &MockAnalyticsMapper{ctrl: ctrl}
	mock.recorder = &MockAnalyticsMapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnalyticsMapper) EXPECT() *MockAnalyticsMapperMockRecorder {
	return m.recorder
}

// ToDayResponse mocks base method.
func (m *MockAnalyticsMapper) ToDayResponse(items []model.DateAnalyticsItem) model0.AnalyticsResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToDayResponse", items)
	ret0, _ := ret[0].(model0.AnalyticsResponse)
	return ret0
}

// ToDayResponse indicates an expected call of ToDayResponse.
func (mr *MockAnalyticsMapperMockRecorder) ToDayResponse(items any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToDayResponse", reflect.TypeOf((*MockAnalyticsMapper)(nil).ToDayResponse), items)
}

// ToMonthResponse mocks base method.
func (m *MockAnalyticsMapper) ToMonthResponse(items []model.DateAnalyticsItem) model0.AnalyticsResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToMonthResponse", items)
	ret0, _ := ret[0].(model0.AnalyticsResponse)
	return ret0
}

// ToMonthResponse indicates an expected call of ToMonthResponse.
func (mr *MockAnalyticsMapperMockRecorder) ToMonthResponse(items any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToMonthResponse", reflect.TypeOf((*MockAnalyticsMapper)(nil).ToMonthResponse), items)
}

// ToResponse mocks base method.
func (m *MockAnalyticsMapper) ToResponse(items []model.AnalyticsItem) model0.AnalyticsResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToResponse", items)
	ret0, _ := ret[0].(model0.AnalyticsResponse)
	return ret0
}

// ToResponse indicates an expected call of ToResponse.
func (mr *MockAnalyticsMapperMockRecorder) ToResponse(items any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToResponse", reflect.TypeOf((*MockAnalyticsMapper)(nil).ToResponse), items)
}
