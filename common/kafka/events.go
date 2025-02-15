package kafka

import (
	"encoding/json"
	"fmt"

	"github.com/bsanzhiev/taco-delivery/common/models"
)

func PublishConsumerCreatedEvent(producer *KafkaProducer, event *models.CustomerCreatedEvent) error {
	return producer.Publish("consumer-created", event)
}

func ParseCostumerCreatedEvent(data []byte) (*models.CustomerCreatedEvent, error) {
	var event models.CustomerCreatedEvent
	err := json.Unmarshal(data, &event)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CustomerCreatedEvent: %v", err)
	}
	return &event, nil
}

func PublishOrderPlacedEvent(producer *KafkaProducer, event *models.OrderPlacedEvent) error {
	return producer.Publish("order-placed", event)
}

func ParseOrderPlacedEvent(data []byte) (*models.OrderPlacedEvent, error) {
	var event models.OrderPlacedEvent
	err := json.Unmarshal(data, &event)
	if err != nil {
		return nil, fmt.Errorf("failed to parse OrderPlacedEvent: %v", err)
	}
	return &event, nil
}
