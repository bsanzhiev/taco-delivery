package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bsanzhiev/taco-delivery/gateway/services"
)

func main() {
	// Context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Signals for graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Initialize gRPC clients for microservices
	// Customer client
	customerClient, err := services.NewCustomerServiceClient("customer-service:8081")
	if err != nil {
		log.Fatalf("Failed to connect to customer-service: %v", err)
	}
	defer customerClient.Close()

	// Order client
	orderClient, err := services.NewOrderServiceClient("order-serviceservice:8082")
	if err != nil {
		log.Fatalf("Failed to connect to order-service: %v", err)
	}
	defer orderClient.Close()

	// Delivery client
	deliveryClient, err := services.NewDeliveryServiceClient("delivery-service:8083")
	if err != nil {
		log.Fatalf("Failed to connect to delivery-service: %v", err)
	}
	defer deliveryClient.Close()

	// Setting up router

	// Start HTTP server
	// Handle signals for graceful shutdown
	// Graceful shutdown
	// Test line
}
