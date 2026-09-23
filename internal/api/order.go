// Package api serves the Orders HTTP contract.
package api

import (
	"encoding/json"
	"net/http"
)

// OrderResponse is the body GET /orders/{id} returns.
type OrderResponse struct {
	OrderID    string `json:"order_id"`
	CustomerID string `json:"customer_id"`
	TotalCents int    `json:"total_cents"`
	Currency   string `json:"currency"`
	Status     string `json:"status"`
	// Channel is the surface the order was placed through: web, ios or partner.
	Channel string `json:"channel,omitempty"`
}

// GetOrder writes one order as JSON.
func GetOrder(w http.ResponseWriter, r *http.Request) {
	order := OrderResponse{
		OrderID:    r.PathValue("id"),
		CustomerID: "cus_8811",
		TotalCents: 4250,
		Currency:   "USD",
		Status:     "settled",
		Channel:    "web",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(order); err != nil {
		http.Error(w, "encode order", http.StatusInternalServerError)
	}
}
