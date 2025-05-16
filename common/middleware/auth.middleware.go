package middleware

import (
	"net/http"
	"sojebsikder/go-boilerplate/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		ctg, _ := config.GetConfig()
		JWT_SECRET := []byte(ctg.Security.JWTKey)

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return JWT_SECRET, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"])
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}
	}
}
