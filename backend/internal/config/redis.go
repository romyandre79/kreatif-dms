package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func InitRedis(redisURL string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisURL,
	})

	// Test connection
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}

	log.Println("Redis connection established")
	return rdb
}
