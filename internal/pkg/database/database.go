package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
func GetConnection() (*gorm.DB, error) {
	dsn := "host=172.18.0.3 user=admin password=mypassword dbname=dsd port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
