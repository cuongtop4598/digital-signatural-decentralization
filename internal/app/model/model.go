package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	PublicKey   string    `json:"public_key" gorm:"primaryKey;not null;unique"`
	CardID      string    `json:"card_id"`
	Phone       string    `json:"phone"`
	Gmail       string    `json:"gmail"`
	DateOfBirth string    `json:"dateo_of_birth,omitempty"`
	Password    string    `json:"password"`
	CreateAt    time.Time `json:"creat_at,omitempty"`
	UpdateAt    time.Time `json:"update_at,omitempty"`
	DeleteAt    time.Time `json:"dalete_at,omitempty"`
}

type Document struct {
	DocID     uuid.UUID `gorm:"primaryKey"`
	Number    int
	Owner     string
	Name      string
	Type      string
	Signature string
	Path      string
	Public    bool
	CreateAt  time.Time `json:"create_at,omitempty"`
	UpdateAt  time.Time `json:"update_at,omitempty"`
	DeleteAt  time.Time `json:"delete_at,omitempty"`
}

type UserAllow struct {
	UserID uuid.UUID
	DocID  uuid.UUID
}

type Transaction struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Timestamp time.Time `json:"time"`
	Type      string
}

type Role struct {
	ID   uuid.UUID `gorm:"primaryKey"`
	Name string
	Type string
}

type UserRole struct {
	UserID uuid.UUID
	RoleID uuid.UUID
}

func (u *User) SantisizePassword() {
	u.Password = ""
}
