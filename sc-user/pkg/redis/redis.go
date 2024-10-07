package pkg

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var Ctx = context.Background()

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(redisURL string) *RedisClient {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("failed to parse Redis URL: %v", err)
	}

	client := redis.NewClient(opts)

	_, err = client.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	return &RedisClient{
		Client: client,
	}
}
