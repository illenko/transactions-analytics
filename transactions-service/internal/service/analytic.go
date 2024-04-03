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
	Analytic(ctx context.Context, direction string, group string) (model.AnalyticResponse, error)
	AnalyticByDates(ctx context.Context, direction string, unit string, calculation string) (analytic model.AnalyticResponse, err error)
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

func (s *analyticService) Analytic(ctx context.Context, direction string, group string) (model.AnalyticResponse, error) {
	analyticItems, err := s.repo.Find(group, s.resolveAmount(direction))
	if err != nil {
		s.log.ErrorContext(ctx, "When retrieving income transaction statistics")
		return model.AnalyticResponse{}, err
	}
	return s.mapper.ToResponse(analyticItems), nil
}

func (s *analyticService) AnalyticByDates(ctx context.Context, direction string, unit string, calculation string) (analytic model.AnalyticResponse, err error) {
	analyticItems, err := s.getAnalyticByDates(calculation, direction, unit)

	if err != nil {
		s.log.ErrorContext(ctx, "When retrieving income transaction statistics by dates")
		return model.AnalyticResponse{}, err
	}

	return s.resolveMappingByUnit(unit, analyticItems)
}

func (s *analyticService) getAnalyticByDates(calculation string, direction string, unit string) ([]dbmodel.DateAnalyticItem, error) {
	if calculation == "absolute" {
		return s.repo.FindByDates(s.resolveAmount(direction), unit)
	} else {
		return s.repo.FindByDatesCumulative(s.resolveAmount(direction), unit)
	}
}

func (s *analyticService) resolveMappingByUnit(unit string, analyticItems []dbmodel.DateAnalyticItem) (model.AnalyticResponse, error) {
	if unit == "day" {
		return s.mapper.ToDayResponse(analyticItems), nil
	} else {
		return s.mapper.ToMonthResponse(analyticItems), nil
	}
}

func (s *analyticService) resolveAmount(direction string) bool {
	return direction == "income"
}
