package mock

import (
	reflect "reflect"

	uuid "github.com/google/uuid"
	model "github.com/illenko/transactions-service/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockTransactionRepository is a mock of TransactionRepository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockTransactionRepository) FindAll() ([]model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockTransactionRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockTransactionRepository)(nil).FindAll))
}

// FindById mocks base method.
func (m *MockTransactionRepository) FindById(id uuid.UUID) (model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockTransactionRepositoryMockRecorder) FindById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockTransactionRepository)(nil).FindById), id)
}
