package mock

import (
	reflect "reflect"

	model "github.com/illenko/transactions-service/internal/model"
	model0 "github.com/illenko/transactions-service/pkg/model"
	gomock "go.uber.org/mock/gomock"
)

// MockTransactionMapper is a mock of TransactionMapper interface.
type MockTransactionMapper struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionMapperMockRecorder
}

// MockTransactionMapperMockRecorder is the mock recorder for MockTransactionMapper.
type MockTransactionMapperMockRecorder struct {
	mock *MockTransactionMapper
}

// NewMockTransactionMapper creates a new mock instance.
func NewMockTransactionMapper(ctrl *gomock.Controller) *MockTransactionMapper {
	mock := &MockTransactionMapper{ctrl: ctrl}
	mock.recorder = &MockTransactionMapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionMapper) EXPECT() *MockTransactionMapperMockRecorder {
	return m.recorder
}

// ToResponse mocks base method.
func (m *MockTransactionMapper) ToResponse(entity model.Transaction) model0.TransactionResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToResponse", entity)
	ret0, _ := ret[0].(model0.TransactionResponse)
	return ret0
}

// ToResponse indicates an expected call of ToResponse.
func (mr *MockTransactionMapperMockRecorder) ToResponse(entity any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToResponse", reflect.TypeOf((*MockTransactionMapper)(nil).ToResponse), entity)
}

// ToResponses mocks base method.
func (m *MockTransactionMapper) ToResponses(entities []model.Transaction) []model0.TransactionResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToResponses", entities)
	ret0, _ := ret[0].([]model0.TransactionResponse)
	return ret0
}

// ToResponses indicates an expected call of ToResponses.
func (mr *MockTransactionMapperMockRecorder) ToResponses(entities any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToResponses", reflect.TypeOf((*MockTransactionMapper)(nil).ToResponses), entities)
}
