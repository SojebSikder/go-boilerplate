package auth

import (
	"time"

	"github.com/sojebsikder/go-boilerplate/config"
	"github.com/sojebsikder/go-boilerplate/model"
	"github.com/sojebsikder/go-boilerplate/pkg/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(model.User) (model.User, error)
	GetAllUsers() ([]model.User, error)
	Login(string, string) (string, error)
	HashPassword(string) (string, error)
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

func (s *service) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	ctg, _ := config.GetConfig()
	tokenString, err := token.SignedString([]byte(ctg.Security.JWTKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *service) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *service) ComparePassword(hashedPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateUser(user model.User) (model.User, error) {
	return s.repo.Update(user)
}

func (s *service) DeleteUser(id string) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(user)
}
