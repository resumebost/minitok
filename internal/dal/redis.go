package dal

import (
	"github.com/redis/go-redis/v9"
	"minitok/internal/conf"
)

var Redis *redis.Client

// InitRedis TODO: 更多配置
func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddress,
		Password: conf.RedisPassword,
		DB:       0,
		PoolSize: 100,
	})
}
