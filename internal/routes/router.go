package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sojebsikder/go-boilerplate/internal/config"
	"github.com/sojebsikder/go-boilerplate/internal/middleware"
	"github.com/sojebsikder/go-boilerplate/internal/modules/metrics"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SetupRouter(lc fx.Lifecycle, ctg *config.Config, r *gin.Engine, log *zap.Logger) {
	// logger
	r.Use(middleware.RequestID())
	r.Use(middleware.RequestLogger(log))

	// prometheus metrics
	metrics.Register()
	r.Use(middleware.Prometheus())
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.GET("/health", func(c *gin.Context) {
		log.Info("health_check")
		c.JSON(200, gin.H{"status": "ok"})
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
