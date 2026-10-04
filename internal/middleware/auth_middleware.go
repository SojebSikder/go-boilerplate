package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sojebsikder/go-boilerplate/internal/config"
	"github.com/sojebsikder/go-boilerplate/internal/utils"
)

type AuthMiddleware struct {
	cfg *config.Config
}

func NewAuthMiddleware(cfg *config.Config) *AuthMiddleware {
	return &AuthMiddleware{cfg: cfg}
}

func (m *AuthMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are not authorized. Please login",
				"success": false,
			})
			return
		}

		userID, err := utils.ParseAccessToken(m.cfg.Security.JWTSecret, utils.BearerToken(authHeader))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
				"success": false,
			})
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
