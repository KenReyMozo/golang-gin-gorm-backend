package controllers

import (
	model "go-backend/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(ctx *gin.Context) {

	var user model.User
	if err := BindModel(ctx, &user); err != nil { return }

	bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	
}