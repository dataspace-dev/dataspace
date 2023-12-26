package connections

import (
	"dataspace/api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	connectionsGroup := router.Group("/")
	{
		connectionsGroup.POST("/connections", middlewares.AuthMiddleware(), handleCreate)
		connectionsGroup.GET("/connections", middlewares.AuthMiddleware(), handleList)
	}
}
