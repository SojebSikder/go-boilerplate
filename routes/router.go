package routes

import (
	"github.com/sojebsikder/go-boilerplate/internal/app/user"
	"github.com/sojebsikder/go-boilerplate/internal/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) *gin.Engine {

	auth.RegisterRoutes(r, db)
	user.RegisterRoutes(r, db)

	return r
}
