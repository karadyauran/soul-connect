package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"soul-connect/sc-kafka/internal/config"
	"soul-connect/sc-kafka/internal/kafka"
	"syscall"
	"time"
)

func main() {
	newConfig, err := config.LoadKafkaConfig("./")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	producer := kafka.NewProducer(newConfig.Broker, newConfig.Topic)
	defer producer.Close()

	go func() {
		time.Sleep(2 * time.Second)

		notification := kafka.Notification{
			UserID:  "user123",
			Message: "New message!",
		}

		value, err := json.Marshal(notification)
		if err != nil {
			log.Printf("Error of serializing: %v", err)
			return
		}

		err = producer.SendMessage(ctx, []byte(notification.UserID), value)
		if err != nil {
			log.Printf("Error of sending info: %v", err)
		}
	}()

	handler := &kafka.NotificationHandler{}
	consumer := kafka.NewConsumer(newConfig.Broker, newConfig.Topic, "notification-service", handler)
	defer consumer.Close()

	go consumer.StartConsuming(ctx)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	log.Println("Shutting down gracefully...")
}
