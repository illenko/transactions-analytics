package service

import (
	"context"
	"github.com/illenko/transactions-service/internal/database"
	"github.com/illenko/transactions-service/internal/mapper"
	dbmodel "github.com/illenko/transactions-service/internal/model"
	"github.com/illenko/transactions-service/pkg/model"
	"log/slog"
)

type AnalyticService interface {
	Analytic(ctx context.Context, analyticType string, groupBy string) (model.Analytic, error)
	AnalyticByDates(ctx context.Context, analyticType string, unit string, period int, category string, merchant string, valueType string) (analytic model.Analytic, err error)
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

func (s *analyticService) AnalyticByDates(ctx context.Context, analyticType string, unit string, period int, category string, merchant string, valueType string) (analytic model.Analytic, err error) {

	var analyticItems []dbmodel.DateAnalyticItem
	if valueType == "absolute" {
		analyticItems, err = s.repo.FindByDates(s.resolveAmount(analyticType), unit, period, category, merchant)
	} else {
		analyticItems, err = s.repo.FindByDatesCumulative(s.resolveAmount(analyticType), unit, period, category, merchant)
	}
	if err != nil {
		s.log.ErrorContext(ctx, "When retrieving income transaction statistics by dates")
		return model.Analytic{}, err
	}

	if unit == "day" {
		return s.mapper.ToDayResponse(analyticItems), nil
	} else {
		return s.mapper.ToMonthResponse(analyticItems), nil
	}
}

func (s *analyticService) resolveAmount(analyticType string) bool {
	return analyticType == "income"
}
