package model

import "time"

type StatisticsBy struct {
	Name   string
	Count  int
	Amount float64
}

type DateAmount struct {
	Date   time.Time
	Amount float64
}
