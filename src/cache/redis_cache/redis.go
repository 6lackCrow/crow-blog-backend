package redis_cache

import (
	"context"
	config "crow-blog-backend/src/config"
	"time"
)

func Get(key string) (string, error) {
	client := config.GetRedisClient()
	ctx := context.Background()
	return client.Get(ctx, key).Result()
}

func GetScan(key string, v interface{}) error {
	client := config.GetRedisClient()
	ctx := context.Background()
	return client.Get(ctx, key).Scan(v)
}

func GetSet(key string, value interface{}) (string, error) {
	client := config.GetRedisClient()
	ctx := context.Background()
	return client.GetSet(ctx, key, value).Result()
}

func SetNX(key string, value interface{}, expireTime time.Duration) (bool, error) {
	client := config.GetRedisClient()
	ctx := context.Background()
	return client.SetNX(ctx, key, value, expireTime).Result()
}

func Set(key string, value interface{}, expireTime time.Duration) error {
	client := config.GetRedisClient()
	ctx := context.Background()
	return client.Set(ctx, key, value, expireTime).Err()
}

func Remove(key string) error {
	client := config.GetRedisClient()
	ctx := context.Background()
	return client.Del(ctx, key).Err()
}
