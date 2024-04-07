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

type AnalyticsResponse struct {
	Count  int              `json:"count"`
	Amount float64          `json:"amount"`
	Groups []AnalyticsGroup `json:"groups"`
}

type AnalyticsGroup struct {
	Name   string  `json:"name"`
	Count  int     `json:"count"`
	Amount float64 `json:"amount"`
}
