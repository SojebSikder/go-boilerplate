package user

import (
	"github.com/sojebsikder/go-boilerplate/model"
	"github.com/sojebsikder/go-boilerplate/pkg/repository"
)

type Service interface {
	CreateUser(model.User) (model.User, error)
	GetAllUsers() ([]model.User, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo}
}

func (s *service) CreateUser(user model.User) (model.User, error) {
	return s.repo.Create(user)
}

func (s *service) GetAllUsers() ([]model.User, error) {
	return s.repo.FindAll()
}
