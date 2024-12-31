package main

import (
	"log"
	"os"

	"github.com/Mohamed-Kalandar-Sulaiman/Identity_Service/src/database"
	"github.com/Mohamed-Kalandar-Sulaiman/Identity_Service/src/models"
	"github.com/Mohamed-Kalandar-Sulaiman/Identity_Service/src/routers"
	"github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
)

func main() {

    // Load environment variables from .env file
	 if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }
    // Connect to database
    if err := database.ConnectDatabase(); err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }

    // Auto migrate the User and Application models
    database.DB.AutoMigrate(&models.User{}, &models.Application{})

    // Initialize Fiber app
    app := fiber.New()

    // Setup routes
    routers.SetupAuthRoutes(app)

    // Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    log.Fatal(app.Listen(":" + port))
}
