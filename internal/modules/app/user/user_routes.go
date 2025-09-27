package user

import (
	"github.com/sojebsikder/go-boilerplate/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB, c *UserController) {
	routes := r.Group("/users")
	routes.Use(middleware.AuthMiddleware())
	routes.POST("/", c.Create)
	routes.GET("/", c.GetAll)
}
