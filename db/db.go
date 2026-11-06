package db

import (
	"log"

	"github.com/glebarez/sqlite"
	"github.com/omarghandour/test1/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes SQLite and runs automatic migrations
func InitDB() {
	var err error
	// Open or create the SQLite database file "books.db"
	DB, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Automatically create or update the "books" table schema
	if err := DB.AutoMigrate(&models.Book{}); err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}

	log.Println("SQLite database connected and migrated successfully.")
}

