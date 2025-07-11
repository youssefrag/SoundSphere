package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/youssefrag/SoundSphere/cache"
	"github.com/youssefrag/SoundSphere/models"
)

const (
  allSongsKey     = "songs:all"
)

type editSongRequest struct {
	Name  string `json:"name"`
	Genre string `json:"genre"`
}

func saveSongHandler(context *gin.Context) {

	var song models.Song
	
	err := context.ShouldBindJSON(&song)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})

		return
	}

	err = song.Save()

	fmt.Println(err)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save song."})
		return
	}

	cache.Client.Del(cache.Ctx, allSongsKey)
	
	context.JSON(http.StatusCreated, gin.H{"message": "Song saved successfully."})
	
}

func getAllSongsHandler(context *gin.Context) {
	songs, err := models.GetAllSongs()
	
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
	}
	context.JSON(http.StatusOK, songs)
}

func deleteSongsHandler(context * gin.Context) {
	songIdStr := context.Param("songId")

	songId, err := strconv.ParseInt(songIdStr, 10, 64)
  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": "invalid song ID"})
    return
  }

	fmt.Println(songId)

	storagePath, err := models.DeleteSong(songId)

	if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": "could not delete song"})
    return
  }

	cache.Client.Del(cache.Ctx, allSongsKey)

	context.JSON(http.StatusOK, gin.H{
		"message":      "song deleted",
		"storage_path": storagePath,
	})
}

func editSongHandler(context * gin.Context) {
	songIdStr := context.Param("songId")

	songId, err := strconv.ParseInt(songIdStr, 10, 64)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": "invalid song ID"})
    return
  }

	var req editSongRequest
	if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	err = models.EditSong(songId, req.Name, req.Genre)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not edit song"})
	}

	cache.Client.Del(cache.Ctx, allSongsKey)

	context.JSON(http.StatusOK, gin.H{
		"message": "edit succesful",
	})
}

func getSongDetailsHandler(context *gin.Context) {
	songIdStr := context.Param("songId")

	songId, err := strconv.ParseInt(songIdStr, 10, 64)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": "invalid song ID"})
    return
  }

	songDetails, err := models.GetSongDetails(songId)

	if err != nil {
		
		context.JSON(http.StatusBadRequest, gin.H{"error": "could fetch song details"})
	}


	context.JSON(http.StatusOK, songDetails)
}