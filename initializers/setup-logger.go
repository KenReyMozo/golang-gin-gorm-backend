package initializers

import (
	"fmt"
	"time"

	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupLogger() {

	today := time.Now().Format("2006-01-02_15-04-05")
	logFilePath := fmt.Sprintf("%s/%s.log", LOG_DIR, today)
	f, err := os.Create(logFilePath)
	if err != nil {
		panic(fmt.Sprintf("failed to create log file: %v", err))
	}
	
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}