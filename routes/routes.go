package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/yourusername/water-management-system/handlers/water"
)

// SetupRoutes initializes all the routes for the application
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Water usage routes
	r.HandleFunc("/api/v1/water-usage", water.GetWaterUsage).Methods("GET")

	return r
}
