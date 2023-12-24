package general

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	generalGroup := router.Group("/")
	{
		generalGroup.GET("/ping", pingHandler)
	}
}
