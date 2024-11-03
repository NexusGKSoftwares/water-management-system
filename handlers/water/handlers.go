package water

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/yourusername/water-management-system/models"
)

var db *sql.DB

// InitializeHandler sets the database connection
func InitializeHandler(database *sql.DB) {
	db = database
}

// GetWaterUsage handles the retrieval of water usage records
func GetWaterUsage(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, user_id, amount, date_recorded FROM water_usage")
	if err != nil {
		http.Error(w, "Error fetching data from the database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var waterUsage []models.WaterUsage
	for rows.Next() {
		var usage models.WaterUsage
		if err := rows.Scan(&usage.ID, &usage.UserID, &usage.Amount, &usage.DateRecorded); err != nil {
			http.Error(w, "Error scanning data", http.StatusInternalServerError)
			return
		}
		waterUsage = append(waterUsage, usage)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(waterUsage)
}
