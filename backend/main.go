package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
    // Initialize a Gin router
    router := gin.Default()

    // Define a route that responds with "Hello, world!"
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, world!",
        })
    })

    // Start the server on port 8080
    if err := router.Run(":8080"); err != nil {
        log.Fatal("Error starting server: ", err)
    }
}
