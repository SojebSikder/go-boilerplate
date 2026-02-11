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

func (h *AuthController) Hello(ctx *gin.Context) {
	result, err := h.authService.Hello(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data fetched successful",
		"success": true,
		"data":    result,
	})
}

func (h *AuthController) Register(ctx *gin.Context) {
	var req AuthRegisterRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"success": false,
		})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	_, err := h.authService.CreateUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Account created successfully",
		"success": true,
	})
}

func (h *AuthController) Login(ctx *gin.Context) {
	var req AuthRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"success": false,
		})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	token, loginErr := h.authService.Login(ctx, req.Email, req.Password)
	if loginErr != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
			"success": false,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"success": true,
		"token":   token,
	})
}
