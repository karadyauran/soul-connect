package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

type Producer struct {
	Writer *kafka.Writer
}

func NewProducer(broker string, topic string) *Producer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   topic,
	})

	return &Producer{
		Writer: writer,
	}
}

func (p *Producer) SendMessage(ctx context.Context, key, value []byte) error {
	msg := kafka.Message{
		Key:   key,
		Value: value,
	}

	err := p.Writer.WriteMessages(ctx, msg)
	if err != nil {
		log.Printf("Error sending message: %v", err)
		return err
	}

	log.Printf("Message sent: key=%s", string(key))
	return nil
}

func (p *Producer) Close() error {
	return p.Writer.Close()
}
