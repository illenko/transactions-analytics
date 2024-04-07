package model

import "time"

type AnalyticsItem struct {
	Name   string
	Count  int
	Amount float64
}

type DateAnalyticsItem struct {
	Date   time.Time
	Count  int
	Amount float64
}
