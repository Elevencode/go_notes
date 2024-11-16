package database

import (
	"fmt"
	"go_notes/envs"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func InitRedis() error {
	redisUrl := fmt.Sprintf("%s:%s", envs.ServerEnvs.REDIS_HOST, envs.ServerEnvs.REDIS_PORT)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "",
		DB:       0,
	})

	status := RedisClient.Ping()
	if status.Val() == "PONG" {
		return nil
	} else {
		return fmt.Errorf("ошибка при подключении к Redis: %v", status)
	}
}
