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
	if err := GetModels(ctx, &posts); err != nil { return }

	ctx.JSON(200, posts)
}

func GetPost(ctx *gin.Context) {
	var post model.Post
	if err := GetModelByID(ctx, &post); err != nil { return }

	ctx.JSON(200, post)
}

func UpdatePost(ctx *gin.Context) {

	var post model.Post
	if err := GetModelByID(ctx, &post); err != nil { return }
	if err := BindModel(ctx, &post); err != nil { return }
	tx := StartTransation(ctx)
	if tx == nil { return }
	if err := UpdateModelByID(ctx, tx, &post); err != nil { return }
	ctx.JSON(200, post)
}

func PatchPost(ctx *gin.Context) {

	var post model.Post
	if err := GetModelByID(ctx, &post); err != nil { return }
	tx := StartTransation(ctx)
	if tx == nil { return }
	if err := UpdateModelByID(ctx, tx, &post); err != nil { return }

	ctx.JSON(200, post)
}

func DeletePost(ctx *gin.Context) {

	var post model.Post
	if err := GetModelByID(ctx, &post); err != nil { return }
	tx := StartTransation(ctx)
	if tx == nil { return }
	if err := DeleteModelByID(ctx, tx, &post); err != nil { return }

	ctx.JSON(http.StatusOK, post);
}