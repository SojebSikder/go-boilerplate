package controller

import (
	"log"
	"net/http"

	model "github.com/SojebSikder/goframe/app/model"
	orm "github.com/SojebSikder/goframe/system/core/ORM"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	model.CreateUser(orm.Ctx, orm.Client)

	var result, err any
	result, err = model.QueryUser(orm.Ctx, orm.Client)

	log.Println("result: ", result, "err: ", err)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello",
	})
}
