package model

import (
	"time"

	uuid "github.com/google/uuid"
)

type Url struct {
	ID        uuid.UUID `json:"id"`
	ShortUrl  string    `json:"short_url"`
	LongUrl   string    `json:"long_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (u *Url) TableName() string {
	return "urls"
}
