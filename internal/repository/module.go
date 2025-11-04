package repository

import (
	"go.uber.org/fx"
)

var Module = fx.Module("modules", fx.Provide(
	NewUserRepository,
))
