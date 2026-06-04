package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hosseinal/UrlShortner/internal/service"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthHandler) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var registerRequest dto.RegisterRequest
	c.ShouldBindJSON(&registerRequest)

	// create user from authService
	err := h.authService.Register(c.Request.Context(), registerRequest.Username, registerRequest.Email, registerRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// create a new JWT token

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
