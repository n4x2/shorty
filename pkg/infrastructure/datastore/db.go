package datastore

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDB creates a new database connection.
func NewDB() *gorm.DB {
	// Load environment variables from .env file.
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	// Open a new connection to database using DSN from environment.
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to PlanetScale: %v", err)
	}

	return db
}
