package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sojebsikder/go-boilerplate/config"
	"github.com/sojebsikder/go-boilerplate/internal/app/user"
	"github.com/sojebsikder/go-boilerplate/internal/auth"
	"go.uber.org/fx"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(lc fx.Lifecycle, ctg config.Config, r *gin.Engine, db *gorm.DB) {

	auth.RegisterRoutes(r, db)
	user.RegisterRoutes(r, db)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Test route is working!"})
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := r.Run(":" + ctg.App.Port); err != nil && err != http.ErrServerClosed {
					fmt.Println("Failed to start server:", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
