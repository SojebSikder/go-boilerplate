package repository

import (
	"github.com/sojebsikder/go-boilerplate/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Update(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Delete(user model.User) error {
	err := r.db.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByID(id string) (model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmailAndPassword(email, password string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
