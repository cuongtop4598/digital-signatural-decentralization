package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresConfig persists the config for our PostgreSQL database connection
type DBConfig struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Database    string `yaml:"database"`
	Environment string `yaml:"env"`
}

func GetConnection(c *DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.Host, c.User, c.Password, c.Database, c.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
