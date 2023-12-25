package api

import (
	"dataspace/api/general"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	// In case of production, set the Gin mode to release.
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create the Gin engine.
	r := gin.Default()

	r.Use(cors.Default())

	general.SetupRoutes(r)

	r.Run(":8080")
}
