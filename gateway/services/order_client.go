package services

import (
	"context"

	"github.com/bsanzhiev/taco-delivery/common/models"
	"google.golang.org/grpc"
)

type OrderServiceClient struct {
	client models.OrderServiceClient
	conn   *grpc.ClientConn
}

func NewOrderServiceClient(addr string) (*OrderServiceClient, error) {
	conn, err := grpc.NewClient(addr) // TODO: add options
	if err != nil {
		return nil, err
	}
	return &OrderServiceClient{
		client: models.NewOrderServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *OrderServiceClient) Close() error {
	return c.conn.Close()
}

func (c *OrderServiceClient) CreateOrder(ctx context.Context, order *models.Order) (*models.OrderResponse, error) {
	return c.client.CreateOrder(ctx, order)
}

func (c *OrderServiceClient) GetOrder(ctx context.Context, orderID *models.OrderID) (*models.Order, error) {
	return c.client.GetOrder(ctx, orderID)
}
