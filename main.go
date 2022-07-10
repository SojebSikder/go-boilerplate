package main

import (
	"github.com/SojebSikder/goframe/app/middleware"
	"github.com/SojebSikder/goframe/config"
	"github.com/SojebSikder/goframe/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the application
	r := gin.Default()

	config.GetConfig()
	r.Static("/static", "./"+config.StaticDir)
	r.LoadHTMLGlob(config.TemplateDir + "/*")

	// Custom middleware call here
	r.Use(middleware.Hello())
	routes.Routes(r)

	r.Run(":8080")
}
