package services

import (
	"context"

	"github.com/bsanzhiev/taco-delivery/common/models"
	"google.golang.org/grpc"
)

type DeliveryServiceClient struct {
	client models.DeliveryServiceClient
	conn   *grpc.ClientConn
}

func NewDeliveryServiceClient(addr string) (*DeliveryServiceClient, error) {
	conn, err := grpc.NewClient(addr)
	if err != nil {
		return nil, err
	}
	return &DeliveryServiceClient{
		client: models.NewDeliveryServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *DeliveryServiceClient) Close() error {
	return c.conn.Close()
}

func (c *DeliveryServiceClient) UpdateDeliveryStatus(ctx context.Context, status *models.DeliveryStatus) (*models.DeliveryResponse, error) {
	return c.client.UpdateDeliveryStatus(ctx, status)
}

func (c *DeliveryServiceClient) GetDelivery(ctx context.Context, deliveryID *models.DeliveryID) (*models.Delivery, error) {
	return c.client.GetDelivery(ctx, deliveryID)
}
