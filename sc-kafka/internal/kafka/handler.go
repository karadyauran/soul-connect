package kafka

import (
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

type Notification struct {
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}

type MessageHandler interface {
	HandleMessage(kafka.Message) error
}

type NotificationHandler struct{}

func (h *NotificationHandler) HandleMessage(msg kafka.Message) error {
	var notification Notification
	err := json.Unmarshal(msg.Value, &notification)
	if err != nil {
		log.Printf("Ошибка при разборе сообщения: %v", err)
		return err
	}

	log.Printf("Sending message to user %s: %s", notification.UserID, notification.Message)

	return nil
}
