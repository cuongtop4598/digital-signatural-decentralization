package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	PublicKey   string    `json:"public_key" gorm:"primaryKey;not null;unique"`
	CardID      string    `json:"card_id" gorm:"not null"`
	Phone       string    `json:"phone" gorm:"primaryKey;not null;unique"`
	Gmail       string    `json:"gmail" gorm:"not null"` // verify Gmail
	DateOfBirth string    `json:"date_of_birth" gorm:"not null"`
	Password    string    `json:"password" gorm:"not null"` // hash MD5
	CreateAt    time.Time `json:"creat_at,omitempty"`
	UpdateAt    time.Time `json:"update_at,omitempty"`
	DeleteAt    time.Time `json:"dalete_at,omitempty"`
}

type Document struct {
	ID            uuid.UUID `gorm:"primaryKey, autoIncrement"`
	IndexOnchain  int
	Owner         []string `gorm:"type:text"`
	Name          string
	BlockNumber   int
	BlockTimetamp int
	Signature     string
	Public        bool
	TypeFile      string
	CreateAt      time.Time `json:"create_at,omitempty"`
	UpdateAt      time.Time `json:"update_at,omitempty"`
	DeleteAt      time.Time `json:"delete_at,omitempty"`
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

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreateAt = time.Now()
	return
}

func (doc *Document) BeforeCreate(tx *gorm.DB) (err error) {
	doc.ID = uuid.New()
	doc.CreateAt = time.Now()
	return
}
