package database

import (
	"github.com/google/uuid"
	"github.com/illenko/analytics-service/internal/database"
	"github.com/illenko/analytics-service/internal/model"
	"github.com/illenko/analytics-service/tests/integration"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type TransactionRepositoryTestSuite struct {
	integration.BaseIntegrationTestSuite
}

func (suite *TransactionRepositoryTestSuite) TestFindAll() {
	suite.Run("Find transaction by id", func() {
		repo := database.NewTransactionRepository(suite.Log, suite.DB)
		transactions, err := repo.FindAll()
		suite.NoError(err)
		suite.NotNil(transactions)
		suite.Equal(len(transactions), 16)
	})

}

func (suite *TransactionRepositoryTestSuite) TestFindById() {
	suite.Run("Find transaction by id", func() {
		repo := database.NewTransactionRepository(suite.Log, suite.DB)
		transaction, err := repo.FindById(uuid.MustParse("44bdcdbc-4eae-443d-9bbd-4d1c1b7e628a"))
		suite.NoError(err)
		suite.Equal(transaction, model.Transaction{
			ID:       uuid.MustParse("44bdcdbc-4eae-443d-9bbd-4d1c1b7e628a"),
			Datetime: time.Date(2024, time.March, 28, 16, 40, 31, 0, time.UTC),
			Amount:   -13,
			Category: "taxi",
			Merchant: "uklon",
		})
	})

	suite.Run("Find not existing transaction by id", func() {
		repo := database.NewTransactionRepository(suite.Log, suite.DB)
		_, err := repo.FindById(uuid.New())
		suite.Error(err)
	})
}

func TestTransactionRepository(t *testing.T) {
	suite.Run(t, new(TransactionRepositoryTestSuite))
}
