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
	StatisticsBy(by string, income bool) (statistics []model.StatisticsBy, err error)
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
	result := t.db.Find(&transactions)
	return transactions, result.Error
}

func (t *transactionRepository) StatisticsBy(by string, income bool) (statistics []model.StatisticsBy, err error) {
	var query string
	if income {
		query = "select " + by + " as name, count(id), sum(amount) as amount from transactions where amount > 0 group by " + by + " order by amount desc"
	} else {
		query = "select " + by + " as name, count(id), sum(amount) as amount from transactions where amount < 0 group by " + by + " order by amount"
	}
	result := t.db.Raw(query).Scan(&statistics)
	return statistics, result.Error
}

func (t *transactionRepository) FindById(id uuid.UUID) (transaction model.Transaction, err error) {
	result := t.db.Find(&transaction, id)
	return transaction, result.Error
}
