package routes

import (
	"github.com/bsanzhiev/taco-delivery/gateway/handlers"
	"github.com/bsanzhiev/taco-delivery/gateway/services"
	"github.com/gorilla/mux"
)

func SetupRoutes(
	customerClient *services.CustomerServiceClient,
	orderClient *services.OrderServiceClient,
	deliveryClient *services.DeliveryServiceClient,
) *mux.Router {
	router := mux.NewRouter()

	// Customer routes
	router.HandleFunc("/api/v1/customers", handlers.CreateCustomer(customerClient)).Methods("POST")

	// Order routes
	router.HandleFunc("/api/v1/orders", handlers.CreateOrder(orderClient)).Methods("POST")

	// Deliver routes
	router.HandleFunc("api/v1/delivers", handlers.UpdateDeliveryStatus(deliveryClient)).Methods("PUT")

	return router
}
