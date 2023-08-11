package dal

import (
	"github.com/redis/go-redis/v9"
	"minitok/internal/conf"
)

// InitRedis TODO: 心跳和更多配置
func InitRedis() *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddress,
		Password: conf.RedisPassword,
		DB:       0,
		PoolSize: 100,
	})
	return c
}
