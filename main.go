package main

import (
	"sojebsikder/go-boilerplate/config"
	"sojebsikder/go-boilerplate/routes"
	orm "sojebsikder/go-boilerplate/system/core/ORM"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctg, _ := config.GetConfig()

	DatabaseURL := ctg.Database.DatabaseURL
	// Initialize ORM
	orm.Init(DatabaseURL)
	// Migrate the schema
	orm.GetDB().AutoMigrate(&Product{})

	// Initialize the application
	r := gin.Default()

	r.Static("/static", "./"+config.StaticDir)
	r.LoadHTMLGlob(config.TemplateDir + "/*")

	// Custom middleware call here
	// Setup the routes
	routes.SetupRouter(r, orm.GetDB())

	r.Run(":" + ctg.App.Port)
}
