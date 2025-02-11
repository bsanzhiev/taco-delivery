package services

import (
	"context"

	"github.com/bsanzhiev/taco-delivery/common/models"
	"google.golang.org/grpc"
)

// CustomerServiceClient представляет клиент для взаимодействия с user-service.
type CustomerServiceClient struct {
	client models.CustomerServiceClient
	conn   *grpc.ClientConn
}

func NewCustomerServiceClient(addr string) (*CustomerServiceClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithAuthority())
	if err != nil {
		return nil, err
	}
	return &CustomerServiceClient{
		client: models.NewCustomerServiceClient(conn),
		conn:   conn,
	}, nil
}

// Close закрывает соединение с user-service.
func (c *CustomerServiceClient) Close() error {
	return c.conn.Close()
}

// CreateUser вызывает метод CreateUser в user-service.
func (c *CustomerServiceClient) CreateCustomer(ctx context.Context, customer *models.Customer) (*models.CustomerResponse, error) {
	return &c.client.CreateCustomer(ctx, customer)
}

func (c *CustomerServiceClient) GetCustomer(ctx context.Context, customerID *models.CustomerID) (*models.Customer, error) {
	return c.client.GetCustomer(ctx, customerID)
}
