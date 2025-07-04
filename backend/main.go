package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/youssefrag/SoundSphere/db"
	"github.com/youssefrag/SoundSphere/routes"
)

func main() {

		if err := godotenv.Load(); err != nil {
    	log.Println("No .env file found — relying on real environment")
  	}

    db.InitDB()

    server := gin.Default()

    	// Configure CORS
		server.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:5173"},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))

    routes.RegisterRoutes(server)

    server.Run(":8080")

}

