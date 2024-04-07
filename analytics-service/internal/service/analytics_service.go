package service

import (
	"context"
	"github.com/illenko/analytics-service/internal/database"
	"github.com/illenko/analytics-service/internal/mapper"
	dbmodel "github.com/illenko/analytics-service/internal/model"
	"github.com/illenko/analytics-service/pkg/model"
	"log/slog"
)

type AnalyticsService interface {
	Analytics(ctx context.Context, direction string, group string) (model.AnalyticsResponse, error)
	AnalyticsByDates(ctx context.Context, direction string, unit string, calculation string) (analytics model.AnalyticsResponse, err error)
}

type analyticsService struct {
	log    *slog.Logger
	repo   database.AnalyticsRepository
	mapper mapper.AnalyticsMapper
}

func NewAnalyticsService(log *slog.Logger, repo database.AnalyticsRepository, mapper mapper.AnalyticsMapper) AnalyticsService {
	return &analyticsService{
		log:    log,
		repo:   repo,
		mapper: mapper,
	}
}

func (s *analyticsService) Analytics(ctx context.Context, direction string, group string) (model.AnalyticsResponse, error) {
	analyticsItems, err := s.repo.Find(group, s.resolveAmount(direction))
	if err != nil {
		s.log.ErrorContext(ctx, "When retrieving income transaction analytics")
		return model.AnalyticsResponse{}, err
	}
	return s.mapper.ToResponse(analyticsItems), nil
}

func (s *analyticsService) AnalyticsByDates(ctx context.Context, direction string, unit string, calculation string) (analytics model.AnalyticsResponse, err error) {
	analyticsItems, err := s.getAnalyticsByDates(calculation, direction, unit)

	if err != nil {
		s.log.ErrorContext(ctx, "When retrieving income transaction analytics by dates")
		return model.AnalyticsResponse{}, err
	}

	return s.resolveMappingByUnit(unit, analyticsItems)
}

func (s *analyticsService) getAnalyticsByDates(calculation string, direction string, unit string) ([]dbmodel.DateAnalyticsItem, error) {
	if calculation == "absolute" {
		return s.repo.FindByDates(s.resolveAmount(direction), unit)
	} else {
		return s.repo.FindByDatesCumulative(s.resolveAmount(direction), unit)
	}
}

func (s *analyticsService) resolveMappingByUnit(unit string, analyticsItems []dbmodel.DateAnalyticsItem) (model.AnalyticsResponse, error) {
	if unit == "day" {
		return s.mapper.ToDayResponse(analyticsItems), nil
	} else {
		return s.mapper.ToMonthResponse(analyticsItems), nil
	}
}

func (s *analyticsService) resolveAmount(direction string) bool {
	return direction == "income"
}
