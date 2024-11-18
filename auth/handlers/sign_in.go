package handlers

import (
	"auth/database"
	"auth/models"
	"auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signIn(ctx *gin.Context) {
	var registerData models.RegisterData

	if err := ctx.ShouldBindBodyWithJSON(&registerData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	var user models.User
	result := database.DB.Where("email = ?", registerData.Email).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	if !utils.CheckPasswordHash(registerData.Password, user.Hash) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	tokens, err := utils.GenerateTokens(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create tokens"})
		return
	}
	userResponse := struct {
		Id    uint   `json:"id"`
		Email string `json:"email"`
	}{
		Id:    user.ID,
		Email: user.Email,
	}

	ctx.JSON(http.StatusOK, gin.H{"tokens": tokens, "user": userResponse})
}
