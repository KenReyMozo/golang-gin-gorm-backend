package initializers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func SetupEnv() {

	mode := os.Getenv("GIN_MODE")
	if mode == "release" {
		mode = gin.ReleaseMode
	} else {
		mode = gin.DebugMode
	}

	// Set Gin mode
	gin.SetMode(mode)

}