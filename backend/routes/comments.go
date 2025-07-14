package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/youssefrag/SoundSphere/models"
)

type addCommentInput struct {
  Email   string `json:"email" binding:"required,email"`
  SongID  int64  `json:"songId,string" binding:"required"`
  Content string `json:"content" binding:"required"`
}

func addNewCommentHandler(context *gin.Context) {

	
	var input addCommentInput
	if err := context.ShouldBindJSON(&input); err != nil {
		fmt.Println("🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑")
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	userID, err := models.GetUserIDByEmail(input.Email)
	if err != nil {
			if err == models.ErrUserNotFound {
					context.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
			} else {
					context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
	}

	commentID, err := models.CreateComment(userID, input.SongID, input.Content)
	if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}

	created := models.Comment{
			ID:      commentID,
	}
	context.JSON(http.StatusCreated, created)
}

func getCommentsHandler(context *gin.Context) {
	songIDStr := context.Param("songId")
	songID, err := strconv.ParseInt(songIDStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid songId"})
		return
	}

	comments, err := models.GetCommentsBySongID(songID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, comments)
}

func deleteCommentHandler(context *gin.Context) {
	commentIdStr := context.Param("commentId")

	commentId, err := strconv.ParseInt(commentIdStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid songId"})
		return
	}

	err = models.DeleteCommentById(commentId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Comment Deleted"})

}