package database

import (
	"github.com/illenko/transactions-service/internal/model"
	"gorm.io/gorm"
	"log/slog"
)

type AnalyticRepository interface {
	Find(groupBy string, positiveAmount bool) (analyticItems []model.AnalyticItem, err error)
	FindByDates(positiveAmount bool, unit string, category string, merchant string) (expenses []model.DateAnalyticItem, err error)
	FindByDatesCumulative(positiveAmount bool, unit string, category string, merchant string) (items []model.DateAnalyticItem, err error)
}

type analyticRepository struct {
	log *slog.Logger
	db  *gorm.DB
}

func NewAnalyticRepository(log *slog.Logger, db *gorm.DB) AnalyticRepository {
	return &analyticRepository{
		log: log,
		db:  db,
	}
}

func (r *analyticRepository) Find(groupBy string, positiveAmount bool) (analyticItems []model.AnalyticItem, err error) {
	result := r.db.Select(groupBy + " as name, count(id), SUM(amount) as amount").
		Table("transactions").
		Where(r.whereAmount(positiveAmount)).
		Group("name").
		Order("amount" + r.orderDirection(positiveAmount)).
		Scan(&analyticItems)
	return analyticItems, result.Error
}

func (r *analyticRepository) FindByDates(positiveAmount bool, unit string, category string, merchant string) (items []model.DateAnalyticItem, err error) {
	result := r.db.Select("DATE_TRUNC('" + unit + "', datetime) AS date, count(amount) as count, SUM(amount) AS amount").
		Table("transactions").
		Where(r.whereAmount(positiveAmount)).
		Where(r.optionalWhere("category", category)).
		Where(r.optionalWhere("merchant", merchant)).
		Group("date").
		Order("date").
		Scan(&items)
	return items, result.Error
}

func (r *analyticRepository) FindByDatesCumulative(positiveAmount bool, unit string, category string, merchant string) (items []model.DateAnalyticItem, err error) {
	result := r.db.Select("DATE_TRUNC('" + unit + "', datetime) AS date, SUM(count(amount)) OVER (ORDER BY DATE_TRUNC('" + unit + "', datetime)) AS count, SUM(SUM(amount)) OVER (ORDER BY DATE_TRUNC('" + unit + "', datetime)) AS amount").
		Table("transactions").
		Where(r.whereAmount(positiveAmount)).
		Where(r.optionalWhere("category", category)).
		Where(r.optionalWhere("merchant", merchant)).
		Group("date").
		Order("date").
		Scan(&items)

	return items, result.Error
}

func (r *analyticRepository) whereAmount(positiveAmount bool) (where string) {
	if positiveAmount {
		return "amount > 0"
	} else {
		return "amount < 0"
	}
}

func (r *analyticRepository) optionalWhere(column, val string) (where string, value string) {
	if val != "" {
		return column + " = ?", val
	}
	return "1 = ?", "1"
}

func (r *analyticRepository) orderDirection(positiveAmount bool) (order string) {
	if positiveAmount {
		return " desc"
	} else {
		return " asc"
	}
}
