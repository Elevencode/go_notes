package server

import (
	"go_notes/envs"
	"go_notes/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()
	auth := router.Group("/")
	auth.Use(handlers.AuthMiddleware())
	{
		/// Edit.
		router.PUT("/note/:id", handlers.UpdateNoteHandler)
		/// Delete.
		router.DELETE("/note/:id", handlers.DeleteNoteHandler)
		/// Get.
		router.GET("/note/:id", handlers.GetNoteHandler)
		/// Create.
		router.POST("/note", handlers.CreateNoteHandler)
		/// Get all.
		router.GET("/notes", handlers.GetNotesHandler)
	}

	router.Run(":" + envs.ServerEnvs.NOTES_PORT)
}
