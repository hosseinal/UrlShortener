package rediscache

import (
	"context"
	"strconv"

	"github.com/hosseinal/UrlShortner/internal/config"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	Client *redis.Client
}

func NewRedisCache(config config.RedisCacheConfig) (*RedisCache, error) {

	port := strconv.Itoa(config.Port)
	addr := config.Host + ":" + port
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       0,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return &RedisCache{Client: client}, nil
}
