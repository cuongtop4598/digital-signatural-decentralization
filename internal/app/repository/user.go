package repository

import (
	"digitalsignature/internal/app/model"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB  *gorm.DB
	log *zap.Logger
}

func NewUserRepository(db *gorm.DB, log *zap.Logger) *UserRepo {
	return &UserRepo{
		DB:  db,
		log: log,
	}
}

func (repo *UserRepo) Create(user model.User) error {
	result := repo.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	repo.log.Info("Create user:", zap.String("Info", strings.Join([]string{user.Phone, user.PublicKey}, "-")))
	return nil
}

func (repo *UserRepo) IsExist(phone string, publicKey string) bool {
	var exist bool
	err := repo.DB.Model(&model.User{}).Select("count(*) > 0").Where("phone = ? or public_key = ?", phone, publicKey).Find(&exist).Error
	if err != nil {
		repo.log.Sugar().Error(err)
		return false
	}
	return exist
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

func (repo *UserRepo) GetPublickeyByPhone(phone string) (string, error) {
	user := model.User{}
	result := repo.DB.Where("phone = ?", phone).First(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return user.PublicKey, nil
}

func (repo *UserRepo) GetUserByGmail(gmail string) (*model.User, error) {
	user := model.User{}
	result := repo.DB.Where("gmail = ?", gmail).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepo) CheckLogin(password string, phone string) (bool, *model.User, error) {
	user := model.User{}
	result := repo.DB.Model(&model.User{}).Where("password = ? and phone = ?", password, phone).First(&user)
	if result.Error != nil {
		repo.log.Sugar().Error(result.Error)
		return false, nil, result.Error
	}
	if user.PublicKey != "" {
		return true, &user, nil
	}
	return false, nil, nil
}

func (repo *UserRepo) GetPhoneByPublickey(publickey string) (string, error) {
	user := model.User{}
	result := repo.DB.Model(&model.User{}).Where("public_key = ? ", publickey).First(&user)
	if result.Error != nil {
		return "", result.Error
	} else {
		return user.Phone, nil
	}
}
