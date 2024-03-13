package middlewares

import (
	"fmt"
	"time"

	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s [%d]- [%s] %s %s %s \n",
			params.ClientIP,
			params.StatusCode,
			params.TimeStamp.Format(time.RFC822),
			params.Method,
			params.Path,
			params.Latency,
		)
	})
}

func LoggerPerDay() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		start := time.Now()

		logFileName := fmt.Sprintf("%s/%s.log", "logs", time.Now().Format("2006-01-02"))

		file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		defer file.Close()

		gin.DefaultWriter = file

		ctx.Next()

		end := time.Now()

		latency := end.Sub(start)

		fmt.Fprintf(file, "[%s] [%s] %d %s %s %s\n",
			time.Now().Format("2006-01-02 15:04:05"),
			ctx.ClientIP(),
			ctx.Writer.Status(),
			ctx.Request.Method,
			ctx.Request.URL.Path,
			latency,
		)
	}
}