package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/users", handleV1Users)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/users", handleV2Users)
	}

	r.Run(":8080")
}

func handleV1Users(c *gin.Context) {
	// Handle v1 user requests
	c.JSON(200, gin.H{"message": "Handling v1 users request"})
}

func handleV2Users(c *gin.Context) {
	// Handle v2 user requests
	c.JSON(200, gin.H{"message": "Handling v2 users request"})
}
