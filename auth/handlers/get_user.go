package handlers

import (
	"auth/database"
	"auth/models"
	"auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUser(ctx *gin.Context) {
	userId, err := utils.ExtractUserId(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
		return
	}

	var user models.User
	result := database.DB.Where("ID = ?", userId).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	userResponse := struct {
		Id    uint   `json:"id"`
		Email string `json:"email"`
	}{
		Id:    user.ID,
		Email: user.Email,
	}

	ctx.JSON(http.StatusOK, gin.H{"user": userResponse})
}
