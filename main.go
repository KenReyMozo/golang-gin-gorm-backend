package main

import (
	"go-backend/controllers"
	"go-backend/initializers"
	"go-backend/middlewares"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.SetupDirs()
	initializers.LoadEnv()
	initializers.SetupEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.New()

	r.Use(
		gin.Recovery(),
		middlewares.LoggerPerDay(),
		middlewares.Cors(),
		middlewares.CustomRateLimiter(3, time.Second/2),
	)

	r.POST("/signup", controllers.SignUpUser)

	r.POST("/login", controllers.LoginUser)
	r.POST("/login-secure", controllers.EncLoginUser)

	r.GET("/validate", controllers.RequireAuth, controllers.ValidateUser)
	r.GET("/me", controllers.RequireAuth, controllers.GetMe)

	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.PATCH("/posts/:id", controllers.PatchPost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run()
}
