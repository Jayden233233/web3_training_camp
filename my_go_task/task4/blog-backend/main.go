package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/my_go_task/task4/blog-backend/config"
	"github.com/my_go_task/task4/blog-backend/routes"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database
	config.ConnectDB()
	config.MigrateDB()

	// Set JWT secret
	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	// Set up router
	router := routes.SetupRouter()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	log.Fatal(router.Run(":" + port))
}