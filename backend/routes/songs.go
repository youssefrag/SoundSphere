package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youssefrag/SoundSphere/cache"
	"github.com/youssefrag/SoundSphere/models"
)

const (
  allSongsKey     = "songs:all"
)

func saveSongHandler(context *gin.Context) {

	var song models.Song

	
	err := context.ShouldBindJSON(&song)

	fmt.Println(song)

	fmt.Println("🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢")

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		fmt.Println("🔴🔴🔴🔴🔴🔴🔴", err)
		return
	}

	fmt.Println("🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣")


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