package apperr

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Error struct {
	Data    any
	Message string
	Status  int
}

func (e *Error) Error() string { return e.Message }

func New(status int, message string) *Error {
	return &Error{Status: status, Message: message}
}

func BadRequest(message string) *Error {
	return New(http.StatusBadRequest, message)
}
func Forbidden(message string) *Error {
	return New(http.StatusForbidden, message)
}
func NotFound(message string) *Error {
	return New(http.StatusNotFound, message)
}
func Conflict(message string, current any) *Error {
	return &Error{Status: http.StatusConflict, Message: message, Data: current}
}

// Respond writes err as a JSON error response.
func Respond(ctx *gin.Context, err error) {
	var appErr *Error

	switch {
	case errors.As(err, &appErr):
	case errors.Is(err, gorm.ErrRecordNotFound):
		appErr = NotFound("Not found")
	default:
		appErr = New(http.StatusInternalServerError, "Internal server error")
		_ = ctx.Error(err)
	}

	body := gin.H{"message": appErr.Message, "success": false}
	if appErr.Data != nil {
		body["data"] = appErr.Data
	}
	ctx.AbortWithStatusJSON(appErr.Status, body)
}
