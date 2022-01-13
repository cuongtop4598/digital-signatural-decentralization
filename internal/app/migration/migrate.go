package migration

import (
	"digitalsignature/internal/app/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.Migrator().CreateTable(model.User{})
	// db.Migrator().CreateTable(&model.UserAllow)
	// db.Migrator().CreateConstraint(&model.UserAllow{}, "User")
	// db.Migrator().CreateConstraint(&model.User{}, "fk_allow_users")

	// db.Migrator().CreateConstraint(&model.Document{}, "User")
	// db.Migrator().CreateConstraint(&model.Document{}, "fk_document_users")

	// db.Migrator().CreateConstraint(&model.UserAllow{}, "Document")
	// db.Migrator().CreateConstraint(&model.UserAllow{}, "fk_allow_documents")

	// db.Migrator().CreateConstraint(&model.UserRole{}, "User")
	// db.Migrator().CreateConstraint(&model.UserRole{}, "fk_role_users")

	// db.Migrator().CreateConstraint(&model.UserRole{}, "Role")
	// db.Migrator().CreateConstraint(&model.UserRole{}, "fk_userrole_roles")

	return nil
}
