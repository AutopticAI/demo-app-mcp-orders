// Package api serves the Orders HTTP contract.
package api

import (
	"encoding/json"
	"net/http"
)

// OrderResponse is the body GET /orders/{id} returns.
//
// The customer is no longer carried here: it moved behind /orders/{id}/party
// so an order can be read without exposing who placed it.
type OrderResponse struct {
	OrderID    string `json:"order_id"`
	TotalCents int    `json:"total_cents"`
	Currency   string `json:"currency"`
	Status     string `json:"status"`
}

// GetOrder writes one order as JSON.
func GetOrder(w http.ResponseWriter, r *http.Request) {
	order := OrderResponse{
		OrderID:    r.PathValue("id"),
		TotalCents: 4250,
		Currency:   "USD",
		Status:     "settled",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(order); err != nil {
		http.Error(w, "encode order", http.StatusInternalServerError)
	}
}
