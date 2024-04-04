package mock

import (
	reflect "reflect"

	transaction "github.com/illenko/transactions-service/internal/transaction"
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

// ToResponse mocks base method.
func (m *MockMapper) ToResponse(entity transaction.Entity) model.TransactionResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToResponse", entity)
	ret0, _ := ret[0].(model.TransactionResponse)
	return ret0
}

// ToResponse indicates an expected call of ToResponse.
func (mr *MockMapperMockRecorder) ToResponse(entity any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToResponse", reflect.TypeOf((*MockMapper)(nil).ToResponse), entity)
}

// ToResponses mocks base method.
func (m *MockMapper) ToResponses(entities []transaction.Entity) []model.TransactionResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToResponses", entities)
	ret0, _ := ret[0].([]model.TransactionResponse)
	return ret0
}

// ToResponses indicates an expected call of ToResponses.
func (mr *MockMapperMockRecorder) ToResponses(entities any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToResponses", reflect.TypeOf((*MockMapper)(nil).ToResponses), entities)
}
