package mock

import (
	reflect "reflect"

	model "github.com/illenko/analytics-service/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockAnalyticsRepository is a mock of AnalyticsRepository interface.
type MockAnalyticsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAnalyticsRepositoryMockRecorder
}

// MockAnalyticsRepositoryMockRecorder is the mock recorder for MockAnalyticsRepository.
type MockAnalyticsRepositoryMockRecorder struct {
	mock *MockAnalyticsRepository
}

// NewMockAnalyticsRepository creates a new mock instance.
func NewMockAnalyticsRepository(ctrl *gomock.Controller) *MockAnalyticsRepository {
	mock := &MockAnalyticsRepository{ctrl: ctrl}
	mock.recorder = &MockAnalyticsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnalyticsRepository) EXPECT() *MockAnalyticsRepositoryMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockAnalyticsRepository) Find(group string, positiveAmount bool) ([]model.AnalyticsItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", group, positiveAmount)
	ret0, _ := ret[0].([]model.AnalyticsItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockAnalyticsRepositoryMockRecorder) Find(group, positiveAmount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockAnalyticsRepository)(nil).Find), group, positiveAmount)
}

// FindByDates mocks base method.
func (m *MockAnalyticsRepository) FindByDates(positiveAmount bool, unit string) ([]model.DateAnalyticsItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByDates", positiveAmount, unit)
	ret0, _ := ret[0].([]model.DateAnalyticsItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByDates indicates an expected call of FindByDates.
func (mr *MockAnalyticsRepositoryMockRecorder) FindByDates(positiveAmount, unit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByDates", reflect.TypeOf((*MockAnalyticsRepository)(nil).FindByDates), positiveAmount, unit)
}

// FindByDatesCumulative mocks base method.
func (m *MockAnalyticsRepository) FindByDatesCumulative(positiveAmount bool, unit string) ([]model.DateAnalyticsItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByDatesCumulative", positiveAmount, unit)
	ret0, _ := ret[0].([]model.DateAnalyticsItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByDatesCumulative indicates an expected call of FindByDatesCumulative.
func (mr *MockAnalyticsRepositoryMockRecorder) FindByDatesCumulative(positiveAmount, unit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByDatesCumulative", reflect.TypeOf((*MockAnalyticsRepository)(nil).FindByDatesCumulative), positiveAmount, unit)
}
