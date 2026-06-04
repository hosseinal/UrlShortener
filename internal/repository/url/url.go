package url

import (
	"context"

	"github.com/hosseinal/UrlShortner/internal/model"

	"gorm.io/gorm"
)

type UrlRepository interface {
	CreateUrl(ctx context.Context, url *model.Url) error
	GetUrlByID(ctx context.Context, id string) (*model.Url, error)
	GetUrlByShortUrl(ctx context.Context, shortUrl string) (*model.Url, error)
	GetUrlByLongUrl(ctx context.Context, longUrl string) (*model.Url, error)
	DeleteUrl(ctx context.Context, id string) error
}

type urlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) UrlRepository {
	return &urlRepository{db: db}
}

func (r *urlRepository) CreateUrl(ctx context.Context, url *model.Url) error {
	return r.db.WithContext(ctx).Create(url).Error
}

func (r *urlRepository) GetUrlByID(ctx context.Context, id string) (*model.Url, error) {
	var url model.Url
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}

func (r *urlRepository) GetUrlByShortUrl(ctx context.Context, shortUrl string) (*model.Url, error) {
	var url model.Url
	if err := r.db.WithContext(ctx).Where("short_url = ?", shortUrl).First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}

func (r *urlRepository) GetUrlByLongUrl(ctx context.Context, longUrl string) (*model.Url, error) {
	var url model.Url
	if err := r.db.WithContext(ctx).Where("long_url = ?", longUrl).First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}

func (r *urlRepository) DeleteUrl(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.Url{}, id).Error
}
