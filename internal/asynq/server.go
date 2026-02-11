package asynqfx

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
	"go.uber.org/fx"
)

func NewServer(
	lc fx.Lifecycle,
	opt asynq.RedisClientOpt,
	mux *asynq.ServeMux,
) *asynq.Server {
	fmt.Println("server running")
	srv := asynq.NewServer(
		opt,
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 5,
				"default":  3,
				"low":      2,
			},
		},
	)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Starting Asynq server...")
			go func() {
				if err := srv.Run(mux); err != nil {
					panic(fmt.Errorf("asynq server failed: %w", err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Shutting down Asynq server...")
			srv.Shutdown()
			return nil
		},
	})

	return srv
}
