package routers

import (
	"crypto/md5"
	"encoding/hex"

	"digitalsignature/internal/app/migration"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewCMD(db *gorm.DB, log *zap.Logger) *CMD {
	return &CMD{
		db:  db,
		log: log,
	}
}

type CMD struct {
	db  *gorm.DB
	log *zap.Logger
}

func (cmd *CMD) MigrateDB() {
	migration.Migrate(cmd.db, cmd.log)
}

func (cmd *CMD) GenerateMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
