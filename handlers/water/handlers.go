package water

import (
	"encoding/json"
	"net/http"
	"github.com/yourusername/water-management-system/models"
)

// GetWaterUsage handles the retrieval of water usage records
func GetWaterUsage(w http.ResponseWriter, r *http.Request) {
	// Sample data for now
	waterUsage := []models.WaterUsage{
		{ID: 1, UserID: 101, Amount: 250.5, DateRecorded: "2024-11-03"},
		{ID: 2, UserID: 102, Amount: 300.0, DateRecorded: "2024-11-02"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(waterUsage)
}
