package auth

import (
	"github.com/sojebsikder/go-boilerplate/pkg/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	routes := r.Group("/auth")

	routes.POST("/register", handler.Register)
	routes.POST("/login", handler.Login)

}
