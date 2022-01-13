package model

import (
	"time"

	"github.com/google/uuid"
)

type Test struct {
	Test1 int
	Test2 string
}

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;uuid.generate()"`
	Name        string
	PublicKey   string
	IDCard      string
	Phone       string
	Gmail       string
	DateOfBirth time.Time
	CreateAt    time.Time
	UpdateAt    time.Time
	DeleteAt    time.Time
}

type UserAllow struct {
	User     []User
	Document []Document
}

type Role struct {
	ID   uuid.UUID `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey"`
	Name string
	Type string
}

type UserRole struct {
	User []User
	Role []Role
}
