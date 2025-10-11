package main

import (
	"fmt"
	"os"

	"github.com/sojebsikder/go-boilerplate/internal/config"
	"github.com/sojebsikder/go-boilerplate/internal/model"
	"github.com/sojebsikder/go-boilerplate/internal/modules/app/auth"
	"github.com/sojebsikder/go-boilerplate/internal/modules/app/user"
	"github.com/sojebsikder/go-boilerplate/internal/routes"
	"github.com/sojebsikder/go-boilerplate/pkg/ORM"
	"github.com/sojebsikder/go-boilerplate/pkg/repository"
	"github.com/spf13/cobra"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/fx"

	"log"
)

var verbose bool

var RootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "MyCLI is a demo CLI tool",
	Long:  "MyCLI is a demonstration of a complex CLI using Cobra in Go.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use --help to see available commands.")
	},
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the web server",
	Run: func(cmd *cobra.Command, args []string) {
		StartServer()
	},
}

// Execute runs the CLI
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	RootCmd.AddCommand(serverCmd)
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	Execute()
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
