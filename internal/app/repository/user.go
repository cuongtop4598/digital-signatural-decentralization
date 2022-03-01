package repository

import (
	"digitalsignature/internal/app/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (repo *UserRepo) Create(user model.User) error {
	result := repo.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *UserRepo) GetUserByPubkey(pubkey string) (*model.User, error) {
	user := model.User{}
	result := repo.DB.Where("public_key = ?", pubkey).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepo) GetUserById(id string) (*model.User, error) {
	user := model.User{}
	result := repo.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepo) GetUserByPhone(phone string) (*model.User, error) {
	user := model.User{}
	result := repo.DB.Where("phone = ?", phone).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepo) GetUserByGmail(gmail string) (*model.User, error) {
	user := model.User{}
	result := repo.DB.Where("gmail = ?", gmail).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepo) CheckLogin(password string, phone string) (bool, error) {
	user := model.User{}
	result := repo.DB.Model(&model.User{}).Where("password = ?", password).Where("phone = ?", phone).First(&user)
	if result.Error != nil {
		return false, result.Error
	}
	if user.PublicKey != "" {
		return true, nil
	}
	return false, nil
}

func (repo *UserRepo) GetPhoneByPublickey(publickey string) (string, error) {
	user := model.User{}
	result := repo.DB.Model(&model.User{}).Where("publickey = ? ", publickey).First(&user)
	if result.Error != nil {
		return "", result.Error
	} else {
		return user.Phone, nil
	}
}
