package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"soul-connect/sc-kafka/internal/config"
	"soul-connect/sc-kafka/internal/generated"
	"soul-connect/sc-kafka/internal/server"
	"soul-connect/sc-kafka/internal/service"
)

func main() {
	newConfig, err := config.LoadKafkaConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	lis, err := net.Listen("tcp", ":"+newConfig.ServerPort)
	if err != nil {
		log.Fatalf("failed to listen on gRPC port %s: %v", newConfig.ServerPort, err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	newService := service.NewService(newConfig.Brokers, newConfig.Topic)
	newServer := server.NewKafkaServer(newService)

	generated.RegisterKafkaServiceServer(grpcServer, newServer)

	// Start serving gRPC
	log.Printf("gRPC server listening on port %s", newConfig.ServerPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}
