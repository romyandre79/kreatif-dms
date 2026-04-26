package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheService struct {
	client *redis.Client
}

func NewCacheService(client *redis.Client) *CacheService {
	return &CacheService{client: client}
}

func (s *CacheService) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	if s.client == nil {
		return nil // Graceful bypass
	}
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.client.Set(ctx, key, data, ttl).Err()
}

func (s *CacheService) Get(ctx context.Context, key string, dest interface{}) (bool, error) {
	if s.client == nil {
		return false, nil // Graceful bypass
	}
	val, err := s.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}

	err = json.Unmarshal([]byte(val), dest)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *CacheService) Delete(ctx context.Context, key string) error {
	if s.client == nil {
		return nil // Graceful bypass
	}
	return s.client.Del(ctx, key).Err()
}
