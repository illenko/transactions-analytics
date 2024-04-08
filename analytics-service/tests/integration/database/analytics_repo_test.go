package database

import (
	"github.com/illenko/analytics-service/internal/database"
	"github.com/illenko/analytics-service/internal/model"
	"github.com/illenko/analytics-service/tests/integration"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type AnalyticsRepositoryTestSuite struct {
	integration.BaseIntegrationTestSuite
}

func (suite *AnalyticsRepositoryTestSuite) TestFind() {
	sut := database.NewAnalyticsRepository(suite.Log, suite.DB)

	cases := []struct {
		direction string
		group     string
		expected  []model.AnalyticsItem
	}{
		{
			direction: "income",
			group:     "category",
			expected: []model.AnalyticsItem{
				{
					Name:   "work",
					Count:  1,
					Amount: 104,
				},
				{
					Name:   "p2p",
					Count:  3,
					Amount: 85,
				},
			},
		},
		{
			direction: "income",
			group:     "merchant",
			expected: []model.AnalyticsItem{
				{
					Name:   "upwork",
					Count:  1,
					Amount: 104},
				{
					Name:   "paypal",
					Count:  1,
					Amount: 55,
				},
				{
					Name:   "stripe",
					Count:  1,
					Amount: 20,
				},
				{
					Name:   "revolut",
					Count:  1,
					Amount: 10,
				},
			},
		},
		{
			direction: "expenses",
			group:     "merchant",
			expected: []model.AnalyticsItem{
				{
					Name:   "uklon",
					Count:  6,
					Amount: -172,
				},
				{
					Name:   "glovo",
					Count:  1,
					Amount: -100},
				{
					Name:   "uber",
					Count:  1,
					Amount: -30,
				},
				{
					Name:   "jysk",
					Count:  1,
					Amount: -24,
				},
				{
					Name:   "mcdonalds",
					Count:  1,
					Amount: -20.05,
				},
				{Name: "paypal",
					Count:  1,
					Amount: -5,
				},
				{
					Name:   "kfc",
					Count:  1,
					Amount: -3.2,
				},
			},
		},
		{
			direction: "expenses",
			group:     "category",
			expected: []model.AnalyticsItem{
				{
					Name:   "taxi",
					Count:  7,
					Amount: -202,
				},
				{
					Name:   "food",
					Count:  3,
					Amount: -123.25,
				},
				{
					Name:   "house",
					Count:  1,
					Amount: -24,
				},
				{
					Name:   "p2p",
					Count:  1,
					Amount: -5,
				},
			},
		},
	}
	for _, tc := range cases {
		suite.Run("Find analytics items: "+tc.direction+" "+tc.group, func() {
			actual, err := sut.Find(tc.group, tc.direction == "income")
			suite.NoError(err)
			suite.Equal(tc.expected, actual)
		})
	}
}

func (suite *AnalyticsRepositoryTestSuite) TestFindByDates() {
	suite.Run("Find analytic date items", func() {
		repo := database.NewAnalyticsRepository(suite.Log, suite.DB)
		dateAnalyticsItems, err := repo.FindByDates(true, "day")
		suite.NoError(err)
		suite.Equal(dateAnalyticsItems,
			[]model.DateAnalyticsItem{
				{
					Date:   time.Date(2024, time.March, 27, 0, 0, 0, 0, time.UTC),
					Count:  2,
					Amount: 75,
				},
				{
					Date:   time.Date(2024, time.March, 28, 0, 0, 0, 0, time.UTC),
					Count:  1,
					Amount: 104,
				},
				{
					Date:   time.Date(2024, time.March, 30, 0, 0, 0, 0, time.UTC),
					Count:  1,
					Amount: 10,
				},
			},
		)
	})
}

func (suite *AnalyticsRepositoryTestSuite) TestFindByDatesCumulative() {
	suite.Run("Find analytic date items", func() {
		repo := database.NewAnalyticsRepository(suite.Log, suite.DB)
		dateAnalyticsItems, err := repo.FindByDatesCumulative(true, "day")
		suite.NoError(err)
		suite.Equal(dateAnalyticsItems,
			[]model.DateAnalyticsItem{
				{
					Date:   time.Date(2024, time.March, 27, 0, 0, 0, 0, time.UTC),
					Count:  2,
					Amount: 75,
				},
				{
					Date:   time.Date(2024, time.March, 28, 0, 0, 0, 0, time.UTC),
					Count:  3,
					Amount: 179,
				},
				{
					Date:   time.Date(2024, time.March, 30, 0, 0, 0, 0, time.UTC),
					Count:  4,
					Amount: 189,
				},
			},
		)
	})
}

func TestAnalyticsRepository(t *testing.T) {
	suite.Run(t, new(AnalyticsRepositoryTestSuite))
}
