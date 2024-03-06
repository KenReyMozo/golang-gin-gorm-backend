package main

import (
	"go-backend/controllers"
	"go-backend/initializers"
	"go-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.SetupLogger()
}

func main() {
	r := gin.New()
	r.Use(
		gin.Recovery(),
		middlewares.Logger(),
	)
	r.POST("/posts", controllers.CreatePost)
	r.Run()
}