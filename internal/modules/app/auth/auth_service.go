package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sojebsikder/go-boilerplate/internal/config"
	"github.com/sojebsikder/go-boilerplate/internal/model"
	"github.com/sojebsikder/go-boilerplate/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user model.User) (model.User, error) {
	return s.repo.Create(user)
}

func (s *AuthService) GetAllUsers() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *AuthService) Login(email, password string) (string, error) {
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

func (s *AuthService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *AuthService) ComparePassword(hashedPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}

func (s *AuthService) UpdateUser(user model.User) (model.User, error) {
	return s.repo.Update(user)
}

func (s *AuthService) DeleteUser(id string) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(user)
}
