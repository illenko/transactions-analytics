package service

import (
	"context"
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

func TestAnalytic(t *testing.T) {

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
			repository := mock.NewMockAnalyticRepository(ctrl)
			mapper := mock.NewMockAnalyticMapper(ctrl)

			var analyticItems []dbmodel.AnalyticItem
			gofakeit.Slice(&analyticItems)

			var expected model.AnalyticResponse
			gofakeit.Slice(&expected)

			repository.EXPECT().Find(tc.group, tc.direction == "income").Return(analyticItems, nil)
			mapper.EXPECT().ToResponse(analyticItems).Return(expected)

			sut := NewAnalyticService(logger, repository, mapper)

			actual, err := sut.Analytic(context.Background(), tc.direction, tc.group)

			require.NoError(t, err)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestAnalyticByDates(t *testing.T) {
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
			repository := mock.NewMockAnalyticRepository(ctrl)
			mapper := mock.NewMockAnalyticMapper(ctrl)

			var analyticItems []dbmodel.DateAnalyticItem
			gofakeit.Slice(&analyticItems)

			var expected model.AnalyticResponse
			gofakeit.Slice(&expected)

			if tc.calculation == "absolute" {
				repository.EXPECT().FindByDates(tc.direction == "income", tc.unit).Return(analyticItems, nil)
			} else if tc.calculation == "cumulative" {
				repository.EXPECT().FindByDatesCumulative(tc.direction == "income", tc.unit).Return(analyticItems, nil)
			}

			if tc.unit == "day" {
				mapper.EXPECT().ToDayResponse(analyticItems).Return(expected)
			} else {
				mapper.EXPECT().ToMonthResponse(analyticItems).Return(expected)
			}

			sut := NewAnalyticService(logger, repository, mapper)

			actual, err := sut.AnalyticByDates(context.Background(), tc.direction, tc.unit, tc.calculation)

			require.NoError(t, err)
			assert.Equal(t, expected, actual)
		})
	}
}
