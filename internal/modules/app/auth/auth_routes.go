package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB, c *AuthController) {
	routes := r.Group("/auth")

	routes.POST("/register", c.Register)
	routes.POST("/login", c.Login)
}
