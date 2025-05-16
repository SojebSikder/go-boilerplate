package repository

import (
	"sojebsikder/go-boilerplate/models"

	"gorm.io/gorm"
)

type Repository interface {
	Create(user models.User) (models.User, error)
	FindAll() ([]models.User, error)
	FindByEmail(email string) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(user models.User) error
	FindByID(id string) (models.User, error)
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

func (r *repository) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *repository) Update(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *repository) Delete(user models.User) error {
	err := r.db.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) FindByID(id string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *repository) FindByEmailAndPassword(email, password string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
