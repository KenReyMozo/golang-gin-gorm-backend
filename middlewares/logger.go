package middlewares

import (
	"fmt"
	"time"

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
