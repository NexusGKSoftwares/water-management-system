package models

// WaterUsage represents a record of water consumption
type WaterUsage struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	Amount      float64 `json:"amount"`
	DateRecorded string `json:"date_recorded"`
}
