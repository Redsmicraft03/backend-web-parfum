package main

import (
	"log"
	"os"
	"web-parfum/backend/config"
	"web-parfum/backend/database"
	"web-parfum/backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error memuat file .env")
	}

	config.InitCloudinary()

	database.InitDB()

	app := fiber.New()

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}