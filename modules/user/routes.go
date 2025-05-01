package user

import (
	"sojebsikder/go-boilerplate/modules/common/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", handler.Create)
		userRoutes.GET("/", handler.GetAll)
	}
}
