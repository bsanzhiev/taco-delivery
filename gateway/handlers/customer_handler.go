package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/bsanzhiev/taco-delivery/common/models"
	"github.com/bsanzhiev/taco-delivery/gateway/services"
)

func CreateCustomer(customerClient *services.CustomerServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customer models.Customer
		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		resp, err := customerClient.CreateCustomer(r.Context(), &customer)
		if err != nil {
			http.Error(w, "Failed to create customer", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
