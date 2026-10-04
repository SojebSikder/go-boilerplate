package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

// BearerToken strips the "Bearer " prefix from an Authorization header value.
func BearerToken(header string) string {
	return strings.TrimSpace(strings.TrimPrefix(header, "Bearer "))
}

// ParseAccessToken validates an HS256 access token and returns its user_id.
func ParseAccessToken(secret, tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	// refresh tokens are opaque, but reject anything explicitly typed otherwise
	if typ, ok := claims["typ"].(string); ok && typ != "access" {
		return "", errors.New("not an access token")
	}
	userID, ok := claims["user_id"].(string)
	if !ok || userID == "" {
		return "", errors.New("token has no user_id")
	}
	return userID, nil
}
