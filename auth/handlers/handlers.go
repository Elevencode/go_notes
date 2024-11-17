package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignInHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Auth success"})
}

func RefreshTokenHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Token refreshed"})
}

func GetUserHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Get user success"})
}

func RegisterUserHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Register success"})
}
