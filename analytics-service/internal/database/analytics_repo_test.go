package database

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/illenko/analytics-service/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"testing"
	"time"
)

func setupAnalyticsRepo(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, AnalyticsRepository) {
	sqlDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	log := slog.Default()
	repo := NewAnalyticsRepository(log, gormDB)

	return gormDB, mock, repo
}

func TestFind(t *testing.T) {
	db, mock, repo := setupAnalyticsRepo(t)
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	testCases := []struct {
		name           string
		group          string
		positiveAmount bool
		rows           *sqlmock.Rows
		expectedResult []model.AnalyticsItem
		expectedError  error
	}{
		{
			name:           "Returns rows, group by name, positive amount",
			group:          "name",
			positiveAmount: true,
			rows:           sqlmock.NewRows([]string{"name", "count", "amount"}).AddRow("test", 1, 100.0),
			expectedResult: []model.AnalyticsItem{
				{
					Name:   "test",
					Count:  1,
					Amount: 100.0,
				},
			},
			expectedError: nil,
		},
		{
			name:           "Returns no rows, group by name, positive amount",
			group:          "name",
			positiveAmount: true,
			rows:           sqlmock.NewRows([]string{"name", "count", "amount"}),
			expectedResult: []model.AnalyticsItem{},
			expectedError:  nil,
		},
		{
			name:           "Returns rows, group by name, negative amount",
			group:          "name",
			positiveAmount: false,
			rows:           sqlmock.NewRows([]string{"name", "count", "amount"}).AddRow("test", 1, -100.0),
			expectedResult: []model.AnalyticsItem{
				{
					Name:   "test",
					Count:  1,
					Amount: -100.0,
				},
			},
			expectedError: nil,
		},
		{
			name:           "Returns no rows, group by name, negative amount",
			group:          "name",
			positiveAmount: false,
			rows:           sqlmock.NewRows([]string{"name", "count", "amount"}),
			expectedResult: []model.AnalyticsItem{},
			expectedError:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			query := `SELECT ` + tc.group + ` as name, count(id), SUM(amount) as amount FROM "transactions" WHERE amount ` + amountCondition(tc.positiveAmount) + ` GROUP BY "name" ORDER BY amount` + orderDirection(tc.positiveAmount)
			mock.ExpectQuery(query).WillReturnRows(tc.rows)

			result, err := repo.Find(tc.group, tc.positiveAmount)

			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestFindByDates(t *testing.T) {
	db, mock, repo := setupAnalyticsRepo(t)
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	testCases := []struct {
		name           string
		positiveAmount bool
		unit           string
		rows           *sqlmock.Rows
		expectedResult []model.DateAnalyticsItem
		expectedError  error
	}{
		{
			name:           "Returns rows, positive amount, unit day",
			positiveAmount: true,
			unit:           "day",
			rows:           sqlmock.NewRows([]string{"date", "count", "amount"}).AddRow(mustParseTime("2006-01-02", "2022-01-01"), 1, 100.0),
			expectedResult: []model.DateAnalyticsItem{
				{
					Date:   mustParseTime("2006-01-02", "2022-01-01"),
					Count:  1,
					Amount: 100.0,
				},
			},
			expectedError: nil,
		},
		{
			name:           "Returns no rows, positive amount, unit day",
			positiveAmount: true,
			unit:           "day",
			rows:           sqlmock.NewRows([]string{"date", "count", "amount"}),
			expectedResult: []model.DateAnalyticsItem{},
			expectedError:  nil,
		},
		{
			name:           "Returns rows, negative amount, unit day",
			positiveAmount: false,
			unit:           "day",
			rows:           sqlmock.NewRows([]string{"date", "count", "amount"}).AddRow(mustParseTime("2006-01-02", "2022-01-01"), 1, -100.0),
			expectedResult: []model.DateAnalyticsItem{
				{
					Date:   mustParseTime("2006-01-02", "2022-01-01"),
					Count:  1,
					Amount: -100.0,
				},
			},
			expectedError: nil,
		},
		{
			name:           "Returns no rows, negative amount, unit day",
			positiveAmount: false,
			unit:           "day",
			rows:           sqlmock.NewRows([]string{"date", "count", "amount"}),
			expectedResult: []model.DateAnalyticsItem{},
			expectedError:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			query := `SELECT DATE_TRUNC('` + tc.unit + `', datetime) AS date, count(amount) as count, SUM(amount) AS amount FROM "transactions" WHERE amount ` + amountCondition(tc.positiveAmount) + ` GROUP BY "date" ORDER BY date`
			mock.ExpectQuery(query).WillReturnRows(tc.rows)

			result, err := repo.FindByDates(tc.positiveAmount, tc.unit)

			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestFindByDatesCumulative(t *testing.T) {
	db, mock, repo := setupAnalyticsRepo(t)
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	testCases := []struct {
		name           string
		positiveAmount bool
		unit           string
		rows           *sqlmock.Rows
		expectedResult []model.DateAnalyticsItem
		expectedError  error
	}{
		{
			name:           "Returns rows, positive amount, unit day",
			positiveAmount: true,
			unit:           "day",
			rows:           sqlmock.NewRows([]string{"date", "count", "amount"}).AddRow(mustParseTime("2006-01-02", "2022-01-01"), 1, 100.0),
			expectedResult: []model.DateAnalyticsItem{
				{
					Date:   mustParseTime("2006-01-02", "2022-01-01"),
					Count:  1,
					Amount: 100.0,
				},
			},
			expectedError: nil,
		},
		{
			name:           "Returns no rows, positive amount, unit day",
			positiveAmount: true,
			unit:           "day",
			rows:           sqlmock.NewRows([]string{"date", "count", "amount"}),
			expectedResult: []model.DateAnalyticsItem{},
			expectedError:  nil,
		},
		{
			name:           "Returns rows, negative amount, unit day",
			positiveAmount: false,
			unit:           "day",
			rows:           sqlmock.NewRows([]string{"date", "count", "amount"}).AddRow(mustParseTime("2006-01-02", "2022-01-01"), 1, -100.0),
			expectedResult: []model.DateAnalyticsItem{
				{
					Date:   mustParseTime("2006-01-02", "2022-01-01"),
					Count:  1,
					Amount: -100.0,
				},
			},
			expectedError: nil,
		},
		{
			name:           "Returns no rows, negative amount, unit day",
			positiveAmount: false,
			unit:           "day",
			rows:           sqlmock.NewRows([]string{"date", "count", "amount"}),
			expectedResult: []model.DateAnalyticsItem{},
			expectedError:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			query := `SELECT DATE_TRUNC('` + tc.unit + `', datetime) AS date, SUM(count(amount)) OVER (ORDER BY DATE_TRUNC('` + tc.unit + `', datetime)) AS count, SUM(SUM(amount)) OVER (ORDER BY DATE_TRUNC('` + tc.unit + `', datetime)) AS amount FROM "transactions" WHERE amount ` + amountCondition(tc.positiveAmount) + ` GROUP BY "date" ORDER BY date`
			mock.ExpectQuery(query).WillReturnRows(tc.rows)

			result, err := repo.FindByDatesCumulative(tc.positiveAmount, tc.unit)

			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func amountCondition(positiveAmount bool) string {
	if positiveAmount {
		return "> 0"
	} else {
		return "< 0"
	}
}

func orderDirection(positiveAmount bool) string {
	if positiveAmount {
		return " desc"
	} else {
		return " asc"
	}
}

// mustParseTime is a helper function that parses a time string and panics if the string cannot be parsed.
// It's useful in test code when you want to create a time.Time value from a string and you know the string is a valid time.
func mustParseTime(layout, value string) time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return t
}
