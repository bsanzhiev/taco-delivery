package main

import (
	"log"
	"net"

	"github.com/bsanzhiev/taco-delivery/common/models"
	"github.com/bsanzhiev/taco-delivery/customer-service/config"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
	}

	// Repositore init
	repo, err := repositories.NewPostgresRepository(cfg.DB)
	if err != nil {
	}

	// Handlers init
	handler := handlers.NewCustomerHandler(repo)

	// Setting up gRPC server
	lis, err := net.Listen("tcp", cfg.GRPC.Port)
	if err != nil {
		log.Fatalf("Failed listen: %v", err)
	}
	server := grpc.Server()
	models.RegisterCustomerServiceServer(server, handler)

	log.Printf("Starting customer-service on port %s...", cfg.GRPC.Port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
