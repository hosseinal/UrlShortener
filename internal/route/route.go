package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hosseinal/UrlShortner/internal/handler"
	"github.com/hosseinal/UrlShortner/internal/middleware"
)

type Route struct {
	authHandler   handler.AuthHandler
	urlHandler    handler.UrlHandler
	jwtMiddleware middleware.JWT
}

func (r *Route) SetupAuthRoutes(router *gin.Engine) {
	router.POST("/register", r.authHandler.Register)
	router.POST("/login", r.authHandler.Login)
}

func (r *Route) SetupUrlRoutes(router *gin.Engine) {
	router.POST("/create-url", r.urlHandler.CreateUrl)
	router.POST("/get-url", r.urlHandler.GetUrl)
	router.POST("/redirect-url", r.urlHandler.RedirectUrl)
	router.POST("/delete-url", r.urlHandler.DeleteUrl)
}
