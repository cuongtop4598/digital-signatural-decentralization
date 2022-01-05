package database

import (
	"context"
	"digitalsignature/config"
	"fmt"

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
func GetConnection(log *zap.Logger, p *config.PostgresConfig) (*pg.DB, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.PostgresqlUser,
		p.PostgresqlPassword,
		p.PostgresqlHost,
		p.PostgresqlPort,
		p.PostgresqlDbname,
	)
	opt, err := pg.ParseURL(dbURL)
	if err != nil {
		return nil, err
	}
	db := pg.Connect(opt)

	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}
	db.AddQueryHook(dbLogger{
		log: log,
	})

	return db, nil
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
