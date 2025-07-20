package repository

import (
	"github.com/sojebsikder/go-boilerplate/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(user model.User) (model.User, error)
	FindAll() ([]model.User, error)
	FindByEmail(email string) (model.User, error)
	Update(user model.User) (model.User, error)
	Delete(user model.User) error
	FindByID(id string) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *repository) Update(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *repository) Delete(user model.User) error {
	err := r.db.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) FindByID(id string) (model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *repository) FindByEmailAndPassword(email, password string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
