package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	err := rdb.ConfigSet(ctx, "stop-writes-on-bgsave-error", "no").Err()
	if err != nil {
		log.Fatalf("Failed to update redis config: %v", err)
	}

	fmt.Println("Successfully set stop-writes-on-bgsave-error to no")
}
