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
	MerchantExpenses(merchant string) (expenses []model.DateAmount, err error)
	DateAmounts(positiveAmount bool) (expenses []model.DateAmount, err error)
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
	result := t.db.Select(groupKey + " as name, count(id), sum(amount) as amount").
		Table("transactions").
		Where(t.whereAmount(positiveAmount)).
		Group(groupKey).
		Order("amount" + t.orderDirection(positiveAmount)).
		Scan(&statistics)
	return statistics, result.Error
}

func (t *transactionRepository) MerchantExpenses(merchant string) (expenses []model.DateAmount, err error) {
	result := t.db.Select("DATE_TRUNC('month', datetime) AS date, SUM(amount) AS amount").
		Table("transactions").
		Where("merchant = ? and datetime > CURRENT_DATE - INTERVAL '6 months'", merchant).
		Group("month").
		Order("month").
		Scan(&expenses)
	return expenses, result.Error
}

func (t *transactionRepository) DateAmounts(positiveAmount bool) (dateAmounts []model.DateAmount, err error) {
	result := t.db.Select("DATE(datetime) as date, sum(sum(amount)) over (order by DATE(datetime)) as amount").
		Table("transactions").
		Where(t.whereAmount(positiveAmount)).
		Group("date").
		Order("date").
		Scan(&dateAmounts)

	return dateAmounts, result.Error
}

func (t *transactionRepository) whereAmount(positiveAmount bool) (where string) {
	if positiveAmount {
		return "amount > 0"
	} else {
		return "amount < 0"
	}
}

func (t *transactionRepository) orderDirection(positiveAmount bool) (order string) {
	if positiveAmount {
		return " desc"
	} else {
		return " asc"
	}
}
