package transaction

import (
	"context"
	"github.com/google/uuid"
	"github.com/illenko/transactions-service/internal/transaction/mock"
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
	repository := mock.NewMockRepository(ctrl)
	mapper := mock.NewMockMapper(ctrl)

	var transactions []Entity
	gofakeit.Slice(&transactions)

	var expected []model.TransactionResponse
	gofakeit.Slice(&expected)

	repository.EXPECT().FindAll().Return(transactions, nil)
	mapper.EXPECT().ToResponses(transactions).Return(expected)

	sut := NewService(logger, repository, mapper)

	actual, err := sut.FindAll(context.Background())

	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestFindById(t *testing.T) {
	ctrl := gomock.NewController(t)

	logger := slog.Default()
	repository := mock.NewMockRepository(ctrl)
	mapper := mock.NewMockMapper(ctrl)
	id := uuid.New()

	var transaction Entity
	gofakeit.Struct(&transaction)

	var expected model.TransactionResponse
	gofakeit.Struct(&expected)

	repository.EXPECT().FindById(id).Return(transaction, nil)
	mapper.EXPECT().ToResponse(transaction).Return(expected)

	sut := NewService(logger, repository, mapper)

	actual, err := sut.FindById(context.Background(), id)

	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}
