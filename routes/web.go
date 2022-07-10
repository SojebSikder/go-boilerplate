package routes

import (
	"net/http"

	controller "github.com/SojebSikder/goframe/app/controller"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/test", controller.Index)
}
