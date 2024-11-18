package handlers

import (
	"fmt"
	"go_notes/envs"
	"strings"

	"github.com/golang-jwt/jwt"
)

// Get user id from token.
func ExtractUserId(tokenString string) (uint, error) {
	str := strings.TrimSpace(strings.TrimPrefix(tokenString, "Bearer"))

	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid sign method: %v", token.Header["alg"])
		}
		return []byte(envs.ServerEnvs.JWT_SECRET), nil
	})

	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["user_id"]

		if userIdFloat, ok := userId.(float64); ok {
			return uint(userIdFloat), nil
		}
	}

	return 0, fmt.Errorf("could not extract user_id from token")
}