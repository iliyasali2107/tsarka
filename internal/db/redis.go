package db

import (
	"tsarka/internal/config"

	"github.com/redis/go-redis/v9"
)

func InitRedisDB(cfg config.Config) *redis.Client {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisUrl,
		Password: "",
		DB:       0,
	})

	return redisDB
}
