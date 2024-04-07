package service

import (
	"context"
	"github.com/illenko/analytics-service/internal/mock"
	dbmodel "github.com/illenko/analytics-service/internal/model"
	"github.com/illenko/analytics-service/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/brianvoe/gofakeit/v6"
	"log/slog"
	"testing"
)

func TestAnalytics(t *testing.T) {

	cases := []struct {
		direction string
		group     string
	}{
		{
			direction: "income",
			group:     "category",
		},
		{
			direction: "income",
			group:     "merchant",
		},
		{
			direction: "expenses",
			group:     "merchant",
		},
		{
			direction: "expenses",
			group:     "category",
		},
	}
	for _, tc := range cases {
		t.Run(tc.direction+" "+tc.group, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			logger := slog.Default()
			repository := mock.NewMockAnalyticsRepository(ctrl)
			mapper := mock.NewMockAnalyticsMapper(ctrl)

			var analyticsItems []dbmodel.AnalyticsItem
			gofakeit.Slice(&analyticsItems)

			var expected model.AnalyticsResponse
			gofakeit.Slice(&expected)

			repository.EXPECT().Find(tc.group, tc.direction == "income").Return(analyticsItems, nil)
			mapper.EXPECT().ToResponse(analyticsItems).Return(expected)

			sut := NewAnalyticsService(logger, repository, mapper)

			actual, err := sut.Analytics(context.Background(), tc.direction, tc.group)

			require.NoError(t, err)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestAnalyticsByDates(t *testing.T) {
	cases := []struct {
		direction   string
		unit        string
		calculation string
	}{
		{
			direction:   "income",
			unit:        "day",
			calculation: "absolute",
		},
		{
			direction:   "income",
			unit:        "month",
			calculation: "absolute",
		},
		{
			direction:   "income",
			unit:        "day",
			calculation: "cumulative",
		},
		{
			direction:   "income",
			unit:        "month",
			calculation: "cumulative",
		},
		{
			direction:   "expenses",
			unit:        "day",
			calculation: "absolute",
		},
		{
			direction:   "expenses",
			unit:        "month",
			calculation: "absolute",
		},
		{
			direction:   "expenses",
			unit:        "day",
			calculation: "cumulative",
		},
		{
			direction:   "expenses",
			unit:        "month",
			calculation: "cumulative",
		},
	}
	for _, tc := range cases {
		t.Run(tc.direction+" "+tc.unit+" "+tc.calculation, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			logger := slog.Default()
			repository := mock.NewMockAnalyticsRepository(ctrl)
			mapper := mock.NewMockAnalyticsMapper(ctrl)

			var analyticsItems []dbmodel.DateAnalyticsItem
			gofakeit.Slice(&analyticsItems)

			var expected model.AnalyticsResponse
			gofakeit.Slice(&expected)

			if tc.calculation == "absolute" {
				repository.EXPECT().FindByDates(tc.direction == "income", tc.unit).Return(analyticsItems, nil)
			} else if tc.calculation == "cumulative" {
				repository.EXPECT().FindByDatesCumulative(tc.direction == "income", tc.unit).Return(analyticsItems, nil)
			}

			if tc.unit == "day" {
				mapper.EXPECT().ToDayResponse(analyticsItems).Return(expected)
			} else {
				mapper.EXPECT().ToMonthResponse(analyticsItems).Return(expected)
			}

			sut := NewAnalyticsService(logger, repository, mapper)

			actual, err := sut.AnalyticsByDates(context.Background(), tc.direction, tc.unit, tc.calculation)

			require.NoError(t, err)
			assert.Equal(t, expected, actual)
		})
	}
}
