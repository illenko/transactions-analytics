package mapper

import (
	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	dbModel "github.com/illenko/analytics-service/internal/model"
	"github.com/illenko/analytics-service/pkg/model"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
	"time"
)

func TestToTransactionResponse(t *testing.T) {
	log := slog.Default()
	mapper := NewTransactionMapper(log)

	gofakeit.Seed(0)
	item := dbModel.Transaction{
		ID:       uuid.MustParse(gofakeit.UUID()),
		Datetime: time.Now(),
		Amount:   gofakeit.Price(1, 1000),
		Category: gofakeit.Word(),
		Merchant: gofakeit.Company(),
	}

	expected := model.TransactionResponse{
		ID:       item.ID,
		Datetime: item.Datetime,
		Amount:   item.Amount,
		Category: item.Category,
		Merchant: item.Merchant,
	}

	result := mapper.ToResponse(item)

	assert.Equal(t, expected, result, "The two values should be the same.")
}

func TestToResponses(t *testing.T) {
	log := slog.Default()
	mapper := NewTransactionMapper(log)

	gofakeit.Seed(0)
	items := []dbModel.Transaction{
		{
			ID:       uuid.MustParse(gofakeit.UUID()),
			Datetime: time.Now(),
			Amount:   gofakeit.Price(1, 1000),
			Category: gofakeit.Word(),
			Merchant: gofakeit.Company(),
		},
		{
			ID:       uuid.MustParse(gofakeit.UUID()),
			Datetime: time.Now(),
			Amount:   gofakeit.Price(1, 1000),
			Category: gofakeit.Word(),
			Merchant: gofakeit.Company(),
		},
	}

	expected := []model.TransactionResponse{
		{
			ID:       items[0].ID,
			Datetime: items[0].Datetime,
			Amount:   items[0].Amount,
			Category: items[0].Category,
			Merchant: items[0].Merchant,
		},
		{
			ID:       items[1].ID,
			Datetime: items[1].Datetime,
			Amount:   items[1].Amount,
			Category: items[1].Category,
			Merchant: items[1].Merchant,
		},
	}

	result := mapper.ToResponses(items)

	assert.Equal(t, expected, result, "The two values should be the same.")
}
