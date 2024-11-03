package main

import (
	"log"
	"net/http"

	"github.com/yourusername/water-management-system/config"
	"github.com/yourusername/water-management-system/internal/db"
	"github.com/yourusername/water-management-system/routes"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	database := db.ConnectDatabase()
	defer database.Close()

	// Setup routes
	r := routes.SetupRoutes()

	// Start the server
	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
