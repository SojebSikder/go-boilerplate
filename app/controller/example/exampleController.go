package example

import (
	"net/http"

	"github.com/SojebSikder/goframe/app/model/user"
	orm "github.com/SojebSikder/goframe/system/core/ORM"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"welcome": "World",
	})
}

func Hello(c *gin.Context) {
	user.CreateUser(orm.Ctx, orm.Client)

	var result any
	result, _ = user.QueryUser(orm.Ctx, orm.Client)

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
