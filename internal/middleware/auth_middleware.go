package middleware

import (
	"net/http"
	"strings"

	"github.com/sojebsikder/go-boilerplate/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are not authorized. Please login",
				"success": false,
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		tokenString = strings.TrimSpace(tokenString)

		ctg, _ := config.NewConfig()
		JWT_SECRET := []byte(ctg.Security.JWTSecret)

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			return JWT_SECRET, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"])
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
				"success": false,
			})
		}
	}
}
