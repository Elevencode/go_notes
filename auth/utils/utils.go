package utils

import (
	"auth/envs"
	"auth/models"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// TODO(arthur): implement personal info hashing

// Hash password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Check password hash.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Generate access and refresh token with HS256 signing.
func GenerateTokens(userId uint) (models.Tokens, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"user_id": userId, "exp": time.Now().Add(time.Hour * 24).Unix()})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"user_id": userId, "exp": time.Now().Add(time.Hour * 600).Unix()})

	signedAccessToken, _ := accessToken.SignedString([]byte(envs.ServerEnvs.JWT_SECRET))
	signedRefreshToken, _ := refreshToken.SignedString([]byte(envs.ServerEnvs.JWT_SECRET))

	return models.Tokens{AccessToken: signedAccessToken, RefreshToken: signedRefreshToken}, nil
}

// Refresh token validator.
func ValidateRefreshToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid sign method: %v", token.Header["alg"])
		}
		return []byte(envs.ServerEnvs.JWT_SECRET), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDValue, ok := claims["user_id"].(float64)

		if !ok {
			return 0, fmt.Errorf("user_id claim is not a float64")
		}
		return uint(userIDValue), nil
	} else {
		return 0, fmt.Errorf("invalid token")
	}
}

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
