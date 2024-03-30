package database

import (
	"github.com/illenko/transactions-service/internal/model"
	"gorm.io/gorm"
	"log/slog"
)

type AnalyticRepository interface {
	Find(groupBy string, positiveAmount bool) (analyticItems []model.AnalyticItem, err error)
	FindByDates(positiveAmount bool, unit string,
		period int, category *string, merchant *string) (expenses []model.DateAnalyticItem, err error)
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
		Table("analytics").
		Where(r.whereAmount(positiveAmount)).
		Group(groupBy).
		Order("amount" + r.orderDirection(positiveAmount)).
		Scan(&analyticItems)
	return analyticItems, result.Error
}

func (r *analyticRepository) FindByDates(positiveAmount bool, unit string, period int, category *string, merchant *string) (expenses []model.DateAnalyticItem, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *analyticRepository) Analytic(groupKey string, positiveAmount bool) (statistics []model.AnalyticItem, err error) {
	result := t.db.Select(groupKey + " as name, count(id), sum(amount) as amount").
		Table("analytics").
		Where(t.whereAmount(positiveAmount)).
		Group(groupKey).
		Order("amount" + t.orderDirection(positiveAmount)).
		Scan(&statistics)
	return statistics, result.Error
}

func (t *analyticRepository) MerchantExpenses(merchant string) (expenses []model.AnalyticItem, err error) {
	result := t.db.Select("DATE_TRUNC('month', datetime) AS date, SUM(amount) AS amount").
		Table("analytics").
		Where("merchant = ? and datetime > CURRENT_DATE - INTERVAL '6 months'", merchant).
		Group("date").
		Order("date").
		Scan(&expenses)
	return expenses, result.Error
}

func (t *analyticRepository) DateAmounts(positiveAmount bool) (dateAmounts []model.AnalyticItem, err error) {
	result := t.db.Select("DATE(datetime) as date, sum(sum(amount)) over (order by DATE(datetime)) as amount").
		Table("analytics").
		Where(t.whereAmount(positiveAmount)).
		Group("date").
		Order("date").
		Scan(&dateAmounts)

	return dateAmounts, result.Error
}

func (t *analyticRepository) whereAmount(positiveAmount bool) (where string) {
	if positiveAmount {
		return "amount > 0"
	} else {
		return "amount < 0"
	}
}

func (t *analyticRepository) orderDirection(positiveAmount bool) (order string) {
	if positiveAmount {
		return " desc"
	} else {
		return " asc"
	}
}
