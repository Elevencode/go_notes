package server

import "github.com/gin-gonic/gin"

func InitRoutes() {
	router := gin.Default()
	/// Edit.
	router.PUT("/note")
	/// Delete.
	router.DELETE("/note/:id")
	/// Get.
	router.GET("/note/:id")
	/// Create.
	router.POST("/note/:id")
	/// Get all.
	router.GET("/notes")

}
