package repository

import (
	"sojebsikder/go-boilerplate/models"

	"gorm.io/gorm"
)

type Repository interface {
	Create(user models.User) (models.User, error)
	FindAll() ([]models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}
