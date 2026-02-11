package user

import (
	"go.uber.org/fx"
)

var HTTPModule = fx.Module("user-http",
	fx.Provide(
		NewUserController,
	),
	fx.Invoke(RegisterRoutes),
)
