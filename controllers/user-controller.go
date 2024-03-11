package controllers

import (
	"go-backend/initializers"
	model "go-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(ctx *gin.Context) {

	var user model.User
	if err := BindModel(ctx, &user); err != nil { return }

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		SetResponse(ctx, http.StatusBadRequest)
		return
	}

	newUser := model.User{ Email: user.Email, Password: string(hashedPass)}
	result := initializers.DB.Create(&newUser)

	if result.Error != nil {
		SetResponse(ctx, http.StatusBadRequest)
		return
	}

	SetResponse(ctx, http.StatusOK)
}