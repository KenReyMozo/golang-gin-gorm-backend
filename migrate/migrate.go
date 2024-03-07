package main

import (
	"go-backend/initializers"
	model "go-backend/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&model.Post{})
	initializers.DB.AutoMigrate(&model.User{})
}