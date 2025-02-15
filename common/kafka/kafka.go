package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/internal/encoding/messageset"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewProducer(broker string) *KafkaProducer {
	return &KafkaProducer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(broker),
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *KafkaProducer) Publish(topic string, message any) error {
	payload, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}
	err = p.writer.WriteMessages(context.Background(), kafka.Message{
		Topic: topic,
		Value: payload,
	})
	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}
	return nil
}

func (p *KafkaProducer) Close() error {
	return p.writer.Close()
}

type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewConsumer(broker, topic string) *KafkaConsumer {
	return &KafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:     []string{broker},
			Topic:       topic,
			GroupID:     "group-id",
			StartOffset: kafka.LastOffset,
		}),
	}
}

func (c *KafkaConsumer) Subscribe(handler func(message []byte)) {
	for {
		msg, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error consuming message: %v", err)
			continue
		}
		handler(msg.Value)
	}
}

func (c *KafkaConsumer) Close() error {
	return c.reader.Close()
}
