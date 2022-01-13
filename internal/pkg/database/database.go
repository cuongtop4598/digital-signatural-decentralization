package database

import (
	"digitalsignature/config"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// PostgresConfig persists the config for our PostgreSQL database connection
type PgConfig struct {
	URL         string `env:"DATABASE_URL"`
	Host        string `env:"DATABASE_HOST" envDefault:"localhost"`
	Port        string `env:"DATABASE_PORT" envDefault:"5432"`
	User        string `env:"DATABASE_USER"`
	Password    string `env:"DATABASE_PASSWORD"`
	Database    string `env:"DATABASE_DB"`
	Environment string `env:"GIN_MODE"`
}

// GetConnection returns our pg database connection
// usage:
// db := config.GetConnection()
// defer db.Close()
func GetConnection(p *config.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		p.PostgresqlHost,
		p.PostgresqlUser,
		p.PostgresqlPassword,
		p.PostgresqlDbname,
		p.PostgresqlPort,
	)
	// define logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
