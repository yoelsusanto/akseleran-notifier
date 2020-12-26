package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/yoelsusanto/akseleran-notifier/config"
)

func CreateRedisClient(config *config.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       0,
	})

	return client
}
