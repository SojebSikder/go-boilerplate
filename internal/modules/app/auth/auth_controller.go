package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *AuthService
}

func NewAuthController(
	authService *AuthService,
) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (h *AuthController) Register(ctx *gin.Context) {
	var req AuthRegisterRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	created, err := h.authService.CreateUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (h *AuthController) Login(ctx *gin.Context) {
	var req AuthRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, loginErr := h.authService.Login(ctx, req.Email, req.Password)

	if loginErr != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
