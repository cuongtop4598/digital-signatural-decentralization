package migration

import (
	"digitalsignature/internal/app/model"

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
	//db.Migrator().CreateConstraint(&model.UserAllow{}, "User")
	return err
}
