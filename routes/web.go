package routes

import (
	"github.com/SojebSikder/goframe/app/controller/example"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", example.Index)
	router.GET("/test", example.Hello)
}
