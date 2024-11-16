package server

import (
	"go_notes/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()
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

	router.Run(":9200")
}
