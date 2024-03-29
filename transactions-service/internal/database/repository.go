package database

import (
	"github.com/google/uuid"
	"github.com/illenko/transactions-service/internal/model"
	"gorm.io/gorm"
	"log/slog"
)

type TransactionRepository interface {
	FindAll() (transactions []model.Transaction, err error)
	FindById(id uuid.UUID) (transaction model.Transaction, err error)
	Statistics(by string, income bool) (statistics []model.StatisticsBy, err error)
	MerchantExpenses(merchant string) (expenses []model.MerchantExpense, err error)
}

type transactionRepository struct {
	log *slog.Logger
	db  *gorm.DB
}

func NewTransactionRepository(log *slog.Logger, db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		log: log,
		db:  db,
	}
}

func (t *transactionRepository) FindAll() (transactions []model.Transaction, err error) {
	result := t.db.Order("datetime").Find(&transactions)
	return transactions, result.Error
}

func (t *transactionRepository) FindById(id uuid.UUID) (transaction model.Transaction, err error) {
	result := t.db.Find(&transaction, id)
	return transaction, result.Error
}

func (t *transactionRepository) Statistics(groupKey string, positiveAmount bool) (statistics []model.StatisticsBy, err error) {
	where, order := t.statisticsParameters(positiveAmount)

	result := t.db.Select(groupKey + " as name, count(id), sum(amount) as amount").Table("transactions").Where(where).Group(groupKey).Order("amount " + order).Scan(&statistics)

	return statistics, result.Error
}

func (t *transactionRepository) MerchantExpenses(merchant string) (expenses []model.MerchantExpense, err error) {
	result := t.db.Select("DATE_TRUNC('month', datetime) AS month, SUM(amount) AS amount").
		Table("transactions").Where("merchant = ? and datetime > CURRENT_DATE - INTERVAL '6 months'", merchant).
		Group("month").Order("month").Scan(&expenses)

	return expenses, result.Error
}

func (t *transactionRepository) statisticsParameters(positiveAmount bool) (where string, order string) {
	if positiveAmount {
		where = "amount > 0"
		order = "desc"
	} else {
		where = "amount < 0"
		order = "asc"
	}
	return
}
