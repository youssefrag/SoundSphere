package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/youssefrag/SoundSphere/db"
)

func main() {

    db.InitDB()

    server := gin.Default()

    server.Run(":8080")

    fmt.Println("reached here")
}

