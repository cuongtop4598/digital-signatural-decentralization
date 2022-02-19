package repository

import (
	"digitalsignature/internal/app/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DocumentRepository struct {
	DB  *gorm.DB
	log *zap.Logger
}

func NewDocumentRepo(db *gorm.DB, log *zap.Logger) *DocumentRepository {
	return &DocumentRepository{
		DB:  db,
		log: log,
	}
}

func (repo *DocumentRepository) Create(doc *model.Document) error {
	result := repo.DB.Create(&doc)
	return result.Error
}

func (repo *DocumentRepository) IsPublic(name string) bool {
	var count int64
	var doc model.Document
	result := repo.DB.Model(&doc).Where("public = ?", true).Where("name = ?", doc.Name).Count(&count)
	if result.Error != nil {
		repo.log.Sugar().Error(result.Error)
	}
	if count > 0 {
		return true
	} else {
		return false
	}
}
