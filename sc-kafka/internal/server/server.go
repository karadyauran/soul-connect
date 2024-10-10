package server

import (
	"context"
	"soul-connect/sc-kafka/internal/generated"
	"soul-connect/sc-kafka/internal/model"
	"soul-connect/sc-kafka/internal/service"
)

type KafkaServer struct {
	generated.UnimplementedKafkaServiceServer
	kafkaService *service.KafkaProducer
}

func NewKafkaServer(kafkaService *service.Service) *KafkaServer {
	return &KafkaServer{
		kafkaService: kafkaService.KafkaProducer,
	}
}

func (ks *KafkaServer) KafkaSend(ctx context.Context, request *generated.MessageRequest) (*generated.MessageResponse, error) {
	sendMessageRequest := model.Message{
		ID:        request.Id,
		Content:   request.Content,
		CreatedAt: request.CreatedAt,
	}

	err := ks.kafkaService.Send(sendMessageRequest)
	if err != nil {
		return &generated.MessageResponse{
			Response: "Error creating message",
		}, err
	}

	return &generated.MessageResponse{Response: "Message sent"}, nil
}
