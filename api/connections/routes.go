package connections

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	connectionsGroup := router.Group("/")
	{
		connectionsGroup.POST("/connections", handleCreate)
	}
}
