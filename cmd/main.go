package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	cfg "github.com/hosseinal/UrlShortner/internal/config"
	redisCache "github.com/hosseinal/UrlShortner/internal/database/cache/redisCache"
	postgres "github.com/hosseinal/UrlShortner/internal/database/sql/postgresdb"
	handlers "github.com/hosseinal/UrlShortner/internal/handler"
	middleware "github.com/hosseinal/UrlShortner/internal/middleware"
	cachedUrl "github.com/hosseinal/UrlShortner/internal/repository/cachedUrl"
	urlRepository "github.com/hosseinal/UrlShortner/internal/repository/url"
	userRepository "github.com/hosseinal/UrlShortner/internal/repository/user"
	route "github.com/hosseinal/UrlShortner/internal/route"
	service "github.com/hosseinal/UrlShortner/internal/service"
)

func main() {

	accessSecret := os.Getenv("ACCESS_SECRET")
	refreshSecret := os.Getenv("REFRESH_SECRET")
	if accessSecret == "" || refreshSecret == "" {
		log.Fatalf("ACCESS_SECRET and REFRESH_SECRET must be set")
	}

	// setup postgres
	postgresPortStr := os.Getenv("POSTGRES_PORT")
	if postgresPortStr == "" {
		postgresPortStr = "5432"
	}
	postgresPort, err := strconv.Atoi(postgresPortStr)
	if err != nil {
		log.Fatalf("failed to convert POSTGRES_PORT to int: %v", err)
	}

	postgresHost := os.Getenv("POSTGRES_HOST")
	if postgresHost == "" {
		postgresHost = "localhost"
	}

	postgresUser := os.Getenv("POSTGRES_USER")
	if postgresUser == "" {
		postgresUser = "postgres"
	}

	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	if postgresPassword == "" {
		postgresPassword = "postgres"
	}

	postgresDBName := os.Getenv("POSTGRES_DB_NAME")
	if postgresDBName == "" {
		postgresDBName = "url_shortner"
	}

	// setup redis
	redisPortStr := os.Getenv("REDIS_PORT")
	if redisPortStr == "" {
		redisPortStr = "6379"
	}
	redisPort, err := strconv.Atoi(redisPortStr)
	if err != nil {
		log.Fatalf("failed to convert REDIS_PORT to int: %v", err)
	}

	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost"
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")

	// setup config
	config := cfg.NewSQLDatabaseConfig(postgresHost, postgresPort, postgresUser, postgresPassword, postgresDBName)

	redisConfig := cfg.NewRedisCacheConfig(redisHost, redisPort, redisPassword)

	// setup database
	db, err := postgres.NewPostgresDB(config)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// setup redis
	redisClient, err := redisCache.NewRedisCache(redisConfig)
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}

	// setup url repository
	urlRepository := urlRepository.NewUrlRepository(db)
	cachedUrlRepository := cachedUrl.NewCachedUrlRepository(redisClient.Client)
	userRepository := userRepository.NewUserRepository(db)

	// setup services
	authService := service.NewAuthService(userRepository)
	urlService := service.NewUrlService(urlRepository, cachedUrlRepository)

	// setup handlers
	authHandler := handlers.NewAuthHandler(authService)
	urlHandler := handlers.NewUrlHandler(urlService)
	jwtMiddleware := middleware.NewJWT(os.Getenv("ACCESS_SECRET"), os.Getenv("REFRESH_SECRET"), authService)

	// setup router
	router := gin.Default()
	route := route.NewRoute(*authHandler, *urlHandler, *jwtMiddleware)
	route.SetupAuthRoutes(router)
	route.SetupUrlRoutes(router)
	router.Run(":8080")
}
