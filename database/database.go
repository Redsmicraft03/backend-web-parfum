package database

import (
	"log"
	"os"
	"web-parfum/backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	log.Println("database connected")

	// buat table users
	DB.AutoMigrate(&models.User{})
	log.Println("migrate user")

	// buat table products
	DB.AutoMigrate(&models.Product{})
	log.Println("migrate product")
}