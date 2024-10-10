# Build stage
FROM golang:1.19-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the binary
RUN go build -o sc-kafka ./cmd/main

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/sc-kafka .

# Copy the environment variables file
COPY dev.env .

# Export environment variables
ENV $(cat .env | xargs)

# Expose application port (if needed)
EXPOSE 8080

# Run the application
CMD ["./sc-kafka"]