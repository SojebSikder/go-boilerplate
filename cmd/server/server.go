package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sojebsikder/go-boilerplate/internal/config"
	"github.com/sojebsikder/go-boilerplate/internal/middleware"
	"github.com/sojebsikder/go-boilerplate/internal/modules/app/auth"
	"github.com/sojebsikder/go-boilerplate/internal/modules/app/user"
	"github.com/sojebsikder/go-boilerplate/internal/repository"
	"github.com/sojebsikder/go-boilerplate/internal/routes"
	"github.com/sojebsikder/go-boilerplate/pkg/ORM"
	utils "github.com/sojebsikder/go-boilerplate/pkg/ratelimitter"
	"github.com/sojebsikder/go-boilerplate/pkg/redis"
	s3Client "github.com/sojebsikder/go-boilerplate/pkg/s3client"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the web server",
	Run: func(cmd *cobra.Command, args []string) {
		StartServer()
	},
}

func StartServer() {
	app := fx.New(
		config.Module,
		redis.Module,
		s3Client.Module,
		auth.Module,
		user.Module,
		repository.Module,
		fx.Provide(
			GinServer,
			ORM.Init,
			utils.NewRateLimiter,
		),
		fx.Invoke(
			routes.SetupRouter,
		),
	)
	app.Run()
}

func GinServer(cfg *config.Config, rateLimiter *utils.RateLimiter) *gin.Engine {
	r := gin.Default()

	// Apply global middleware
	r.Use(middleware.CorsMiddleware())
	r.Use(rateLimiter.Limit())

	// Serve static files and templates
	r.Static("/static", "./"+cfg.App.StaticDir)
	r.LoadHTMLGlob(cfg.App.TemplateDir + "/*")

	return r
}
