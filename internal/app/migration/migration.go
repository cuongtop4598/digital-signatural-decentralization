package migration

import (
	"digitalsignature/internal/app/model"
	"digitalsignature/internal/app/utils"
	"fmt"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.Migrator().AutoMigrate(
		&model.User{},
		&model.Document{},
		&model.UserAllow{},
		&model.Transaction{},
		&model.Role{},
		&model.UserRole{},
	)
	db.Exec(
		"alter table documents add constraint fk_user_documents foreign key (owner) references users(id)",
	)
	db.Exec(
		"alter table user_allows add constraint fk_user_allow_users foreign key (user_id) references users(id)",
	)
	db.Exec(
		"alter table user_allows add constraint fk_user_allow_documents foreign key (doc_id) references documents(doc_id)",
	)
	db.Exec(
		"alter table user_roles add constraint fk_user_roles_roles foreign key (role_id) references roles(id)",
	)
	db.Exec(
		"alter table user_roles add constraint fk_user_roles_users foreign key (user_id) references users(id)",
	)
	err = InsertData(db)
	if err != nil {
		fmt.Printf("Error %s", err.Error())
		return err
	}
	return nil
}

func InsertData(db *gorm.DB) error {
	user := []model.User{}
	if err := utils.Json2struct("./internal/app/mockup/user.json", &user); err != nil {
		return err
	}
	db.Create(&user)
	document := []model.Document{}
	if err := utils.Json2struct("./internal/app/mockup/document.json", &document); err != nil {
		return err
	}
	db.Create(&document)
	userallow := []model.UserAllow{}
	if err := utils.Json2struct("./internal/app/mockup/userallow.json", &userallow); err != nil {
		return err
	}
	db.Create(&userallow)
	role := []model.Role{}
	if err := utils.Json2struct("./internal/app/mockup/role.json", &role); err != nil {
		return err
	}
	db.Create(&role)
	userrole := []model.UserRole{}
	if err := utils.Json2struct("./internal/app/mockup/userrole.json", &userrole); err != nil {
		return err
	}
	db.Create(&userrole)
	return nil
}
