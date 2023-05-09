package src

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	// Set up logger
	log.New(os.Stdout, "gin: ", log.LstdFlags|log.Lshortfile)

	// Creates a gin router with default middleware
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve
}
