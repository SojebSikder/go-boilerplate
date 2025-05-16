package main

import (
	"sojebsikder/go-boilerplate/config"
	"sojebsikder/go-boilerplate/models"
	"sojebsikder/go-boilerplate/routes"
	ORM "sojebsikder/go-boilerplate/system/core/ORM"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctg, _ := config.GetConfig()

	DatabaseURL := ctg.Database.DatabaseURL
	// Initialize ORM
	ORM.Init(DatabaseURL)
	// Migrate the schema
	ORM.GetDB().AutoMigrate(&models.User{})

	// Initialize the application
	r := gin.Default()

	r.Static("/static", "./"+config.StaticDir)
	r.LoadHTMLGlob(config.TemplateDir + "/*")

	// Custom middleware call here
	// Setup the routes
	routes.SetupRouter(r, ORM.GetDB())

	r.Run(":" + ctg.App.Port)
}
