package asynqfx

import "github.com/hibiken/asynq"

func NewClient(opt asynq.RedisClientOpt) *asynq.Client {
	return asynq.NewClient(opt)
}
