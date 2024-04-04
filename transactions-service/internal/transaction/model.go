package transaction

import (
	"github.com/google/uuid"
	"time"
)

type Entity struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Datetime time.Time
	Amount   float64
	Category string
	Merchant string
}
