package app

import (
	asynqfx "github.com/sojebsikder/go-boilerplate/internal/asynq"
	"github.com/sojebsikder/go-boilerplate/internal/config"
	"github.com/sojebsikder/go-boilerplate/internal/modules/auth"
	"github.com/sojebsikder/go-boilerplate/internal/modules/user"
	"github.com/sojebsikder/go-boilerplate/internal/repository"
	"github.com/sojebsikder/go-boilerplate/pkg/ORM"
	"github.com/sojebsikder/go-boilerplate/pkg/logger"
	"github.com/sojebsikder/go-boilerplate/pkg/redis"
	"github.com/sojebsikder/go-boilerplate/pkg/s3client"
	"go.uber.org/fx"
)

func BaseModules() fx.Option {
	return fx.Options(
		// infra
		config.Module,
		logger.Module,
		redis.Module,
		// ORM.Init,
		s3client.Module,

		// async / integrations
		asynqfx.Module,

		// data + domain
		repository.Module,

		// business modules
		auth.Module,
		user.Module,

		fx.Provide(
			ORM.Init,
		),
	)
}
