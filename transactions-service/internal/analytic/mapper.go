package analytic

import (
	"github.com/illenko/transactions-service/pkg/model"
	"github.com/samber/lo"
	"log/slog"
)

type Mapper interface {
	ToResponse(items []Item) model.AnalyticResponse
	ToDayResponse(items []DateItem) model.AnalyticResponse
	ToMonthResponse(items []DateItem) model.AnalyticResponse
}

type mapper struct {
	log *slog.Logger
}

func NewMapper(log *slog.Logger) Mapper {
	return &mapper{log: log}
}

func (m *mapper) ToResponse(items []Item) model.AnalyticResponse {
	return model.AnalyticResponse{
		Count:  lo.Sum(lo.Map(items, func(item Item, _ int) int { return item.Count })),
		Amount: lo.Sum(lo.Map(items, func(item Item, _ int) float64 { return item.Amount })),
		Groups: lo.Map(items, func(item Item, _ int) model.AnalyticGroup { return m.toGroup(item) }),
	}
}

func (m *mapper) ToDayResponse(items []DateItem) model.AnalyticResponse {
	return m.toDatedResponse(items, m.toDayGroup)
}

func (m *mapper) ToMonthResponse(items []DateItem) model.AnalyticResponse {
	return m.toDatedResponse(items, m.toMonthGroup)
}
func (m *mapper) toDatedResponse(items []DateItem, formatter datedFormatter) model.AnalyticResponse {
	return model.AnalyticResponse{
		Count:  lo.Sum(lo.Map(items, func(item DateItem, _ int) int { return item.Count })),
		Amount: lo.Sum(lo.Map(items, func(item DateItem, _ int) float64 { return item.Amount })),
		Groups: lo.Map(items, func(item DateItem, _ int) model.AnalyticGroup { return formatter(item) }),
	}
}

type datedFormatter func(DateItem) model.AnalyticGroup

func (m *mapper) toGroup(item Item) (group model.AnalyticGroup) {
	return model.AnalyticGroup{
		Name:   item.Name,
		Count:  item.Count,
		Amount: item.Amount,
	}
}

func (m *mapper) toMonthGroup(item DateItem) model.AnalyticGroup {
	return model.AnalyticGroup{
		Name:   item.Date.Month().String()[:3],
		Count:  item.Count,
		Amount: item.Amount,
	}
}

func (m *mapper) toDayGroup(item DateItem) model.AnalyticGroup {
	return model.AnalyticGroup{
		Name:   item.Date.Format("01-02-2006"),
		Count:  item.Count,
		Amount: item.Amount,
	}
}
