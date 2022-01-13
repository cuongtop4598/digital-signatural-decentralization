package model

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	DocID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey`
	OwnerID  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Owner    User
	Name     string
	Type     string
	SavePath string
	Public   bool
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
}
