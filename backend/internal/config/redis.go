package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func InitRedis(redisURL string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisURL,
	})

	// Test connection
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("WARNING: Redis connection failed: %v. Application will run without Redis.\n", err)
		return nil, err // Return nil so services know it's not available
	}

	log.Println("Redis connection established")
	return rdb, nil
}
