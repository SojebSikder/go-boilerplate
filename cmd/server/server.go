package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sojebsikder/go-boilerplate/internal/config"
	"github.com/sojebsikder/go-boilerplate/internal/model"
	"github.com/sojebsikder/go-boilerplate/internal/modules/app/auth"
	"github.com/sojebsikder/go-boilerplate/internal/modules/app/user"
	"github.com/sojebsikder/go-boilerplate/internal/routes"
	"github.com/sojebsikder/go-boilerplate/pkg/ORM"
	"github.com/sojebsikder/go-boilerplate/pkg/repository"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the web server",
	Run: func(cmd *cobra.Command, args []string) {
		StartServer()
	},
}

func StartServer() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fx.New(
		config.Module,
		auth.Module,
		user.Module,
		repository.Module,
		fx.Provide(
			GinServer,
			ORM.Init,
		),
		fx.Invoke(
			AutoMigrate,
			routes.SetupRouter,
		),
	)
	app.Run()
}

func GinServer() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./"+config.StaticDir)
	r.LoadHTMLGlob(config.TemplateDir + "/*")
	return r
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
