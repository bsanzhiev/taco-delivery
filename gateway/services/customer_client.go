package services

import (
	"context"

	modelspb "github.com/bsanzhiev/taco-delivery/common/models"
	models "github.com/bsanzhiev/taco-delivery/common/models/customer"
	"google.golang.org/grpc"
)

// CustomerServiceClient представляет клиент для взаимодействия с user-service.
type CustomerServiceClient struct {
	client modelspb.CustomerServiceClient
	conn   *grpc.ClientConn
}

func NewCustomerServiceClient(addr string) (*CustomerServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithAuthority(addr))
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
func (c *CustomerServiceClient) CreateCustomer(ctx context.Context, customer *modelspb.Customer) (*modelspb.CustomerResponse, error) {
	return &c.client.CreateCustomer(ctx, customer)
}

func (c *CustomerServiceClient) GetCustomer(ctx context.Context, customerID *modelspb.CustomerID) (*modelspb.Customer, error) {
	return c.client.GetCustomer(ctx, customerID)
}
