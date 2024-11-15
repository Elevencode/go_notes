package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "GetNoteHandler")
}

func GetNotesHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "GetNotesHandler")
}

func DeleteNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "DeleteNoteHandler")
}

func UpdateNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "UpdateNoteHandler")
}

func CreateNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "CreateNoteHandler")
}
