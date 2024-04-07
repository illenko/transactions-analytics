package database

import (
	"github.com/illenko/analytics-service/internal/model"
	"gorm.io/gorm"
	"log/slog"
)

type AnalyticsRepository interface {
	Find(group string, positiveAmount bool) (analyticsItems []model.AnalyticsItem, err error)
	FindByDates(positiveAmount bool, unit string) (expenses []model.DateAnalyticsItem, err error)
	FindByDatesCumulative(positiveAmount bool, unit string) (items []model.DateAnalyticsItem, err error)
}

type analyticsRepository struct {
	log *slog.Logger
	db  *gorm.DB
}

func NewAnalyticsRepository(log *slog.Logger, db *gorm.DB) AnalyticsRepository {
	return &analyticsRepository{
		log: log,
		db:  db,
	}
}

func (r *analyticsRepository) Find(group string, positiveAmount bool) (analyticsItems []model.AnalyticsItem, err error) {
	result := r.db.Select(group + " as name, count(id), SUM(amount) as amount").
		Table("transactions").
		Where(r.whereAmount(positiveAmount)).
		Group("name").
		Order("amount" + r.orderDirection(positiveAmount)).
		Scan(&analyticsItems)
	return analyticsItems, result.Error
}

func (r *analyticsRepository) FindByDates(positiveAmount bool, unit string) (items []model.DateAnalyticsItem, err error) {
	result := r.db.Select("DATE_TRUNC('" + unit + "', datetime) AS date, count(amount) as count, SUM(amount) AS amount").
		Table("transactions").
		Where(r.whereAmount(positiveAmount)).
		Group("date").
		Order("date").
		Scan(&items)
	return items, result.Error
}

func (r *analyticsRepository) FindByDatesCumulative(positiveAmount bool, unit string) (items []model.DateAnalyticsItem, err error) {
	result := r.db.Select("DATE_TRUNC('" + unit + "', datetime) AS date, SUM(count(amount)) OVER (ORDER BY DATE_TRUNC('" + unit + "', datetime)) AS count, SUM(SUM(amount)) OVER (ORDER BY DATE_TRUNC('" + unit + "', datetime)) AS amount").
		Table("transactions").
		Where(r.whereAmount(positiveAmount)).
		Group("date").
		Order("date").
		Scan(&items)

	return items, result.Error
}

func (r *analyticsRepository) whereAmount(positiveAmount bool) (where string) {
	if positiveAmount {
		return "amount > 0"
	} else {
		return "amount < 0"
	}
}

func (r *analyticsRepository) orderDirection(positiveAmount bool) (order string) {
	if positiveAmount {
		return " desc"
	} else {
		return " asc"
	}
}
