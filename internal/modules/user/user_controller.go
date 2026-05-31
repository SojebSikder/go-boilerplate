package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sojebsikder/go-boilerplate/internal/model"
)

type UserController struct {
	userService *UserService
}

func NewUserController(
	userService *UserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (h *UserController) Create(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.userService.CreateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, created)
}

func (h *UserController) GetAll(ctx *gin.Context) {
	users, err := h.userService.GetAllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
