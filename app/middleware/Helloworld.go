package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Hello World")
		c.Next()
	}
}
