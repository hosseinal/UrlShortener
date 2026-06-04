package cachedurl

import (
	"context"

	"github.com/hosseinal/UrlShortner/internal/model"

	"github.com/redis/go-redis/v9"
)

type CachedUrlRepository interface {
	CreateCachedUrl(ctx context.Context, cachedUrl *model.Url) error
	GetCachedUrlByID(ctx context.Context, id string) (*model.Url, error)
	GetCachedUrlByShortUrl(ctx context.Context, shortUrl string) (*model.Url, error)
	GetCachedUrlByLongUrl(ctx context.Context, longUrl string) (*model.Url, error)
}

type cachedUrlRepository struct {
	redis *redis.Client
}

func NewCachedUrlRepository(redis *redis.Client) CachedUrlRepository {
	return &cachedUrlRepository{redis: redis}
}

func (r *cachedUrlRepository) CreateCachedUrl(ctx context.Context, cachedUrl *model.Url) error {
	return r.redis.Set(ctx, cachedUrl.ShortUrl, cachedUrl, 0).Err()
}

func (r *cachedUrlRepository) GetCachedUrlByID(ctx context.Context, id string) (*model.Url, error) {
	var url model.Url
	if err := r.redis.Get(ctx, id).Scan(&url); err != nil {
		return nil, err
	}
	return &url, nil
}

func (r *cachedUrlRepository) GetCachedUrlByShortUrl(ctx context.Context, shortUrl string) (*model.Url, error) {
	var url model.Url
	if err := r.redis.Get(ctx, shortUrl).Scan(&url); err != nil {
		return nil, err
	}
	return &url, nil
}

func (r *cachedUrlRepository) GetCachedUrlByLongUrl(ctx context.Context, longUrl string) (*model.Url, error) {
	var url model.Url
	if err := r.redis.Get(ctx, longUrl).Scan(&url); err != nil {
		return nil, err
	}
	return &url, nil
}
