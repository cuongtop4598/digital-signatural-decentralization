package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey"`
	Timestamp time.Time
	Type      string
}
