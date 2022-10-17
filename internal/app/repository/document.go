package repository

import (
	"digitalsignature/internal/app/model"

	"github.com/google/uuid"
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
	doc.ID = uuid.New()
	result := repo.DB.Create(&doc)
	return result.Error
}

func (repo *DocumentRepo) CreateMultiSignerDocument(doc *model.MultipleDocument) error {
	doc.ID = uuid.New()
	result := repo.DB.Create(&doc)
	return result.Error
}

func (repo *DocumentRepo) UpdateMultiSignerState(doc *model.MultipleDocument) error {
	result := repo.DB.Model(&doc).Where("id = ?", doc.ID).Updates(doc)
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

func (repo *DocumentRepo) AllMultiSignerByPartner(partner string) ([]model.MultipleDocument, error) {
	docs := []model.MultipleDocument{}
	result := repo.DB.Model(&model.MultipleDocument{}).Where("partner_a = ? or partner_b = ? ", partner, partner).Find(&docs)
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

func (repo *DocumentRepo) IsExisted(digest string) bool {
	var doc model.MultipleDocument
	result := repo.DB.Model(&doc).Where("digest = ?", digest).First(&doc)
	if result.Error != nil {
		repo.log.Sugar().Error(result.Error)
		return false
	}
	if doc.Digest != "" {
		return true
	} else {
		return false
	}
}

func (repo *DocumentRepo) GetMultiByDigest(digest string) *model.MultipleDocument {
	doc := model.MultipleDocument{}
	result := repo.DB.Model(&doc).Where("digest = ?", digest).First(&doc)
	if result.Error != nil {
		return nil
	} else {
		return &doc
	}
}
