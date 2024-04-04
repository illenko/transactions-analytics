package analytic

import "time"

type Item struct {
	Name   string
	Count  int
	Amount float64
}

type DateItem struct {
	Date   time.Time
	Count  int
	Amount float64
}
