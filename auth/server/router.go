package server

import (
	"auth/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()
	// Register.
	router.PUT("/user", handlers.RegisterUserHandler)
	// Auth.
	router.POST("/user", handlers.SignInHandler)
	// Token refresh.
	router.POST("/refresh", handlers.RefreshTokenHandler)
	// Get user.
	router.GET("/user", handlers.GetUserHandler)

	router.Run(":9204")
}
