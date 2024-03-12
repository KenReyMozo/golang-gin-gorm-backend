package main

import (
	"go-backend/controllers"
	"go-backend/initializers"
	"go-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.SetupEnv()
	initializers.ConnectToDB()
	initializers.SetupLogger()
	initializers.SetupOAuth2()
}

func main() {
	r := gin.New()
	r.Use(
		gin.Recovery(),
		middlewares.Logger(),
		middlewares.Cors(),
	)
	
	r.GET("/oauth/login", controllers.OAuthLogin)
	r.GET("/oauth/callback", controllers.OAuthCallback)
	r.POST("/signup", controllers.SignUpUser)
	r.POST("/login", controllers.LoginUser)
	r.GET("/validate", controllers.RequireAuth, controllers.ValidateUser)

	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.PATCH("/posts/:id", controllers.PatchPost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run()
}