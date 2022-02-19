package repository

import (
	"digitalsignature/internal/app/model"

	"gorm.io/gorm"
)

type DocumentRepository struct {
	DB *gorm.DB
}

func NewDocumentRepo(db *gorm.DB) *DocumentRepository {
	return &DocumentRepository{
		DB: db,
	}
}

func (repo *DocumentRepository) Create(doc *model.Document) error {
	result := repo.DB.Create(&doc)
	return result.Error
}
