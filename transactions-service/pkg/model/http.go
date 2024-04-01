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

type AnalyticResponse struct {
	Count  int             `json:"count"`
	Amount float64         `json:"amount"`
	Groups []AnalyticGroup `json:"groups"`
}

type AnalyticGroup struct {
	Name   string  `json:"name"`
	Count  int     `json:"count"`
	Amount float64 `json:"amount"`
}
