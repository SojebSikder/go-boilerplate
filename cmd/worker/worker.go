package worker

import (
	"os"

	"github.com/hibiken/asynq"
	"github.com/sojebsikder/go-boilerplate/internal/app"
	asynqfx "github.com/sojebsikder/go-boilerplate/internal/asynq"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

var WorkerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start the worker",
	Run: func(cmd *cobra.Command, args []string) {
		StartWorker()
	},
}

func StartWorker() {
	app := fx.New(
		app.BaseModules(),
		asynqfx.Registry,
		fx.Provide(
			asynqfx.NewMux,
			asynqfx.NewServer,
		),
		fx.Invoke(func(*asynq.Server) {
			// This forces Fx to call NewServer, which registers the lifecycle hooks
		}),
		fx.WithLogger(func() fxevent.Logger {
			return &fxevent.ConsoleLogger{
				W: os.Stdout,
			}
		}),
	)

	app.Run()
}
