package controllers

import (
	"go-backend/initializers"
	model "go-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {

	var created model.Post
	err := ctx.ShouldBindJSON(&created)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	post := model.Post{Title: created.Body, Body: created.Title}

	result := initializers.DB.Create(&post)
	
	if result.Error != nil {
		ctx.Status(http.StatusBadRequest)
		return 
	}

	ctx.JSON(200, post)
}

func GetPosts(ctx *gin.Context) {
	var posts []model.Post
	initializers.DB.Find(&posts)

	ctx.JSON(200, posts)
}

func GetPost(ctx *gin.Context) {
	id := ctx.Param("id")

	var post model.Post
	initializers.DB.First(&post, id)

	ctx.JSON(200, post)
}

func UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")

	var post model.Post
	if initializers.DB.First(&post, id).Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	err := ctx.ShouldBindJSON(&post)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	initializers.DB.Save(&post)

	ctx.JSON(200, post)
}