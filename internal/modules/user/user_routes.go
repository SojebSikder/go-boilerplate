package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sojebsikder/go-boilerplate/internal/middleware"
)

func RegisterRoutes(r *gin.Engine, c *UserController, auth *middleware.AuthMiddleware) {
	routes := r.Group("/api/users")
	routes.Use(auth.Handle())
	routes.POST("/", c.Create)
	routes.GET("/", c.GetAll)
}
