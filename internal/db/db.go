package db

import (
	"database/sql"
	"log"

	"github.com/yourusername/water-management-system/config"
	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

// ConnectDatabase sets up a connection to the MySQL database
func ConnectDatabase() *sql.DB {
	cfg := config.LoadConfig()
	db, err := sql.Open("mysql", cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not ping the database: %v", err)
	}

	log.Println("Successfully connected to the database")
	return db
}
