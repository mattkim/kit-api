package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/pbk/kit-api/models"
)

// SetupDB db
func SetupDB() *gorm.DB {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	// defer db.Close()
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	// db.DropTableIfExists(&models.User{})
	// db.DropTableIfExists(&models.Event{})
	// db.DropTableIfExists(&models.Message{})
	db.CreateTable(&models.User{})
	db.CreateTable(&models.Event{})
	db.CreateTable(&models.Message{})
	db.LogMode(true)

	return db
}
