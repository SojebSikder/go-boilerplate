package auth

import (
	"go.uber.org/fx"
)

var HTTPModule = fx.Module("auth-http",
	fx.Provide(
		NewAuthController,
	),
	fx.Invoke(RegisterRoutes),
)
