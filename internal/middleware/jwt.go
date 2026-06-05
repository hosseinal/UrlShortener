package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hosseinal/UrlShortner/internal/service"
	"github.com/hosseinal/UrlShortner/internal/toolbox"
)

type JWT struct {
	AccessSecret  string
	RefreshSecret string
	authService   service.AuthService
}

func NewJWT(accessSecret string, refreshSecret string, authService service.AuthService) *JWT {
	return &JWT{AccessSecret: accessSecret, RefreshSecret: refreshSecret, authService: authService}
}

func (j *JWT) JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		}

		claims, err := toolbox.ParseJWT(token, j.AccessSecret)
		switch err {
		case toolbox.ErrorAccessExpired:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token expired"})
			c.Abort()
		case toolbox.ErrorInvalidToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		case toolbox.ErrorInvalidClaims:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
			c.Abort()
		default:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		}

		user, err := j.authService.GetUserByID(c.Request.Context(), claims.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
		}

		c.Set("user", user)
		c.Next()
	}
}
