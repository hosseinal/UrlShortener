package dto

import "time"

type CreateUrlRequest struct {
	LongUrl string `json:"long_url"`
}

type CreateUrlResponse struct {
	ID        string    `json:"id"`
	ShortUrl  string    `json:"short_url"`
	LongUrl   string    `json:"long_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

type GetUrlRequest struct {
	ShortUrl string `json:"short_url"`
}

type GetUrlResponse struct {
	ID        string    `json:"id"`
	ShortUrl  string    `json:"short_url"`
	LongUrl   string    `json:"long_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

type DeleteUrlRequest struct {
	ID string `json:"id"`
}

type DeleteUrlResponse struct {
	Message string `json:"message"`
}
