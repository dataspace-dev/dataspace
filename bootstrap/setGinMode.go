package bootstrap

import (
	"os"

	"github.com/gin-gonic/gin"
)

func SetGitMode() {
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
}