package main

import (
	"github.com/sojebsikder/go-boilerplate/config"
	"github.com/sojebsikder/go-boilerplate/model"
	"github.com/sojebsikder/go-boilerplate/pkg/ORM"
	"github.com/sojebsikder/go-boilerplate/routes"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/fx"

	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GinServer() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./"+config.StaticDir)
	r.LoadHTMLGlob(config.TemplateDir + "/*")
	return r
}

func AutoMigrate(db *gorm.DB) {
	ORM.GetDB().AutoMigrate(&model.User{})
}

func main() {
	app := fx.New(
		config.Module,
		fx.Provide(
			GinServer,
			ORM.Init,
		),
		fx.Invoke(
			routes.SetupRouter,
			AutoMigrate,
		),
	)
	app.Run()

}
