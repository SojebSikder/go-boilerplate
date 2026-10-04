package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sojebsikder/go-boilerplate/internal/apperr"
)

func OK(ctx *gin.Context, message string, data any) {
	ctx.JSON(http.StatusOK, gin.H{"message": message, "success": true, "data": data})
}

func Created(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusCreated, gin.H{"message": message, "success": true})
}

func ServerError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "success": false})
}

func BadRequest(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, gin.H{"message": message, "success": false})
}

func BindJSON(ctx *gin.Context, req any) bool {
	if err := ctx.ShouldBindJSON(req); err != nil {
		apperr.Respond(ctx, apperr.BadRequest("Invalid request: "+err.Error()))
		return false
	}
	return true
}

// UUIDParam parses a path param as a uuid
func UUIDParam(ctx *gin.Context, name string) (uuid.UUID, bool) {
	id, err := uuid.Parse(ctx.Param(name))
	if err != nil {
		apperr.Respond(ctx, apperr.BadRequest(name+" must be a uuid"))
		return uuid.Nil, false
	}
	return id, true
}

// UserAndUUIDParam returns the authenticated user plus a uuid path param.
func UserAndUUIDParam(ctx *gin.Context, name string) (uuid.UUID, uuid.UUID, bool) {
	userID, err := GetUserID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token", "success": false})
		return uuid.Nil, uuid.Nil, false
	}

	id, ok := UUIDParam(ctx, name)
	return userID, id, ok
}
