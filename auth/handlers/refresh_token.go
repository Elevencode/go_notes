package handlers

import (
	"auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func refreshToken(ctx *gin.Context) {
	var token RefreshTokenRequest

	if err := ctx.ShouldBindBodyWithJSON(&token); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Request error"})
		return
	}

	userId, err := utils.ValidateRefreshToken(token.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token" + err.Error()})
		return
	}

	tokens, err := utils.GenerateTokens(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Create tokens error"})
		return
	}

	ctx.JSON(http.StatusOK, tokens)
}
