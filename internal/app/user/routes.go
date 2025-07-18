package user

import (
	"github.com/sojebsikder/go-boilerplate/pkg/middleware"
	"github.com/sojebsikder/go-boilerplate/pkg/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	repo := repository.NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	routes := r.Group("/users")
	routes.Use(middleware.AuthMiddleware())
	routes.POST("/", handler.Create)
	routes.GET("/", handler.GetAll)

}
