start: ## Run the project
	go run cmd/general/main.go

test: ## Run all tests
	go test ./...

install: ## Install dependencies
	go mod tidy

swag: ##  Generate swagger specification
	swag init -g cmd/general/main.go

proto-generate-auth-service:
	protoc --proto_path=../proto --go_out=./internal --go-grpc_out=./internal ../proto/auth.proto

help: ## Display help message
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' Makefile