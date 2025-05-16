package routes

import (
	"sojebsikder/go-boilerplate/modules/app/user"
	"sojebsikder/go-boilerplate/modules/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) *gin.Engine {

	auth.RegisterRoutes(r, db)
	user.RegisterRoutes(r, db)

	return r
}
