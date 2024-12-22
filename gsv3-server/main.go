package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	s := gin.Default()
	s.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := s.Run(":8080")
	if err != nil {
		return
	}
}
