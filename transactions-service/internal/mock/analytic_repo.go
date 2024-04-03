package mock

import (
	reflect "reflect"

	model "github.com/illenko/transactions-service/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockAnalyticRepository is a mock of AnalyticRepository interface.
type MockAnalyticRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAnalyticRepositoryMockRecorder
}

// MockAnalyticRepositoryMockRecorder is the mock recorder for MockAnalyticRepository.
type MockAnalyticRepositoryMockRecorder struct {
	mock *MockAnalyticRepository
}

// NewMockAnalyticRepository creates a new mock instance.
func NewMockAnalyticRepository(ctrl *gomock.Controller) *MockAnalyticRepository {
	mock := &MockAnalyticRepository{ctrl: ctrl}
	mock.recorder = &MockAnalyticRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnalyticRepository) EXPECT() *MockAnalyticRepositoryMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockAnalyticRepository) Find(group string, positiveAmount bool) ([]model.AnalyticItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", group, positiveAmount)
	ret0, _ := ret[0].([]model.AnalyticItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockAnalyticRepositoryMockRecorder) Find(group, positiveAmount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockAnalyticRepository)(nil).Find), group, positiveAmount)
}

// FindByDates mocks base method.
func (m *MockAnalyticRepository) FindByDates(positiveAmount bool, unit string) ([]model.DateAnalyticItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByDates", positiveAmount, unit)
	ret0, _ := ret[0].([]model.DateAnalyticItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByDates indicates an expected call of FindByDates.
func (mr *MockAnalyticRepositoryMockRecorder) FindByDates(positiveAmount, unit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByDates", reflect.TypeOf((*MockAnalyticRepository)(nil).FindByDates), positiveAmount, unit)
}

// FindByDatesCumulative mocks base method.
func (m *MockAnalyticRepository) FindByDatesCumulative(positiveAmount bool, unit string) ([]model.DateAnalyticItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByDatesCumulative", positiveAmount, unit)
	ret0, _ := ret[0].([]model.DateAnalyticItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByDatesCumulative indicates an expected call of FindByDatesCumulative.
func (mr *MockAnalyticRepositoryMockRecorder) FindByDatesCumulative(positiveAmount, unit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByDatesCumulative", reflect.TypeOf((*MockAnalyticRepository)(nil).FindByDatesCumulative), positiveAmount, unit)
}
