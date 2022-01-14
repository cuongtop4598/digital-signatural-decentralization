package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;"`
	Name        string
	PublicKey   string `gorm:"not null"`
	IDCard      string
	Phone       string
	Gmail       string
	DateOfBirth time.Time `json:"dateo_f_birth,omitempty"`
	CreateAt    time.Time `json:"creat_at,omitempty"`
	UpdateAt    time.Time `json:"update_at,omitempty"`
	DeleteAt    time.Time `json:"dalete_at,omitempty"`
}

type Document struct {
	DocID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Owner    uuid.UUID `gorm:"type:uuid;"`
	Name     string
	Type     string
	Path     string
	Public   bool
	CreateAt time.Time `json:"create_at,omitempty"`
	UpdateAt time.Time `json:"update_at,omitempty"`
	DeleteAt time.Time `json:"delete_at,omitempty"`
}

type UserAllow struct {
	UserID uuid.UUID `gorm:"type:uuid"`
	DocId  uuid.UUID `gorm:"type:uuid"`
}

type Transaction struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Timestamp time.Time `json:"time"`
	Type      string
}

type Role struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string
	Type string
}

type UserRole struct {
	UserID uuid.UUID `gorm:"type:uuid;"`
	RoleID uuid.UUID `gorm:"type:uuid;"`
}
