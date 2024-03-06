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
	f, _ := os.Create(fmt.Sprintf("[%s].log", today))
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}