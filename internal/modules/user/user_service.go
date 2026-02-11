package user

import (
	"github.com/sojebsikder/go-boilerplate/internal/model"
	"github.com/sojebsikder/go-boilerplate/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(user model.User) (model.User, error) {
	return s.repo.Create(user)
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repo.FindAll()
}
