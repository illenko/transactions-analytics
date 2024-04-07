package mapper

import (
	dbmodel "github.com/illenko/analytics-service/internal/model"
	"github.com/illenko/analytics-service/pkg/model"
	"github.com/samber/lo"
	"log/slog"
)

type AnalyticsMapper interface {
	ToResponse(items []dbmodel.AnalyticsItem) model.AnalyticsResponse
	ToDayResponse(items []dbmodel.DateAnalyticsItem) model.AnalyticsResponse
	ToMonthResponse(items []dbmodel.DateAnalyticsItem) model.AnalyticsResponse
}

type analyticsMapper struct {
	log *slog.Logger
}

func NewAnalyticsMapper(log *slog.Logger) AnalyticsMapper {
	return &analyticsMapper{log: log}
}

func (m *analyticsMapper) ToResponse(items []dbmodel.AnalyticsItem) model.AnalyticsResponse {
	return model.AnalyticsResponse{
		Count:  lo.Sum(lo.Map(items, func(item dbmodel.AnalyticsItem, _ int) int { return item.Count })),
		Amount: lo.Sum(lo.Map(items, func(item dbmodel.AnalyticsItem, _ int) float64 { return item.Amount })),
		Groups: lo.Map(items, func(item dbmodel.AnalyticsItem, _ int) model.AnalyticsGroup { return m.toGroup(item) }),
	}
}

func (m *analyticsMapper) ToDayResponse(items []dbmodel.DateAnalyticsItem) model.AnalyticsResponse {
	return m.toDatedResponse(items, m.toDayGroup)
}

func (m *analyticsMapper) ToMonthResponse(items []dbmodel.DateAnalyticsItem) model.AnalyticsResponse {
	return m.toDatedResponse(items, m.toMonthGroup)
}
func (m *analyticsMapper) toDatedResponse(items []dbmodel.DateAnalyticsItem, formatter datedFormatter) model.AnalyticsResponse {
	return model.AnalyticsResponse{
		Count:  lo.Sum(lo.Map(items, func(item dbmodel.DateAnalyticsItem, _ int) int { return item.Count })),
		Amount: lo.Sum(lo.Map(items, func(item dbmodel.DateAnalyticsItem, _ int) float64 { return item.Amount })),
		Groups: lo.Map(items, func(item dbmodel.DateAnalyticsItem, _ int) model.AnalyticsGroup { return formatter(item) }),
	}
}

type datedFormatter func(dbmodel.DateAnalyticsItem) model.AnalyticsGroup

func (m *analyticsMapper) toGroup(item dbmodel.AnalyticsItem) (group model.AnalyticsGroup) {
	return model.AnalyticsGroup{
		Name:   item.Name,
		Count:  item.Count,
		Amount: item.Amount,
	}
}

func (m *analyticsMapper) toMonthGroup(item dbmodel.DateAnalyticsItem) model.AnalyticsGroup {
	return model.AnalyticsGroup{
		Name:   item.Date.Month().String()[:3],
		Count:  item.Count,
		Amount: item.Amount,
	}
}

func (m *analyticsMapper) toDayGroup(item dbmodel.DateAnalyticsItem) model.AnalyticsGroup {
	return model.AnalyticsGroup{
		Name:   item.Date.Format("01-02-2006"),
		Count:  item.Count,
		Amount: item.Amount,
	}
}
