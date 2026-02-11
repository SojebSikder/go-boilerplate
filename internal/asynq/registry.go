package asynqfx

import (
	authtask "github.com/sojebsikder/go-boilerplate/internal/modules/auth/task"
	"go.uber.org/fx"
)

var Registry = fx.Module("asynqregistryfx",
	fx.Provide(
		authtask.NewAuthProcessHandler,
	),
)
