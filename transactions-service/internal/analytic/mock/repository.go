package mock

import (
	reflect "reflect"

	analytic "github.com/illenko/transactions-service/internal/analytic"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockRepository) Find(group string, positiveAmount bool) ([]analytic.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", group, positiveAmount)
	ret0, _ := ret[0].([]analytic.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockRepositoryMockRecorder) Find(group, positiveAmount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRepository)(nil).Find), group, positiveAmount)
}

// FindByDates mocks base method.
func (m *MockRepository) FindByDates(positiveAmount bool, unit string) ([]analytic.DateItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByDates", positiveAmount, unit)
	ret0, _ := ret[0].([]analytic.DateItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByDates indicates an expected call of FindByDates.
func (mr *MockRepositoryMockRecorder) FindByDates(positiveAmount, unit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByDates", reflect.TypeOf((*MockRepository)(nil).FindByDates), positiveAmount, unit)
}

// FindByDatesCumulative mocks base method.
func (m *MockRepository) FindByDatesCumulative(positiveAmount bool, unit string) ([]analytic.DateItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByDatesCumulative", positiveAmount, unit)
	ret0, _ := ret[0].([]analytic.DateItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByDatesCumulative indicates an expected call of FindByDatesCumulative.
func (mr *MockRepositoryMockRecorder) FindByDatesCumulative(positiveAmount, unit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByDatesCumulative", reflect.TypeOf((*MockRepository)(nil).FindByDatesCumulative), positiveAmount, unit)
}
