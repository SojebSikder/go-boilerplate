package routes

import (
	"sojebsikder/go-boilerplate/modules/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) *gin.Engine {

	user.RegisterRoutes(r, db)

	return r
}
