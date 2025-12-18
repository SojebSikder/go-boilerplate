package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserID(ctx *gin.Context) (uuid.UUID, error) {
	userIDStr := ctx.MustGet("user_id").(string)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user_id in context: %v", err)
	}
	return userID, nil
}
