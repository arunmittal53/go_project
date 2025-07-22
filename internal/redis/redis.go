package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	redisClient *redis.Client
}

func NewRedisClient() *RedisClient {
	redisOpt := &redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}

	redisClient := redis.NewClient(redisOpt)

	return &RedisClient{
		redisClient: redisClient,
	}
}

func (rc *RedisClient) SaveKey(ctx context.Context, key string, value string, expiration time.Duration) error {
	err := rc.redisClient.Set(ctx, key, value, expiration).Err()
	return err
}

func (rc *RedisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := rc.redisClient.Get(ctx, key).Result()
	return val, err
}

func (rc *RedisClient) Remove(ctx context.Context, key string) error {
	return rc.redisClient.Del(ctx, key).Err()
}
