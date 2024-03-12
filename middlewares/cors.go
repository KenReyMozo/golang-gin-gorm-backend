package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func Cors() gin.HandlerFunc {

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * int(time.Hour), 
})

	return func(c *gin.Context) {
			corsHandler.HandlerFunc(c.Writer, c.Request)
			c.Next()
	}
}

