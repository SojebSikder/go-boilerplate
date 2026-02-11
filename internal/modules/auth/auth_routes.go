package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *AuthController) {
	routes := r.Group("/api/auth")

	routes.GET("/hello", c.Hello)
	routes.POST("/register", c.Register)
	routes.POST("/login", c.Login)
}
