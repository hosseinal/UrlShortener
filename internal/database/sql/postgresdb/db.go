package postgres

import (
	"database/sql"

	"github.com/hosseinal/UrlShortner/internal/config"

	"fmt"
	"time"

	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB(config config.SQLDatabaseConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s timezone=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, "require", "UTC",
	)

	db, err := gorm.Open(gormPostgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("getting raw DB: %w", err)
	}
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	return db, nil
}
