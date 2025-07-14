package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefrag/SoundSphere/utils"
)

func RegisterRoutes(server *gin.Engine) {

	server.POST("/signup", signup)
	server.POST("/login", login)
	server.POST("/refresh", refresh)
	server.POST("/logout", logout)

	server.POST("/saveSong", saveSongHandler)
	server.GET("/allSongs", getAllSongsHandler)
	server.DELETE("/songs/:songId", utils.AuthMiddleware(), deleteSongsHandler)
	server.PUT("/songs/:songId", utils.AuthMiddleware(), editSongHandler)
	server.GET("/song-details/:songId", getSongDetailsHandler)

	server.POST("/comments/addNew", utils.AuthMiddleware(), addNewCommentHandler)
	server.GET("/comments/:songId", getCommentsHandler)
	server.DELETE("/comments/:commentId", utils.AuthMiddleware(), deleteCommentHandler)
}