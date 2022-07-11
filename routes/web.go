package routes

import (
	"github.com/SojebSikder/goframe/app/controller"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", controller.Index)
	router.GET("/test", controller.Hello)
}
