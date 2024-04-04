package transaction

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/illenko/transactions-service/pkg/model"
	"log/slog"
)

type Service interface {
	FindAll(ctx context.Context) ([]model.TransactionResponse, error)
	FindById(ctx context.Context, id uuid.UUID) (model.TransactionResponse, error)
}

type service struct {
	log    *slog.Logger
	repo   Repository
	mapper Mapper
}

func NewService(log *slog.Logger, repo Repository, mapper Mapper) Service {
	return &service{
		log:    log,
		repo:   repo,
		mapper: mapper,
	}
}

func (t *service) FindAll(ctx context.Context) ([]model.TransactionResponse, error) {
	t.log.InfoContext(ctx, "Retrieving all transactions")
	transactionEntities, err := t.repo.FindAll()
	if err != nil {
		t.log.ErrorContext(ctx, "When retrieving all transactions")
		return nil, err
	}
	t.log.InfoContext(ctx, fmt.Sprintf("Found: %v", spew.Sdump(transactionEntities)))
	return t.mapper.ToResponses(transactionEntities), nil

}

func (t *service) FindById(ctx context.Context, id uuid.UUID) (model.TransactionResponse, error) {
	t.log.InfoContext(ctx, "Retrieving transaction by id")
	transaction, err := t.repo.FindById(id)
	if err != nil {
		t.log.ErrorContext(ctx, "When retrieving transaction by id")
		return model.TransactionResponse{}, err
	}
	t.log.InfoContext(ctx, fmt.Sprintf("Found by id: %v", spew.Sdump(transaction)))
	return t.mapper.ToResponse(transaction), nil
}
