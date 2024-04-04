package transaction

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log/slog"
)

type Repository interface {
	FindAll() (transactions []Entity, err error)
	FindById(id uuid.UUID) (transaction Entity, err error)
}

type repository struct {
	log *slog.Logger
	db  *gorm.DB
}

func NewRepository(log *slog.Logger, db *gorm.DB) Repository {
	return &repository{
		log: log,
		db:  db,
	}
}

func (t *repository) FindAll() (transactions []Entity, err error) {
	result := t.db.Order("datetime").Find(&transactions)
	return transactions, result.Error
}

func (t *repository) FindById(id uuid.UUID) (transaction Entity, err error) {
	result := t.db.Find(&transaction, id)
	return transaction, result.Error
}
