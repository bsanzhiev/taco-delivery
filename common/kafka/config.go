package kafka

import (
	"fmt"
	"os"
)

type Config struct {
	Broker string
}

func LoadConfig() (*Config, error) {
	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		return nil, fmt.Errorf("KAFKA BROKER is not set")
	}
	return &Config{
		Broker: broker,
	}, nil
}
