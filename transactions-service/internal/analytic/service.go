package analytic

import (
	"context"
	"github.com/illenko/transactions-service/pkg/model"
	"log/slog"
)

type Service interface {
	Analytic(ctx context.Context, direction string, group string) (model.AnalyticResponse, error)
	AnalyticByDates(ctx context.Context, direction string, unit string, calculation string) (analytic model.AnalyticResponse, err error)
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

func (s *service) Analytic(ctx context.Context, direction string, group string) (model.AnalyticResponse, error) {
	analyticItems, err := s.repo.Find(group, s.resolveAmount(direction))
	if err != nil {
		s.log.ErrorContext(ctx, "When retrieving income transaction statistics")
		return model.AnalyticResponse{}, err
	}
	return s.mapper.ToResponse(analyticItems), nil
}

func (s *service) AnalyticByDates(ctx context.Context, direction string, unit string, calculation string) (analytic model.AnalyticResponse, err error) {
	analyticItems, err := s.getAnalyticByDates(calculation, direction, unit)

	if err != nil {
		s.log.ErrorContext(ctx, "When retrieving income transaction statistics by dates")
		return model.AnalyticResponse{}, err
	}

	return s.resolveMappingByUnit(unit, analyticItems)
}

func (s *service) getAnalyticByDates(calculation string, direction string, unit string) ([]DateItem, error) {
	if calculation == "absolute" {
		return s.repo.FindByDates(s.resolveAmount(direction), unit)
	} else {
		return s.repo.FindByDatesCumulative(s.resolveAmount(direction), unit)
	}
}

func (s *service) resolveMappingByUnit(unit string, analyticItems []DateItem) (model.AnalyticResponse, error) {
	if unit == "day" {
		return s.mapper.ToDayResponse(analyticItems), nil
	} else {
		return s.mapper.ToMonthResponse(analyticItems), nil
	}
}

func (s *service) resolveAmount(direction string) bool {
	return direction == "income"
}
