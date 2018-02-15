package db

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/orvice/kit/mod"
)

var (
	redisClients map[string]*redis.Client

	RedisClientNotFoundErr = fmt.Errorf("Redis client not founnd. ")
)

func InitRedis(cfgs map[string]mod.Redis) error {
	if redisClients == nil {
		redisClients = make(map[string]*redis.Client, 0)
	}
	for k, v := range cfgs {
		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", v.Host, v.Port),
			Password: v.Password, // no password set
			DB:       v.DB,       // use default DB
		})
		redisClients[k] = client
	}
	return nil
}

func GetRedis(k string) (*redis.Client, error) {
	c, ok := redisClients[k]
	if !ok {
		return nil, RedisClientNotFoundErr
	}
	return c, nil
}
