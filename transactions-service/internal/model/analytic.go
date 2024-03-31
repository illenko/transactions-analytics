package model

import "time"

type AnalyticItem struct {
	Name   string
	Count  int
	Amount float64
}

type DateAnalyticItem struct {
	Date   time.Time
	Count  int
	Amount float64
}
