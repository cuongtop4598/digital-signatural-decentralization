package database

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewDBConnection(log *zap.Logger, sc *string) *gorm.DB {
	config := getPgConfig()
	connection := config.ToConnectionString()
	fmt.Println(connection)
	public := "public"
	if sc == nil {
		sc = &public
	}
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   *sc + ".",
			SingularTable: false,
		},
	})
	if err != nil {
		log.Error(err.Error())
	}
	return db
}

type PgConfig struct {
	URL         string `env:"DATABASE_URL"`
	Host        string `env:"DATABASE_HOST"`
	Port        int32  `env:"DATABASE_PORT"`
	User        string `env:"DATABASE_USER"`
	SSLMode     string `env:"SSL_MODE"`
	Password    string `env:"DATABASE_PASSWORD"`
	Database    string `env:"DATABASE_DB"`
	Environment string `env:"GIN_MODE"`
}

func (c PgConfig) ToConnectionString() string {
	return fmt.Sprintf("host=%s user=%s dbname=%s port=%d password=%s",
		c.Host,
		c.User,
		c.Database,
		c.Port,
		c.Password,
	)
}

func getPgConfig() *PgConfig {
	c := PgConfig{}
	if err := env.Parse(&c); err != nil {
		log.Fatal(err)
	}
	return &c
}
