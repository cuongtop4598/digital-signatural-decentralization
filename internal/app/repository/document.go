package repository

import (
	"digitalsignature/internal/app/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DocumentRepo struct {
	DB  *gorm.DB
	log *zap.Logger
}

func NewDocumentRepo(db *gorm.DB, log *zap.Logger) *DocumentRepo {
	return &DocumentRepo{
		DB:  db,
		log: log,
	}
}

func (repo *DocumentRepo) Create(doc *model.Document) error {
	result := repo.DB.Create(&doc)
	return result.Error
}

func (repo *DocumentRepo) AllIsPublic() ([]model.Document, error) {
	docs := []model.Document{}
	result := repo.DB.Model(&model.Document{}).Where("public = ? ", true).Find(&docs)
	if result.Error != nil {
		return nil, result.Error
	}
	return docs, nil
}

func (repo *DocumentRepo) AllByOwner(publickey string) ([]model.Document, error) {
	docs := []model.Document{}
	result := repo.DB.Model(&model.Document{}).Where("owner = ? ", publickey).Find(&docs)
	if result.Error != nil {
		return nil, result.Error
	}
	return docs, nil
}

func (repo *DocumentRepo) IsPublic(name string) (bool, error) {
	var count int64
	var doc model.Document
	result := repo.DB.Model(&doc).Where("public = ?", true).Where("name = ?", name).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (repo *DocumentRepo) UpdateDocumentNumber(signature string, number int) error {

	var doc model.Document
	result := repo.DB.Model(&doc).Where("signature = ?", signature).Update("number", number)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
