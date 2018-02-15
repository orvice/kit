package cache

import (
	"time"

	"github.com/go-redis/redis"
)

type SimpleCache interface {
	Set(key, value string, ttl time.Duration) error
	Get(key string) (string, bool, error)
}

type RedisSimpleCache struct {
	client *redis.Client
}

func (r *RedisSimpleCache) Set(key, value string, ttl time.Duration) error {
	return r.client.Set(key, value, ttl).Err()
}

func (r *RedisSimpleCache) Get(key string) (string, bool, error) {
	i, err := r.client.Exists(key).Result()
	if err != nil {
		return "", false, err
	}
	if i == 0 {
		return "", false, nil
	}

	s, err := r.client.Get(key).Result()
	return s, true, nil
}
