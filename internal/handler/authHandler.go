package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hosseinal/UrlShortner/internal/dto"
	"github.com/hosseinal/UrlShortner/internal/service"
	"github.com/hosseinal/UrlShortner/internal/toolbox"
)

type AuthHandler struct {
	authService   service.AuthService
	accessSecret  string
	refreshSecret string
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var registerRequest dto.RegisterRequest
	c.ShouldBindJSON(&registerRequest)

	// create user from authService
	user, err := h.authService.Register(c.Request.Context(), registerRequest.Username, registerRequest.Email, registerRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// create a new JWT token
	accessToken, refreshToken, err := toolbox.GenerateJWT(user.Username, h.accessSecret, h.refreshSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	registerResponse := dto.RegisterResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusCreated, registerResponse)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest
	c.ShouldBindJSON(&loginRequest)

	user, err := h.authService.Login(c.Request.Context(), loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := toolbox.GenerateJWT(user.Username, h.accessSecret, h.refreshSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	loginResponse := dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, loginResponse)
}
