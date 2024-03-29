package model

import "time"

type StatisticsBy struct {
	Name   string
	Count  int
	Amount float64
}

type MerchantExpense struct {
	Month  time.Time
	Amount float64
}
