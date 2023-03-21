package redis_cache

import (
	"bytes"
	"context"
	config "crow-blog-backend/src/config"
	"encoding/gob"
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

func GetDecode(key string, v interface{}) error {
	client := config.GetRedisClient()
	ctx := context.Background()
	result, err := client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(bytes.NewReader(result))
	err = dec.Decode(v)
	return err
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
