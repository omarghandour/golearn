package db

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/omarghandour/test1/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes SQLite and runs automatic migrations
func InitDB() {
	var err error

	// Support configurable DB path for container volumes (defaults to local books.db)
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "books.db"
	}

	// Open or create the SQLite database file
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database at %s: %v", dbPath, err)
	}

	// Automatically create or update the "books" table schema
	if err := DB.AutoMigrate(&models.Book{}); err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}

	log.Printf("SQLite database connected (%s) and migrated successfully.\n", dbPath)
}
