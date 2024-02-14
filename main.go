package main

import "github.com/gin-gonic/gin"

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Define a route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
