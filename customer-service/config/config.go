package config

import "os"

type Config struct {
	DB   string
	GRPC struct {
		Port string
	}
	Kafka struct {
		Broker string
	}
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		DB: os.Getenv("DB_URL"),
	}
	cfg.GRPC.Port = os.Getenv("GRPC_PORT")
	cfg.Kafka.Broker = os.Getenv("KAFKA_BROKER")

	return cfg, nil
}
