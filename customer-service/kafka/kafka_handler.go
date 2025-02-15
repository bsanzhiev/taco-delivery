package kafka

import (
	"github.com/bsanzhiev/taco-delivery/common/kafka"
	"github.com/bsanzhiev/taco-delivery/common/models"
)

func PublishCustomerCreated(customer *models.Customer) error {
	producer := kafka.NewProducer("kafka:9092")
	defer producer.Close()

	event := &models.CustomerCreatedEvent{
		CustomerID: customer.Id,
		Name:       customer.Name,
		Email:      customer.Email,
		Phone:      customer.Phone,
	}

	return producer.Publish("customer-created", event)
}
