package service

import (
	"context"

	"github.com/hosseinal/UrlShortner/internal/model"
	cachedurl "github.com/hosseinal/UrlShortner/internal/repository/cachedUrl"
	"github.com/hosseinal/UrlShortner/internal/repository/url"
)

type UrlService interface {
	CreateUrl(ctx context.Context, url *model.Url) error
	GetUrlByID(ctx context.Context, id string) (*model.Url, error)
	GetUrlByShortUrl(ctx context.Context, shortUrl string) (*model.Url, error)
}

type urlService struct {
	urlRepository       url.UrlRepository
	cachedUrlRepository cachedurl.CachedUrlRepository
}

func NewUrlService(urlRepository url.UrlRepository, cachedUrlRepository cachedurl.CachedUrlRepository) UrlService {
	return &urlService{urlRepository: urlRepository}
}

func (s *urlService) CreateUrl(ctx context.Context, url *model.Url) error {

	return s.urlRepository.CreateUrl(ctx, url)
}

func (s *urlService) GetUrlByID(ctx context.Context, id string) (*model.Url, error) {
	return s.urlRepository.GetUrlByID(ctx, id)
}

func (s *urlService) GetUrlByShortUrl(ctx context.Context, shortUrl string) (*model.Url, error) {

	cachedUrl, err := s.cachedUrlRepository.GetCachedUrlByShortUrl(ctx, shortUrl)
	if err == nil {
		return cachedUrl, nil
	}

	url, err := s.urlRepository.GetUrlByShortUrl(ctx, shortUrl)
	if err != nil {
		return nil, err
	}
	err = s.cachedUrlRepository.CreateCachedUrl(ctx, url)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (s *urlService) GetUrlByLongUrl(ctx context.Context, longUrl string) (*model.Url, error) {

	cachedUrl, err := s.cachedUrlRepository.GetCachedUrlByLongUrl(ctx, longUrl)
	if err == nil {
		return cachedUrl, nil
	}
	return s.urlRepository.GetUrlByLongUrl(ctx, longUrl)
}
