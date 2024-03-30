package service

import (
	"context"
	"github.com/illenko/transactions-service/internal/database"
	"github.com/illenko/transactions-service/internal/mapper"
	"github.com/illenko/transactions-service/pkg/model"
	"log/slog"
)

type AnalyticService interface {
	Analytic(ctx context.Context, analyticType string, groupBy string) (model.Analytic, error)
	AnalyticByDates(ctx context.Context, analyticType string, unit string,
		period int, category *string, merchant *string) (model.Analytic, error)
}

type analyticService struct {
	log    *slog.Logger
	repo   database.AnalyticRepository
	mapper mapper.AnalyticMapper
}

func NewAnalyticService(log *slog.Logger, repo database.AnalyticRepository, mapper mapper.AnalyticMapper) AnalyticService {
	return &analyticService{
		log:    log,
		repo:   repo,
		mapper: mapper,
	}
}

func (s *analyticService) Analytic(ctx context.Context, analyticType string, groupBy string) (model.Analytic, error) {
	analyticItems, err := s.repo.Find(groupBy, s.resolveAmount(analyticType))
	if err != nil {
		s.log.ErrorContext(ctx, "When retrieving income transaction statistics")
		return model.Analytic{}, err
	}
	return s.mapper.ToResponse(analyticItems), nil
}

func (s *analyticService) AnalyticByDates(ctx context.Context, analyticType string, unit string, period int, category *string, merchant *string) (model.Analytic, error) {
	//TODO implement me
	panic("implement me")
}

func (s *analyticService) resolveAmount(analyticType string) bool {
	return analyticType == "income"
}
