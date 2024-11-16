package handlers

import (
	"fmt"
	"go_notes/database"
	"go_notes/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func GetNoteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var note models.Note

	collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", 1))
	filter := bson.M{"id": id}

	errFind := collection.FindOne(ctx, filter).Decode(&note)
	if errFind != nil {
		ctx.JSON(http.StatusOK, "Note not found")
	}

	ctx.JSON(http.StatusOK, &note)

}

func GetNotesHandler(ctx *gin.Context) {
	authorId := 1
	var notes []models.Note
	collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", authorId))

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var note models.Note
		err := cursor.Decode(&note)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		notes = append(notes, note)
	}

	if err := cursor.Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(notes) == 0 {
		ctx.JSON(http.StatusOK, "Заметок не найдено")
	} else {
		ctx.JSON(http.StatusOK, notes)
	}

}

func DeleteNoteHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", 1))
	filter := bson.M{"id": id}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
	}

	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusOK, "Note not found")
	} else {
		ctx.JSON(http.StatusOK, "Note delete success")
	}
}

func UpdateNoteHandler(ctx *gin.Context) {
	authorId := 1
	id := ctx.Param("id")

	var note models.Note
	if err := ctx.ShouldBindJSON(&note); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", authorId))
	updateFields := bson.M{}

	if note.Name != nil {
		updateFields["name"] = note.Name
	}

	if note.Content != nil {
		updateFields["content"] = note.Content
	}

	update := bson.M{"$set": updateFields}
	filter := bson.M{"id": id}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.MatchedCount == 0 {
		ctx.JSON(http.StatusOK, "Note not found")
	} else {
		ctx.JSON(http.StatusOK, "Note update success")
	}

}

func CreateNoteHandler(ctx *gin.Context) {
	var note models.Note

	if err := ctx.ShouldBindBodyWithJSON(&note); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	note.Id = uuid.New().String()
	note.AuthorID = 1
	collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", note.AuthorID))

	_, errInsert := collection.InsertOne(ctx, note)
	if errInsert != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": errInsert.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"note":    note,
		"message": "Note create success"})

}
