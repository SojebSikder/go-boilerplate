package asynqfx

import (
	"go.uber.org/fx"
)

var Module = fx.Module("asynqfx",
	fx.Provide(
		NewClient,
		AsynqRedisOpt,
	),
)
