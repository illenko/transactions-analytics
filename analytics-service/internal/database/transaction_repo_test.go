package database

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/illenko/analytics-service/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"testing"
	"time"
)

func setupTransactionRepo(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, TransactionRepository) {
	sqlDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	log := slog.Default()
	repo := NewTransactionRepository(log, gormDB)

	return gormDB, mock, repo
}

func TestTransactionRepository_FindAll(t *testing.T) {
	db, mock, repo := setupTransactionRepo(t)
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	testCases := []struct {
		name           string
		rows           *sqlmock.Rows
		expectedLength int
		expectedError  error
	}{
		{
			name: "Returns multiple rows",
			rows: sqlmock.NewRows([]string{"id", "datetime", "amount"}).
				AddRow(uuid.New(), time.Now(), 100.0).
				AddRow(uuid.New(), time.Now(), 200.0),
			expectedLength: 2,
			expectedError:  nil,
		},
		{
			name: "Returns single row",
			rows: sqlmock.NewRows([]string{"id", "datetime", "amount"}).
				AddRow(uuid.New(), time.Now(), 100.0),
			expectedLength: 1,
			expectedError:  nil,
		},
		{
			name:           "Returns no rows",
			rows:           sqlmock.NewRows([]string{"id", "datetime", "amount"}),
			expectedLength: 0,
			expectedError:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock.ExpectQuery(`SELECT * FROM "transactions" ORDER BY datetime`).WillReturnRows(tc.rows)

			result, err := repo.FindAll()

			assert.Equal(t, tc.expectedError, err)
			assert.Len(t, result, tc.expectedLength)
		})
	}
}

func TestTransactionRepository_FindById(t *testing.T) {
	db, mock, repo := setupTransactionRepo(t)
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	id := uuid.New()
	now := time.Now()

	testCases := []struct {
		name           string
		id             uuid.UUID
		row            *sqlmock.Rows
		expectedResult model.Transaction
		expectedError  error
	}{
		{
			name: "Returns a row",
			id:   uuid.New(),
			row:  sqlmock.NewRows([]string{"id", "datetime", "amount"}).AddRow(id, now, 100.0), // Use id directly
			expectedResult: model.Transaction{
				ID:       id, // Use id directly
				Datetime: now,
				Amount:   100.0,
			},
			expectedError: nil,
		},
		{
			name:           "Returns no rows",
			id:             uuid.New(),
			row:            sqlmock.NewRows([]string{"id", "datetime", "amount"}),
			expectedResult: model.Transaction{},
			expectedError:  gorm.ErrRecordNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			limit := 1
			mock.ExpectQuery(`SELECT * FROM "transactions" WHERE "transactions"."id" = $1 ORDER BY "transactions"."id" LIMIT $2`).WithArgs(tc.id, limit).WillReturnRows(tc.row)

			result, err := repo.FindById(tc.id)

			assert.Equal(t, tc.expectedError, err)
			if err == nil {
				assert.Equal(t, tc.expectedResult.ID, result.ID)
			}
		})
	}
}
