package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *AuthController) {
	routes := r.Group("/api/auth")

	routes.POST("/register", c.Register)
	routes.POST("/login", c.Login)
}
