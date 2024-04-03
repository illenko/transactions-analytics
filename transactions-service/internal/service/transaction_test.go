package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/illenko/transactions-service/internal/mock"
	dbmodel "github.com/illenko/transactions-service/internal/model"
	"github.com/illenko/transactions-service/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/brianvoe/gofakeit/v6"
	"log/slog"
	"testing"
)

func TestFindAll(t *testing.T) {
	ctrl := gomock.NewController(t)

	logger := slog.Default()
	repository := mock.NewMockTransactionRepository(ctrl)
	mapper := mock.NewMockTransactionMapper(ctrl)

	var transactions []dbmodel.Transaction
	gofakeit.Slice(&transactions)

	var expected []model.TransactionResponse
	gofakeit.Slice(&expected)

	repository.EXPECT().FindAll().Return(transactions, nil)
	mapper.EXPECT().ToResponses(transactions).Return(expected)

	sut := NewTransactionService(logger, repository, mapper)

	actual, err := sut.FindAll(context.Background())

	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestFindById(t *testing.T) {
	ctrl := gomock.NewController(t)

	logger := slog.Default()
	repository := mock.NewMockTransactionRepository(ctrl)
	mapper := mock.NewMockTransactionMapper(ctrl)
	id := uuid.New()

	var transaction dbmodel.Transaction
	gofakeit.Struct(&transaction)

	var expected model.TransactionResponse
	gofakeit.Struct(&expected)

	repository.EXPECT().FindById(id).Return(transaction, nil)
	mapper.EXPECT().ToResponse(transaction).Return(expected)

	sut := NewTransactionService(logger, repository, mapper)

	actual, err := sut.FindById(context.Background(), id)

	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}
