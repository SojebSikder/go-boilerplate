package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sojebsikder/go-boilerplate/internal/app"
	"github.com/sojebsikder/go-boilerplate/internal/config"
	"github.com/sojebsikder/go-boilerplate/internal/middleware"
	"github.com/sojebsikder/go-boilerplate/internal/modules/auth"
	"github.com/sojebsikder/go-boilerplate/internal/modules/user"
	"github.com/sojebsikder/go-boilerplate/internal/routes"

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
		app.BaseModules(),
		auth.HTTPModule,
		user.HTTPModule,

		fx.Provide(
			GinServer,
			// utils.NewRateLimiter, // uncomment this if you want to enable rate limiting
		),
		fx.Invoke(
			routes.SetupRouter,
		),
	)
	app.Run()
}

func GinServer(cfg *config.Config,

// rateLimiter *utils.RateLimiter, // uncomment this if you want to enable rate limiting
) *gin.Engine {
	r := gin.Default()

	// Apply global middleware
	r.Use(middleware.CorsMiddleware())
	// r.Use(rateLimiter.Limit()) // uncomment this if you want to enable rate limiting

	// Serve static files and templates
	r.Static("/static", "./"+cfg.App.StaticDir)
	r.LoadHTMLGlob(cfg.App.TemplateDir + "/*")

	return r
}
