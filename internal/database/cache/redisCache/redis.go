package rediscache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	Client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{Client: client}
}

func (c *RedisCache) Get(key string) (string, error) {
	return c.Client.Get(context.Background(), key).Result()
}

func (c *RedisCache) Set(key string, value string) error {
	return c.Client.Set(context.Background(), key, value, 0).Err()
}

func (c *RedisCache) Delete(key string) error {
	return c.Client.Del(context.Background(), key).Err()
}
