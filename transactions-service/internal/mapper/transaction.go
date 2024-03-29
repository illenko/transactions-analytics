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
	ToStatisticsResponse(income []dbmodel.StatisticsBy, expenses []dbmodel.StatisticsBy) model.StatisticsResponse
	ToMonthExpenses(expenses []dbmodel.MerchantExpense) []model.MonthExpense
}

type transactionMapper struct {
	log  *slog.Logger
	repo database.TransactionRepository
}

func NewTransactionMapper(log *slog.Logger) TransactionMapper {
	return &transactionMapper{log: log}
}

func (t *transactionMapper) ToResponse(entity dbmodel.Transaction) model.TransactionResponse {
	return model.TransactionResponse{
		ID:       entity.ID,
		Datetime: entity.Datetime,
		Amount:   entity.Amount,
		Category: entity.Category,
		Merchant: entity.Merchant,
	}
}

func (t *transactionMapper) ToResponseList(entities []dbmodel.Transaction) []model.TransactionResponse {
	return lo.Map(entities, func(item dbmodel.Transaction, _ int) model.TransactionResponse { return t.ToResponse(item) })
}

func (t *transactionMapper) ToStatisticsResponse(income []dbmodel.StatisticsBy, expenses []dbmodel.StatisticsBy) model.StatisticsResponse {
	return model.StatisticsResponse{
		Income:   t.toStatistics(income),
		Expenses: t.toStatistics(expenses),
	}
}

func (t *transactionMapper) ToMonthExpenses(expenses []dbmodel.MerchantExpense) []model.MonthExpense {
	return lo.Map(expenses, func(item dbmodel.MerchantExpense, _ int) model.MonthExpense {
		return model.MonthExpense{
			Month:  item.Month.Month().String()[:3],
			Amount: item.Amount,
		}
	})
}

func (t *transactionMapper) toStatistics(entities []dbmodel.StatisticsBy) model.Statistics {
	return model.Statistics{
		Count:  lo.Sum(lo.Map(entities, func(item dbmodel.StatisticsBy, _ int) int { return item.Count })),
		Amount: lo.Sum(lo.Map(entities, func(item dbmodel.StatisticsBy, _ int) float64 { return item.Amount })),
		Groups: lo.Map(entities, func(item dbmodel.StatisticsBy, _ int) model.StatisticsGroup { return t.toStatisticsGroup(item) }),
	}
}

func (t *transactionMapper) toStatisticsGroup(entity dbmodel.StatisticsBy) model.StatisticsGroup {
	return model.StatisticsGroup{
		Name:   entity.Name,
		Count:  entity.Count,
		Amount: entity.Amount,
	}
}
