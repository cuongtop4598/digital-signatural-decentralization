package database

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/go-pg/pg/v10"
	"go.uber.org/zap"
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
func GetConnection(log *zap.Logger) *pg.DB {
	c := GetPgConfig()
	fmt.Println(c.Host)
	fmt.Println(c.User)
	// if DATABASE_URL is valid, we will use its constituent values in preference
	validConfig, err := validPostgresURL(c.URL)
	if err == nil {
		c = validConfig
	}
	db := pg.Connect(&pg.Options{
		Addr:     c.Host + ":" + c.Port,
		User:     c.User,
		Password: c.Password,
		Database: c.Database,
		PoolSize: 150,
	})
	db.AddQueryHook(dbLogger{
		log: log,
	})
	return db
}

type dbLogger struct {
	log *zap.Logger
}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	fq, _ := q.FormattedQuery()
	d.log.Sugar().Info(string(fq))
	return nil
}

// GetPostgresConfig returns a PostgresConfig pointer with the correct Postgres Config values
func GetPgConfig() *PgConfig {
	c := PgConfig{}
	if err := env.Parse(&c); err != nil {
		fmt.Printf("%+v\n", err)
		fmt.Println("Lỗi")
	}
	return &c
}

func validPostgresURL(URL string) (*PgConfig, error) {
	if URL == "" || strings.TrimSpace(URL) == "" {
		return nil, errors.New("database url is blank")
	}

	validURL, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}
	c := &PgConfig{}
	c.URL = URL
	c.Host = validURL.Host
	c.Database = validURL.Path
	c.Port = validURL.Port()
	c.User = validURL.User.Username()
	c.Password, _ = validURL.User.Password()
	return c, nil
}
