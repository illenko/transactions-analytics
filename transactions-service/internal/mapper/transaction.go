package mapper

import (
	"github.com/illenko/transactions-service/internal/database"
	dbmodel "github.com/illenko/transactions-service/internal/model"
	"github.com/illenko/transactions-service/pkg/model"
	"github.com/samber/lo"
	"log/slog"
)

type TransactionMapper interface {
	ToResponse(entity dbmodel.Transaction) model.TransactionResponse
	ToResponseList(entities []dbmodel.Transaction) []model.TransactionResponse
}

type transactionMapper struct {
	log  *slog.Logger
	repo database.TransactionRepository
}

func NewTransactionMapper(log *slog.Logger) TransactionMapper {
	return &transactionMapper{log: log}
}

func (t transactionMapper) ToResponse(entity dbmodel.Transaction) model.TransactionResponse {
	return model.TransactionResponse{
		ID:       entity.ID,
		Datetime: entity.Datetime,
		Amount:   entity.Amount,
		Category: entity.Category,
		Merchant: entity.Merchant,
	}
}

func (t transactionMapper) ToResponseList(entities []dbmodel.Transaction) []model.TransactionResponse {
	return lo.Map(entities, func(item dbmodel.Transaction, _ int) model.TransactionResponse { return t.ToResponse(item) })
}
