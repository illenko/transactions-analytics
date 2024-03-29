package service

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/illenko/transactions-service/internal/database"
	"github.com/illenko/transactions-service/internal/mapper"
	"github.com/illenko/transactions-service/pkg/model"
	"log/slog"
)

type TransactionService interface {
	FindAll(ctx context.Context) ([]model.TransactionResponse, error)
	Statistics(ctx context.Context, by string) (model.StatisticsResponse, error)
	MerchantExpenses(ctx context.Context, by string) ([]model.MonthAmount, error)
	FindById(ctx context.Context, id uuid.UUID) (model.TransactionResponse, error)
}

type transactionService struct {
	log    *slog.Logger
	repo   database.TransactionRepository
	mapper mapper.TransactionMapper
}

func NewTransactionService(log *slog.Logger, repo database.TransactionRepository, mapper mapper.TransactionMapper) TransactionService {
	return &transactionService{
		log:    log,
		repo:   repo,
		mapper: mapper,
	}
}

func (t *transactionService) FindAll(ctx context.Context) ([]model.TransactionResponse, error) {
	t.log.InfoContext(ctx, "Retrieving all transactions")
	transactionEntities, err := t.repo.FindAll()
	if err != nil {
		t.log.ErrorContext(ctx, "When retrieving all transactions")
		return nil, err
	}
	t.log.InfoContext(ctx, fmt.Sprintf("Found: %v", spew.Sdump(transactionEntities)))
	return t.mapper.ToResponseList(transactionEntities), nil

}

func (t *transactionService) FindById(ctx context.Context, id uuid.UUID) (model.TransactionResponse, error) {
	t.log.InfoContext(ctx, "Retrieving transaction by id")
	transaction, err := t.repo.FindById(id)
	if err != nil {
		t.log.ErrorContext(ctx, "When retrieving transaction by id")
		return model.TransactionResponse{}, err
	}
	t.log.InfoContext(ctx, fmt.Sprintf("Found by id: %v", spew.Sdump(transaction)))
	return t.mapper.ToResponse(transaction), nil
}

func (t *transactionService) Statistics(ctx context.Context, by string) (model.StatisticsResponse, error) {
	t.log.InfoContext(ctx, fmt.Sprintf("Retrieving transaction statistics by: %v", by))
	income, err := t.repo.Statistics(by, true)
	if err != nil {
		t.log.ErrorContext(ctx, "When retrieving income transaction statistics")
		return model.StatisticsResponse{}, err
	}
	incomeDateAmounts, err := t.repo.DateAmounts(true)
	if err != nil {
		t.log.ErrorContext(ctx, "When retrieving income date amounts")
		return model.StatisticsResponse{}, err
	}
	expenses, err := t.repo.Statistics(by, false)
	if err != nil {
		t.log.ErrorContext(ctx, "When retrieving expenses transaction statistics")
		return model.StatisticsResponse{}, err
	}
	expensesDateAmounts, err := t.repo.DateAmounts(false)
	if err != nil {
		t.log.ErrorContext(ctx, "When retrieving expenses date amounts")
		return model.StatisticsResponse{}, err
	}
	return t.mapper.ToStatisticsResponse(income, expenses, incomeDateAmounts, expensesDateAmounts), nil
}

func (t *transactionService) MerchantExpenses(ctx context.Context, merchant string) ([]model.MonthAmount, error) {
	t.log.InfoContext(ctx, fmt.Sprintf("Retrieving merchant expenses by: %v", merchant))
	expenses, err := t.repo.MerchantExpenses(merchant)
	if err != nil {
		t.log.ErrorContext(ctx, "When merchant expenses")
		return []model.MonthAmount{}, err
	}
	return t.mapper.ToMonthAmounts(expenses), nil
}
