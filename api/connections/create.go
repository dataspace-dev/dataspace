package connections

import "github.com/gin-gonic/gin"

func handleCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}