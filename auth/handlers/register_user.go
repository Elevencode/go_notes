package handlers

import (
	"auth/database"
	"auth/models"
	"auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerUserHandler(ctx *gin.Context) {
	var user models.User
	var data models.RegisterData

	if err := ctx.ShouldBindBodyWithJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	hashedPass, err := utils.HashPassword(data.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Hashing password error"})
		return
	}
	user.Email = data.Email
	user.Hash = hashedPass
	result := database.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Create user error"})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "User create success"})
	}
}
