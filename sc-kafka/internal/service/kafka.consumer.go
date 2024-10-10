package service

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

// todo: implement it in sc-notification

type KafkaConsumer struct {
	Consumer *kafka.Consumer
}

func NewKafkaConsumer(brokers, groupID string) *KafkaConsumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
		return nil
	}
	return &KafkaConsumer{
		Consumer: c,
	}
}

func (kc *KafkaConsumer) Subscribe(topics []string) error {
	return kc.Consumer.SubscribeTopics(topics, nil)
}

func (kc *KafkaConsumer) PollMessages() {
	for {
		msg, err := kc.Consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
			continue
		}
		log.Printf("Received message: %s\n", string(msg.Value))
	}
}
