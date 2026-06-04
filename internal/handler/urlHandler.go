package handler

import (
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hosseinal/UrlShortner/internal/dto"
	"github.com/hosseinal/UrlShortner/internal/model"
	"github.com/hosseinal/UrlShortner/internal/service"

	"github.com/google/uuid"
	"github.com/hosseinal/UrlShortner/internal/toolbox"
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

	if !toolbox.IsValidURL(url.LongUrl) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	modelUrl := model.Url{
		ID:        uuid.New(),
		LongUrl:   url.LongUrl,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	shortUrl := generate256BitCode(url.LongUrl)
	modelUrl.ShortUrl = shortUrl

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

func (h *UrlHandler) GetUrl(c *gin.Context) {
	var url dto.GetUrlRequest
	c.ShouldBindJSON(&url)

	urlResponse, err := h.urlService.GetUrlByShortUrl(c.Request.Context(), url.ShortUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.GetUrlResponse{
		ID:        urlResponse.ID.String(),
		ShortUrl:  urlResponse.ShortUrl,
		LongUrl:   urlResponse.LongUrl,
		CreatedAt: urlResponse.CreatedAt,
		UpdatedAt: urlResponse.UpdatedAt,
		ExpiresAt: urlResponse.ExpiresAt,
	})
}

func (h *UrlHandler) RedirectUrl(c *gin.Context) {
	var url dto.GetUrlRequest
	c.ShouldBindJSON(&url)

	urlResponse, err := h.urlService.GetUrlByShortUrl(c.Request.Context(), url.ShortUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, urlResponse.LongUrl)
}

func (h *UrlHandler) DeleteUrl(c *gin.Context) {
	var url dto.DeleteUrlRequest
	c.ShouldBindJSON(&url)

	err := h.urlService.DeleteUrl(c.Request.Context(), url.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.DeleteUrlResponse{Message: "Url deleted successfully"})
}
