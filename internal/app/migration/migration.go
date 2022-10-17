package migration

import (
	"digitalsignature/internal/app/model"
	"digitalsignature/internal/app/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, log *zap.Logger) error {
	tx := db.Migrator()
	if tx.HasTable("documents") {
		tx.DropTable("documents")
	}
	if tx.HasTable("multiple_documents") {
		tx.DropTable("multiple_documents")
	}
	if tx.HasTable("users") {
		tx.DropTable("users")
	}
	err := tx.AutoMigrate(
		&model.User{},
		&model.Document{},
		&model.MultipleDocument{},
		&model.UserAllow{},
		&model.Transaction{},
		&model.Role{},
		&model.UserRole{},
	)
	if err != nil {
		log.Sugar().Error(err)
		return err
	}
	// if !tx.HasConstraint(&model.Document{}, "fk_user_documents") {
	// 	db.Exec(
	// 		"alter table documents add constraint fk_user_documents foreign key (owner) references users(public_key)",
	// 	)
	// }
	// db.Exec(
	// 	"alter table user_allows add constraint fk_user_allow_users foreign key (user_id) references users(id)",
	// )
	// db.Exec(
	// 	"alter table user_allows add constraint fk_user_allow_documents foreign key (doc_id) references documents(doc_id)",
	// )
	// db.Exec(
	// 	"alter table user_roles add constraint fk_user_roles_roles foreign key (role_id) references roles(id)",
	// )
	// db.Exec(
	// 	"alter table user_roles add constraint fk_user_roles_users foreign key (user_id) references users(id)",
	// )
	err = InsertData(db)
	if err != nil {
		log.Sugar().Error(err)
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
	return nil
}
