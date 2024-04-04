package mapper

import (
	dbmodel "github.com/illenko/transactions-service/internal/model"
	"github.com/illenko/transactions-service/pkg/model"
	"github.com/samber/lo"
	"log/slog"
)

type AnalyticMapper interface {
	ToResponse(items []dbmodel.AnalyticItem) model.AnalyticResponse
	ToDayResponse(items []dbmodel.DateAnalyticItem) model.AnalyticResponse
	ToMonthResponse(items []dbmodel.DateAnalyticItem) model.AnalyticResponse
}

type analyticMapper struct {
	log *slog.Logger
}

func NewAnalyticMapper(log *slog.Logger) AnalyticMapper {
	return &analyticMapper{log: log}
}

func (m *analyticMapper) ToResponse(items []dbmodel.AnalyticItem) model.AnalyticResponse {
	return model.AnalyticResponse{
		Count:  lo.Sum(lo.Map(items, func(item dbmodel.AnalyticItem, _ int) int { return item.Count })),
		Amount: lo.Sum(lo.Map(items, func(item dbmodel.AnalyticItem, _ int) float64 { return item.Amount })),
		Groups: lo.Map(items, func(item dbmodel.AnalyticItem, _ int) model.AnalyticGroup { return m.toGroup(item) }),
	}
}

func (m *analyticMapper) ToDayResponse(items []dbmodel.DateAnalyticItem) model.AnalyticResponse {
	return m.toDatedResponse(items, m.toDayGroup)
}

func (m *analyticMapper) ToMonthResponse(items []dbmodel.DateAnalyticItem) model.AnalyticResponse {
	return m.toDatedResponse(items, m.toMonthGroup)
}
func (m *analyticMapper) toDatedResponse(items []dbmodel.DateAnalyticItem, formatter datedFormatter) model.AnalyticResponse {
	return model.AnalyticResponse{
		Count:  lo.Sum(lo.Map(items, func(item dbmodel.DateAnalyticItem, _ int) int { return item.Count })),
		Amount: lo.Sum(lo.Map(items, func(item dbmodel.DateAnalyticItem, _ int) float64 { return item.Amount })),
		Groups: lo.Map(items, func(item dbmodel.DateAnalyticItem, _ int) model.AnalyticGroup { return formatter(item) }),
	}
}

type datedFormatter func(dbmodel.DateAnalyticItem) model.AnalyticGroup

func (m *analyticMapper) toGroup(item dbmodel.AnalyticItem) (group model.AnalyticGroup) {
	return model.AnalyticGroup{
		Name:   item.Name,
		Count:  item.Count,
		Amount: item.Amount,
	}
}

func (m *analyticMapper) toMonthGroup(item dbmodel.DateAnalyticItem) model.AnalyticGroup {
	return model.AnalyticGroup{
		Name:   item.Date.Month().String()[:3],
		Count:  item.Count,
		Amount: item.Amount,
	}
}

func (m *analyticMapper) toDayGroup(item dbmodel.DateAnalyticItem) model.AnalyticGroup {
	return model.AnalyticGroup{
		Name:   item.Date.Format("01-02-2006"),
		Count:  item.Count,
		Amount: item.Amount,
	}
}
