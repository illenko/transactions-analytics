package analytic

import (
	"gorm.io/gorm"
	"log/slog"
)

type Repository interface {
	Find(group string, positiveAmount bool) (analyticItems []Item, err error)
	FindByDates(positiveAmount bool, unit string) (expenses []DateItem, err error)
	FindByDatesCumulative(positiveAmount bool, unit string) (items []DateItem, err error)
}

type repository struct {
	log *slog.Logger
	db  *gorm.DB
}

func NewRepository(log *slog.Logger, db *gorm.DB) Repository {
	return &repository{
		log: log,
		db:  db,
	}
}

func (r *repository) Find(group string, positiveAmount bool) (analyticItems []Item, err error) {
	result := r.db.Select(group + " as name, count(id), SUM(amount) as amount").
		Table("transactions").
		Where(r.whereAmount(positiveAmount)).
		Group("name").
		Order("amount" + r.orderDirection(positiveAmount)).
		Scan(&analyticItems)
	return analyticItems, result.Error
}

func (r *repository) FindByDates(positiveAmount bool, unit string) (items []DateItem, err error) {
	result := r.db.Select("DATE_TRUNC('" + unit + "', datetime) AS date, count(amount) as count, SUM(amount) AS amount").
		Table("transactions").
		Where(r.whereAmount(positiveAmount)).
		Group("date").
		Order("date").
		Scan(&items)
	return items, result.Error
}

func (r *repository) FindByDatesCumulative(positiveAmount bool, unit string) (items []DateItem, err error) {
	result := r.db.Select("DATE_TRUNC('" + unit + "', datetime) AS date, SUM(count(amount)) OVER (ORDER BY DATE_TRUNC('" + unit + "', datetime)) AS count, SUM(SUM(amount)) OVER (ORDER BY DATE_TRUNC('" + unit + "', datetime)) AS amount").
		Table("transactions").
		Where(r.whereAmount(positiveAmount)).
		Group("date").
		Order("date").
		Scan(&items)

	return items, result.Error
}

func (r *repository) whereAmount(positiveAmount bool) (where string) {
	if positiveAmount {
		return "amount > 0"
	} else {
		return "amount < 0"
	}
}

func (r *repository) orderDirection(positiveAmount bool) (order string) {
	if positiveAmount {
		return " desc"
	} else {
		return " asc"
	}
}
