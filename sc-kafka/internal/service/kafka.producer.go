package service

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"soul-connect/sc-kafka/internal/model"
)

type KafkaProducer struct {
	Producer *kafka.Producer
	Topic    string
}

func NewKafkaProducer(brokers, topic string) *KafkaProducer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokers})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &KafkaProducer{
		Producer: p,
		Topic:    topic,
	}
}

func (kp *KafkaProducer) Send(message model.Message) error {
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kp.Topic, Partition: kafka.PartitionAny},
		Value:          []byte(message.Content),
	}
	err := kp.Producer.Produce(msg, nil)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return err
	}
	return nil
}
