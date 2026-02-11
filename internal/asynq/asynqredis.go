package asynqfx

import (
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/sojebsikder/go-boilerplate/internal/config"
)

func AsynqRedisOpt(cfg *config.Config) (asynq.RedisClientOpt, error) {
	opt, err := asynq.ParseRedisURI(cfg.Redis.RedisURL)
	if err != nil {
		return asynq.RedisClientOpt{}, fmt.Errorf("invalid redis url: %w", err)
	}

	redisOpt, ok := opt.(asynq.RedisClientOpt)
	if !ok {
		return asynq.RedisClientOpt{}, fmt.Errorf("redis url is not a valid RedisClientOpt")
	}

	return redisOpt, nil
}
