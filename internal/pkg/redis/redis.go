package redis

import (
	"context"
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/go-redis/redis/v8"
)

type Config struct {
	Host     string `env:"REDIS_HOST" envDefault:"localhost"`
	Port     string `env:"REDIS_PORT" envDefault:"6379"`
	Password string `env:"REDIS_PASSWORD"`
	Database int    `env:"REDIS_DB"`
}

func GetConnection() *redis.Client {
	c := GetRedisConfig()
	fmt.Printf("Connect redis host: %s:%s", c.Host, c.Port)

	redis := redis.NewClient(&redis.Options{
		Addr:     c.Host + ":" + c.Port,
		Password: c.Password,
		DB:       c.Database,
	})
	if err := redis.Ping(context.TODO()).Err(); err != nil {
		fmt.Println(err)
	}
	return redis
}

func GetRedisConfig() *Config {
	c := Config{}
	if err := env.Parse(&c); err != nil {
		fmt.Printf("%+v\n", err)
		fmt.Println("redis env error!")
	}
	return &c
}
