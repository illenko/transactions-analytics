package database

import (
	"github.com/google/uuid"
	"github.com/illenko/analytics-service/internal/model"
	"gorm.io/gorm"
	"log/slog"
)

type TransactionRepository interface {
	FindAll() (transactions []model.Transaction, err error)
	FindById(id uuid.UUID) (transaction *model.Transaction, err error)
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

func (t *transactionRepository) FindById(id uuid.UUID) (*model.Transaction, error) {
	var transaction model.Transaction
	result := t.db.First(&transaction, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &transaction, nil
}
