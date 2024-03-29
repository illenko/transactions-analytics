package model

import (
	"github.com/google/uuid"
	"time"
)

type TransactionResponse struct {
	ID       uuid.UUID `json:"id"`
	Datetime time.Time `json:"datetime"`
	Amount   float64   `json:"amount"`
	Category string    `json:"category"`
	Merchant string    `json:"merchant"`
}

type StatisticsResponse struct {
	Income   Statistics `json:"income"`
	Expenses Statistics `json:"expenses"`
}

type Statistics struct {
	Count       int               `json:"count"`
	Amount      float64           `json:"amount"`
	Groups      []StatisticsGroup `json:"groups"`
	DateAmounts []DateAmount      `json:"dateAmounts"`
}

type StatisticsGroup struct {
	Name   string  `json:"name"`
	Count  int     `json:"count"`
	Amount float64 `json:"amount"`
}

type MonthAmount struct {
	Month  string  `json:"month"`
	Amount float64 `json:"amount"`
}

type DateAmount struct {
	Date   string  `json:"date"`
	Amount float64 `json:"amount"`
}
