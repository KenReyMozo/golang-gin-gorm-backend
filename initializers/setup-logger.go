package initializers

import (
	"fmt"
	"time"

	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupLogger() {
	logsDir := "logs"
	err := os.MkdirAll(logsDir, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("failed to create logs directory: %v", err))
	}

	today := time.Now().Format("2006-01-02_15-04-05")
	logFilePath := fmt.Sprintf("%s/%s.log", logsDir, today)
	f, err := os.Create(logFilePath)
	if err != nil {
		panic(fmt.Sprintf("failed to create log file: %v", err))
	}
	
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}