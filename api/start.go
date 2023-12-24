package api

import (
	"dataspace/api/general"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	r.Use(cors.Default())

	general.SetupRoutes(r)

	r.Run(":8080") 
}