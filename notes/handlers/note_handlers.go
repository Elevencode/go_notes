package handlers

import (
	"encoding/json"
	"fmt"
	"go_notes/database"
	"go_notes/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func GetNoteHandler(ctx *gin.Context) {
	authorId, err := ExtractUserId(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No access"})
		return
	}

	id := ctx.Param("id")
	var note models.Note

	collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", authorId))
	filter := bson.M{"id": id}

	errFind := collection.FindOne(ctx, filter).Decode(&note)
	if errFind != nil {
		ctx.JSON(http.StatusOK, "Note not found")
	}

	ctx.JSON(http.StatusOK, &note)

}

func GetNotesHandler(ctx *gin.Context) {
	authorId, err := ExtractUserId(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No access"})
		return
	}

	var notes []models.Note
	val, err := database.RedisClient.Get(fmt.Sprintf("notes/%d", authorId)).Result()
	if err == redis.Nil {
		log.Printf("Chache not found. Loading from DB...")
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
			ctx.JSON(http.StatusOK, "Notes not found.")
		} else {
			recordToCache(notes, authorId)
			ctx.JSON(http.StatusOK, notes)
		}
	} else {
		log.Printf("Cache found. Loading from cache...")
		getFromCache(val, ctx)
	}

}

func DeleteNoteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	authorId, err := ExtractUserId(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No access"})
		return
	}

	collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", authorId))
	filter := bson.M{"id": id}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
	}

	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusOK, "Note not found")
	} else {
		resetCache(fmt.Sprintf("notes/%d", authorId))
		ctx.JSON(http.StatusOK, "Note delete success")
	}
}

func UpdateNoteHandler(ctx *gin.Context) {
	authorId, err := ExtractUserId(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No access"})
		return
	}

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
		resetCache(fmt.Sprintf("notes/%d", note.AuthorID))
		ctx.JSON(http.StatusOK, "Note update success")
	}

}

func CreateNoteHandler(ctx *gin.Context) {
	authorId, err := ExtractUserId(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No access"})
		return
	}

	var note models.Note

	if err := ctx.ShouldBindBodyWithJSON(&note); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	note.Id = uuid.New().String()
	note.AuthorID = authorId
	collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", note.AuthorID))

	_, errInsert := collection.InsertOne(ctx, note)
	if errInsert != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": errInsert.Error()})
	}

	resetCache(fmt.Sprintf("notes/%d", note.AuthorID))

	ctx.JSON(http.StatusOK, gin.H{
		"note":    note,
		"message": "Note create success"})

}

func getFromCache(val string, ctx *gin.Context) {
	notes := make([]models.Note, 0)
	json.Unmarshal([]byte(val), &notes)
	ctx.JSON(http.StatusOK, notes)
}

func recordToCache(notes []models.Note, authorId uint) {
	notesJSON, err := json.Marshal(notes)
	if err != nil {
		log.Printf("Notes serialize error: %v", err)
	} else {
		err := database.RedisClient.Set(fmt.Sprintf("notes/%d", authorId), string(notesJSON), 1440*time.Minute).Err()
		if err != nil {
			log.Printf("Record to cache error: %v", err)

		}
	}
}

func resetCache(val string) {
	database.RedisClient.Del(val)
}
