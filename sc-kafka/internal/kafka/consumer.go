package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

type Consumer struct {
	Reader  *kafka.Reader
	Handler MessageHandler
}

func NewConsumer(broker string, topic string, groupID string, handler MessageHandler) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		GroupID: groupID,
		Topic:   topic,
	})

	return &Consumer{
		Reader:  reader,
		Handler: handler,
	}
}

func (c *Consumer) StartConsuming(ctx context.Context) {
	for {
		m, err := c.Reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Ошибка при чтении сообщения: %v", err)
			break
		}

		err = c.Handler.HandleMessage(m)
		if err != nil {
			log.Printf("Error proccessing message: %v", err)
			// Можно добавить логику повторной обработки или отправки в dead-letter queue
		}
	}
}

func (c *Consumer) Close() error {
	return c.Reader.Close()
}
