package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID              uuid.UUID `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name" gorm:"not null"`
	PublicKey       string    `json:"public_key" gorm:"primaryKey;not null;unique"`
	CardID          string    `json:"card_id" gorm:"not null"`
	Phone           string    `json:"phone" gorm:"primaryKey;not null;unique"`
	Gmail           string    `json:"gmail" gorm:"not null"` // verify Gmail
	DateOfBirth     string    `json:"date_of_birth" gorm:"not null"`
	Password        string    `json:"password" gorm:"not null"` // hash MD5
	BlockNumber     int       `json:"block_number,omitempty"`
	BlockTimetamp   int       `json:"block_timestamp,omitempty"`
	BlockHash       string    `json:"block_hash,omitempty"`
	TransactionHash string    `json:"transaction_hash"`
	CreateAt        time.Time `json:"creat_at,omitempty"`
	UpdateAt        time.Time `json:"update_at,omitempty"`
	DeleteAt        time.Time `json:"dalete_at,omitempty"`
}

type Document struct {
	ID              uuid.UUID `gorm:"primaryKey, autoIncrement"`
	IndexOnchain    int       `json:"index_onchain"`
	Owner           string    `gorm:"type:text"`
	Name            string    `json:"name"`
	BlockNumber     string    `json:"block_number,omitempty"`
	BlockHash       string    `json:"block_hash,omitempty"`
	TransactionHash string    `json:"transaction_hash"`
	Signature       string    `json:"signature"`
	Public          bool      `json:"public"`
	TypeFile        string    `json:"type_file"`
	CreateAt        time.Time `json:"create_at,omitempty"`
	UpdateAt        time.Time `json:"update_at,omitempty"`
	DeleteAt        time.Time `json:"delete_at,omitempty"`
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
