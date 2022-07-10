package main

import (
	"github.com/SojebSikder/goframe/app/middleware"
	"github.com/SojebSikder/goframe/config"
	"github.com/SojebSikder/goframe/routes"
	"github.com/gin-gonic/gin"

	"context"
	"log"

	"github.com/SojebSikder/goframe/ent"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Initialize database connection
	client, err := ent.Open("mysql", "root@tcp(localhost:3306)/go-example?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

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
