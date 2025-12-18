package auth

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sojebsikder/go-boilerplate/internal/config"
	"github.com/sojebsikder/go-boilerplate/internal/model"
	"github.com/sojebsikder/go-boilerplate/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
	config   *config.Config
}

func NewAuthService(
	userRepo *repository.UserRepository,
	config *config.Config,
) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		config:   config,
	}
}

func (s *AuthService) CreateUser(ctx *gin.Context, req *AuthRegisterRequest) (model.User, error) {
	// Check if user already exists
	if _, err := s.userRepo.FindByEmail(req.Email); err == nil {
		return model.User{}, errors.New("User with this email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return model.User{}, errors.New("Failed to hash password")
	}

	hashedPasswordStr := string(hashedPassword)
	user := model.User{
		Name:     &req.Name,
		Email:    &req.Email,
		Password: &hashedPasswordStr,
	}
	return s.userRepo.Create(user)
}

func (s *AuthService) GetAllUsers() ([]model.User, error) {
	return s.userRepo.FindAll()
}

func (s *AuthService) Login(ctx *gin.Context, email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	tokenString, err := token.SignedString([]byte(s.config.Security.JWTSecret))
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

func (s *AuthService) UpdateUser(ctx *gin.Context, user model.User) (model.User, error) {
	return s.userRepo.Update(user)
}

func (s *AuthService) DeleteUser(ctx *gin.Context, id string) error {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return err
	}
	return s.userRepo.Delete(user)
}
