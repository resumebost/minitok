package dal

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"minitok/internal/constant"
	"time"
)

var redisConstants = &constant.AllConstants.Redis

func InitRedis() *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr:     redisConstants.Addr(),
		Password: redisConstants.Password,
		DB:       0,
		PoolSize: 100,
	})

	sec := time.Duration(5)
	ctx, cancel := context.WithTimeout(context.Background(), sec*time.Second)
	defer cancel()

	_, err := c.Ping(ctx).Result()
	if err != nil {
		klog.Fatalf("Redis ping is unreachable in %v seconds: %v", sec, err)
	}

	return c
}
