package handlers

import (
	"context"

	"github.com/bsanzhiev/taco-delivery/common/models"
	"github.com/bsanzhiev/taco-delivery/customer-service/repositories"
)

type CustomerHandler struct {
	repo repositories.CustomerRepository
}

func NewCustomerhandler(repo repositories.CustomerRepository) *CustomerHandler {
	return &CustomerHandler{
		repo: repo,
	}
}

func (h *CustomerHandler) CreateCustomer(ctx context.Context, customer *models.Customer) (*models.CustomerResponse, error) {
	id, err := h.repo.CreateCustomer(customer)
	if err != nil {
		return nil, err
	}

	return &models.CustomerResponse{
		Id:     id,
		Status: "created",
	}, nil
}

func (h *CustomerHandler) GetCustomer(ctx context.Context, customerID *models.CustomerID) (*models.Customer, error) {
	customer, err := h.repo.GetCustomer(customerID.Id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
