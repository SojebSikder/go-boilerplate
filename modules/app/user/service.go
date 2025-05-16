package user

import (
	"sojebsikder/go-boilerplate/common/repository"
	"sojebsikder/go-boilerplate/models"
)

type Service interface {
	CreateUser(models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo}
}

func (s *service) CreateUser(user models.User) (models.User, error) {
	return s.repo.Create(user)
}

func (s *service) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}
