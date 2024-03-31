package database

import (
	"github.com/illenko/transactions-service/internal/model"
	"gorm.io/gorm"
	"log/slog"
	"strconv"
)

type AnalyticRepository interface {
	Find(groupBy string, positiveAmount bool) (analyticItems []model.AnalyticItem, err error)
	FindByDates(positiveAmount bool, unit string, period int, category string, merchant string) (expenses []model.DateAnalyticItem, err error)
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
	result := r.db.Select(groupBy + " as name, count(id), sum(amount) as amount").
		Table("transactions").
		Where(r.whereAmount(positiveAmount)).
		Group("name").
		Order("amount" + r.orderDirection(positiveAmount)).
		Scan(&analyticItems)
	return analyticItems, result.Error
}

func (r *analyticRepository) FindByDates(positiveAmount bool, unit string, period int, category string, merchant string) (items []model.DateAnalyticItem, err error) {
	result := r.db.Select("DATE_TRUNC('"+unit+"', datetime) AS name, count(*) as count, SUM(amount) AS amount").
		Table("transactions").
		Where("datetime > CURRENT_DATE - INTERVAL '"+strconv.Itoa(period)+" "+unit+"s'", merchant).
		Where(r.whereAmount(positiveAmount)).
		Group("name").
		Order("name").
		Scan(&items)
	return items, result.Error
}

func (r *analyticRepository) Analytic(groupKey string, positiveAmount bool) (statistics []model.AnalyticItem, err error) {
	result := r.db.Select(groupKey + " as name, count(id), sum(amount) as amount").
		Table("analytics").
		Where(r.whereAmount(positiveAmount)).
		Group(groupKey).
		Order("amount" + r.orderDirection(positiveAmount)).
		Scan(&statistics)
	return statistics, result.Error
}

func (r *analyticRepository) DateAmounts(positiveAmount bool) (dateAmounts []model.AnalyticItem, err error) {
	result := r.db.Select("DATE(datetime) as date, sum(sum(amount)) over (order by DATE(datetime)) as amount").
		Table("analytics").
		Where(r.whereAmount(positiveAmount)).
		Group("date").
		Order("date").
		Scan(&dateAmounts)

	return dateAmounts, result.Error
}

func (r *analyticRepository) whereAmount(positiveAmount bool) (where string) {
	if positiveAmount {
		return "amount > 0"
	} else {
		return "amount < 0"
	}
}

func (r *analyticRepository) orderDirection(positiveAmount bool) (order string) {
	if positiveAmount {
		return " desc"
	} else {
		return " asc"
	}
}
