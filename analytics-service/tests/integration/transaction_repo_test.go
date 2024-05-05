package integration

import (
	"github.com/google/uuid"
	"github.com/illenko/analytics-service/internal/database"
	"github.com/illenko/analytics-service/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"log/slog"
	"testing"
	"time"
)

type TransactionRepositoryTestSuite struct {
	BaseIntegrationTestSuite
	repo database.TransactionRepository
}

func (suite *TransactionRepositoryTestSuite) SetupTest() {
	suite.repo = database.NewTransactionRepository(slog.Default(), suite.DB)
}

func (suite *TransactionRepositoryTestSuite) TestFindById() {
	testCases := []struct {
		name           string
		id             uuid.UUID
		expectedResult model.Transaction
		expectedError  error
	}{
		{
			name: "Returns a row",
			id:   uuid.MustParse("44bdcdbc-4eae-443d-9bbd-4d1c1b7e628a"),
			expectedResult: model.Transaction{
				ID:       uuid.MustParse("44bdcdbc-4eae-443d-9bbd-4d1c1b7e628a"),
				Datetime: time.Date(2024, 3, 28, 16, 40, 31, 0, time.UTC),
				Amount:   -13.00,
				Category: "taxi",
				Merchant: "uklon",
			},
			expectedError: nil,
		},
		{
			name:           "Returns no rows",
			id:             uuid.New(),
			expectedResult: model.Transaction{},
			expectedError:  gorm.ErrRecordNotFound,
		},
	}

	for _, tc := range testCases {
		suite.T().Run(tc.name, func(t *testing.T) {
			result, err := suite.repo.FindById(tc.id)
			assert.Equal(t, tc.expectedError, err)
			if err == nil {
				assert.Equal(t, tc.expectedResult, *result)
			}
		})
	}
}

func TestTransactionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionRepositoryTestSuite))
}
