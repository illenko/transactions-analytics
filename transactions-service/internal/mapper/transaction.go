package mapper

import (
	dbmodel "github.com/illenko/transactions-service/internal/model"
	"github.com/illenko/transactions-service/pkg/model"
	"github.com/samber/lo"
	"log/slog"
)

type TransactionMapper interface {
	ToResponse(entity dbmodel.Transaction) model.TransactionResponse
	ToResponseList(entities []dbmodel.Transaction) []model.TransactionResponse
	ToStatisticsResponse(income []dbmodel.StatisticsBy,
		expenses []dbmodel.StatisticsBy,
		incomeDateAmounts []dbmodel.DateAmount,
		expensesDateAmounts []dbmodel.DateAmount) model.StatisticsResponse
	ToMonthAmounts(expenses []dbmodel.DateAmount) []model.MonthAmount
}

type transactionMapper struct {
	log *slog.Logger
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

func (t *transactionMapper) ToStatisticsResponse(income []dbmodel.StatisticsBy,
	expenses []dbmodel.StatisticsBy,
	incomeDateAmounts []dbmodel.DateAmount,
	expensesDateAmounts []dbmodel.DateAmount) model.StatisticsResponse {
	return model.StatisticsResponse{
		Income:   t.toStatistics(income, incomeDateAmounts),
		Expenses: t.toStatistics(expenses, expensesDateAmounts),
	}
}

func (t *transactionMapper) ToMonthAmounts(expenses []dbmodel.DateAmount) []model.MonthAmount {
	return lo.Map(expenses, func(item dbmodel.DateAmount, _ int) model.MonthAmount {
		return model.MonthAmount{
			Month:  item.Date.Month().String()[:3],
			Amount: item.Amount,
		}
	})
}

func (t *transactionMapper) toDateAmounts(dateAmounts []dbmodel.DateAmount) []model.DateAmount {
	return lo.Map(dateAmounts, func(item dbmodel.DateAmount, _ int) model.DateAmount {
		return model.DateAmount{
			Date:   item.Date.Format("01-02-2006"),
			Amount: item.Amount,
		}
	})
}

func (t *transactionMapper) toStatistics(entities []dbmodel.StatisticsBy, dateAmounts []dbmodel.DateAmount) model.Statistics {
	return model.Statistics{
		Count:       lo.Sum(lo.Map(entities, func(item dbmodel.StatisticsBy, _ int) int { return item.Count })),
		Amount:      lo.Sum(lo.Map(entities, func(item dbmodel.StatisticsBy, _ int) float64 { return item.Amount })),
		Groups:      lo.Map(entities, func(item dbmodel.StatisticsBy, _ int) model.StatisticsGroup { return t.toStatisticsGroup(item) }),
		DateAmounts: t.toDateAmounts(dateAmounts),
	}
}

func (t *transactionMapper) toStatisticsGroup(entity dbmodel.StatisticsBy) model.StatisticsGroup {
	return model.StatisticsGroup{
		Name:   entity.Name,
		Count:  entity.Count,
		Amount: entity.Amount,
	}
}
