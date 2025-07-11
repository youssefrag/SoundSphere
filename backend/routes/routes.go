package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)
	server.POST("/refresh", refresh)
	server.POST("/logout", logout)

	server.POST("/saveSong", saveSongHandler)
	server.GET("/allSongs", getAllSongsHandler)
	server.DELETE("songs/:songId", deleteSongsHandler)
	server.PUT("songs/:songId", editSongHandler)
	server.GET("/song-details/:songId", getSongDetailsHandler)

	server.POST("/comments/addNew", addNewCommentHandler)
	server.GET("/comments/:songId", getCommentsHandler)
}