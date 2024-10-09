package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"soul-connect/sc-auth/internal/config"
	"soul-connect/sc-auth/internal/generated"
	"soul-connect/sc-auth/internal/server"
	"soul-connect/sc-auth/internal/services"
	"soul-connect/sc-auth/pkg/database"
)

func main() {
	// Load configuration
	newConfig, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	newPool, err := database.NewPostgresDB(&newConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer newPool.Close()

	lis, err := net.Listen("tcp", ":"+newConfig.ServerPort)
	if err != nil {
		log.Fatalf("failed to listen on gRPC port %s: %v", newConfig.ServerPort, err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	newService := services.NewService(newPool)
	newServer := server.NewAuthServer(newService)

	generated.RegisterAuthServiceServer(grpcServer, newServer)

	// Start serving gRPC
	log.Printf("gRPC server listening on port %s", newConfig.ServerPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}
