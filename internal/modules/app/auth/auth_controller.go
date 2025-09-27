package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sojebsikder/go-boilerplate/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	service *AuthService
}

func NewAuthController(service *AuthService) *AuthController {
	return &AuthController{service}
}

func (h *AuthController) Register(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := model.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hashedPassword),
	}

	created, err := h.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *AuthController) Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, loginErr := h.service.Login(body.Email, body.Password)

	if loginErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
