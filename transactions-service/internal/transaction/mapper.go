package transaction

import (
	"github.com/illenko/transactions-service/pkg/model"
	"github.com/samber/lo"
	"log/slog"
)

type Mapper interface {
	ToResponse(entity Entity) model.TransactionResponse
	ToResponses(entities []Entity) []model.TransactionResponse
}

type mapper struct {
	log *slog.Logger
}

func NewMapper(log *slog.Logger) Mapper {
	return &mapper{log: log}
}

func (t *mapper) ToResponse(entity Entity) model.TransactionResponse {
	return model.TransactionResponse{
		ID:       entity.ID,
		Datetime: entity.Datetime,
		Amount:   entity.Amount,
		Category: entity.Category,
		Merchant: entity.Merchant,
	}
}

func (t *mapper) ToResponses(entities []Entity) []model.TransactionResponse {
	return lo.Map(entities, func(item Entity, _ int) model.TransactionResponse {
		return t.ToResponse(item)
	})
}
