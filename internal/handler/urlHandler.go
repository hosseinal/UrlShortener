package handler

import (
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hosseinal/UrlShortner/internal/dto"
	"github.com/hosseinal/UrlShortner/internal/model"
	"github.com/hosseinal/UrlShortner/internal/service"

	"github.com/google/uuid"
)

type UrlHandler struct {
	urlService service.UrlService
}

func NewUrlHandler(urlService service.UrlService) *UrlHandler {
	return &UrlHandler{urlService: urlService}
}

func (h *UrlHandler) CreateUrl(c *gin.Context) {
	var url dto.CreateUrlRequest
	c.ShouldBindJSON(&url)

	modelUrl := model.Url{
		ID:        uuid.New(),
		LongUrl:   url.LongUrl,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err := h.urlService.CreateUrl(c.Request.Context(), &modelUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.CreateUrlResponse{
		ID:        modelUrl.ID.String(),
		ShortUrl:  modelUrl.ShortUrl,
		LongUrl:   modelUrl.LongUrl,
		CreatedAt: modelUrl.CreatedAt,
		UpdatedAt: modelUrl.UpdatedAt,
		ExpiresAt: modelUrl.ExpiresAt,
	})

}
