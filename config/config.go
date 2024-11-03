package config

import (
	"log"
	"os"

	"github.com/joho/godotenv" // Import for loading .env files
)

// Config holds the application configuration
type Config struct {
	Port       string
	DatabaseURL string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	return Config{
		Port:       getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/waterdb"),
	}
}

// Helper function to get environment variables or fallback to a default
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
