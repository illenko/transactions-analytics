package mapper

import (
	"github.com/brianvoe/gofakeit"
	dbModel "github.com/illenko/analytics-service/internal/model"
	"github.com/illenko/analytics-service/pkg/model"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
	"time"
)

func TestToResponse(t *testing.T) {
	log := slog.Default()
	mapper := NewAnalyticsMapper(log)

	gofakeit.Seed(0)
	items := []dbModel.AnalyticsItem{
		{
			Name:   gofakeit.Name(),
			Count:  gofakeit.Number(1, 100),
			Amount: gofakeit.Price(1, 1000),
		},
		{
			Name:   gofakeit.Name(),
			Count:  gofakeit.Number(1, 100),
			Amount: gofakeit.Price(1, 1000),
		},
	}

	expected := model.AnalyticsResponse{
		Count:  items[0].Count + items[1].Count,
		Amount: items[0].Amount + items[1].Amount,
		Groups: []model.AnalyticsGroup{
			{
				Name:   items[0].Name,
				Count:  items[0].Count,
				Amount: items[0].Amount,
			},
			{
				Name:   items[1].Name,
				Count:  items[1].Count,
				Amount: items[1].Amount,
			},
		},
	}

	result := mapper.ToResponse(items)

	assert.Equal(t, expected, result, "The two values should be the same.")
}

func TestToDayResponse(t *testing.T) {
	log := slog.Default()
	mapper := NewAnalyticsMapper(log)

	gofakeit.Seed(0)
	items := []dbModel.DateAnalyticsItem{
		{
			Date:   time.Now(),
			Count:  gofakeit.Number(1, 100),
			Amount: gofakeit.Price(1, 1000),
		},
		{
			Date:   time.Now(),
			Count:  gofakeit.Number(1, 100),
			Amount: gofakeit.Price(1, 1000),
		},
	}

	expected := model.AnalyticsResponse{
		Count:  items[0].Count + items[1].Count,
		Amount: items[0].Amount + items[1].Amount,
		Groups: []model.AnalyticsGroup{
			{
				Name:   items[0].Date.Format("01-02-2006"),
				Count:  items[0].Count,
				Amount: items[0].Amount,
			},
			{
				Name:   items[1].Date.Format("01-02-2006"),
				Count:  items[1].Count,
				Amount: items[1].Amount,
			},
		},
	}

	result := mapper.ToDayResponse(items)

	assert.Equal(t, expected, result, "The two values should be the same.")
}

func TestToMonthResponse(t *testing.T) {
	log := slog.Default()
	mapper := NewAnalyticsMapper(log)

	gofakeit.Seed(0)
	items := []dbModel.DateAnalyticsItem{
		{
			Date:   time.Now(),
			Count:  gofakeit.Number(1, 100),
			Amount: gofakeit.Price(1, 1000),
		},
		{
			Date:   time.Now(),
			Count:  gofakeit.Number(1, 100),
			Amount: gofakeit.Price(1, 1000),
		},
	}

	expected := model.AnalyticsResponse{
		Count:  items[0].Count + items[1].Count,
		Amount: items[0].Amount + items[1].Amount,
		Groups: []model.AnalyticsGroup{
			{
				Name:   items[0].Date.Month().String()[:3],
				Count:  items[0].Count,
				Amount: items[0].Amount,
			},
			{
				Name:   items[1].Date.Month().String()[:3],
				Count:  items[1].Count,
				Amount: items[1].Amount,
			},
		},
	}

	result := mapper.ToMonthResponse(items)

	assert.Equal(t, expected, result, "The two values should be the same.")
}
