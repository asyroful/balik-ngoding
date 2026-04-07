package database

import (
	"log"
	"os"

	"balik-ngoding-backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=balik_ngoding port=5434 sslmode=disable"
	}
	// GORM postgres driver supports both DSN format (host=...) and URL format (postgres://...)
	// Railway injects DATABASE_URL as postgres:// URL which is handled automatically

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := DB.AutoMigrate(&models.Problem{}, &models.TestCase{}, &models.Submission{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database connected and migrations applied")
}
