package auth

import (
	"go.uber.org/fx"
)

var Module = fx.Module("auth",
	fx.Provide(
		NewAuthController,
		NewAuthService,
	),
	fx.Invoke(RegisterRoutes),
)
