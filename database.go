package config

import (
	"fmt"                     // For printing to console
	"log"                     // For logging errors
	"os"                      // For accessing environment variables
	"gorm.io/driver/postgres" // GORM PostgreSQL driver
	"gorm.io/gorm"            // GORM ORM package
	"bookapi/models"          // Your schema models
)

// Global DB connection
var DB *gorm.DB

// ConnectDatabase initializes the database connection using environment variables
func ConnectDatabase() {
	// Build DSN string using environment variables from .env
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Try to connect to PostgreSQL
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database! ", err)
	}

	// Auto-migrate your models (creates tables if not present)
	// Order matters: User first, then Books & Movies (since they reference UserID)
	err = database.AutoMigrate(&models.User{}, &models.Books{}, &models.Movies{})
	if err != nil {
		log.Fatal("❌ Migration failed: ", err)
	}

	fmt.Println("✅ Database connected and migrated successfully!")

	// Save the connection globally
	DB = database
}
