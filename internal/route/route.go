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

func NewRoute(authHandler handler.AuthHandler, urlHandler handler.UrlHandler, jwtMiddleware middleware.JWT) *Route {
	return &Route{
		authHandler:   authHandler,
		urlHandler:    urlHandler,
		jwtMiddleware: jwtMiddleware,
	}
}

func (r *Route) SetupAuthRoutes(router *gin.Engine) {
	router.POST("/register", r.authHandler.Register)
	router.POST("/login", r.authHandler.Login)
}

func (r *Route) SetupUrlRoutes(router *gin.Engine) {
	router.POST("/:shortUrl", r.urlHandler.RedirectUrl)
	url := router.Group("/url").Use(r.jwtMiddleware.JWT())
	url.POST("/create-url", r.urlHandler.CreateUrl)
	url.POST("/get-url", r.urlHandler.GetUrl)
	url.POST("/delete-url", r.urlHandler.DeleteUrl)
}
