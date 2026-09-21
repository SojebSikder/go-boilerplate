package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hibiken/asynq"
	"github.com/sojebsikder/go-boilerplate/internal/config"
	"github.com/sojebsikder/go-boilerplate/internal/model"
	authtask "github.com/sojebsikder/go-boilerplate/internal/modules/auth/task"
	"github.com/sojebsikder/go-boilerplate/internal/repository"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	logger      *zap.Logger
	userRepo    *repository.UserRepository
	config      *config.Config
	asynqClient *asynq.Client
}

func NewAuthService(
	logger *zap.Logger,
	userRepo *repository.UserRepository,
	config *config.Config,
	asynqClient *asynq.Client,
) *AuthService {
	return &AuthService{
		logger:      logger,
		userRepo:    userRepo,
		config:      config,
		asynqClient: asynqClient,
	}
}

func (s *AuthService) Hello(ctx context.Context) (string, error) {
	// add task to queue
	task, err := authtask.NewAuthTask("Sojeb")
	if err != nil {
		s.logger.Error("failed to create task", zap.Error(err))
		return "", errors.New("failed to create task")
	}

	_, err = s.asynqClient.Enqueue(task,
		asynq.Queue("critical"),
		// asynq.MaxRetry(1),
		asynq.Timeout(30*time.Second),
	)
	if err != nil {
		s.logger.Error("failed to enqueue task", zap.Error(err))
		return "", errors.New("failed to enqueue task")
	}
	return fmt.Sprintf("Hello, %s!", "World"), nil
}

func (s *AuthService) CreateUser(ctx context.Context, req *AuthRegisterRequest) (model.User, error) {
	// Check if user already exists
	if _, err := s.userRepo.FindByEmail(ctx, req.Email); err == nil {
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
	return s.userRepo.Create(ctx, user)
}

func (s *AuthService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.userRepo.FindAll(ctx)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
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

func (s *AuthService) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	return s.userRepo.Update(ctx, user)
}

func (s *AuthService) DeleteUser(ctx context.Context, id string) error {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return s.userRepo.Delete(ctx, user)
}
