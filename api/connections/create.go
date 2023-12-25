package connections

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleCreate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}