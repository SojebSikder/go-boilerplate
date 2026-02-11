package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/sojebsikder/go-boilerplate/internal/config"
)

type Redis struct {
	Config *config.Config
	Client *redis.Client
}

func NewRedis(config *config.Config) (*Redis, error) {
	opt, _ := redis.ParseURL(config.Redis.RedisURL)
	client := redis.NewClient(opt)

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return &Redis{
		Client: client,
	}, nil
}

func (r *Redis) Close() error {
	return r.Client.Close()
}
