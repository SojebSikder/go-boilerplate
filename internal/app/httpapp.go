package app

import (
	"github.com/sojebsikder/go-boilerplate/internal/modules/auth"
	"github.com/sojebsikder/go-boilerplate/internal/modules/user"
	"go.uber.org/fx"
)

func BaseHTTPModules() fx.Option {
	return fx.Options(
		auth.HTTPModule,
		user.HTTPModule,
	)
}
