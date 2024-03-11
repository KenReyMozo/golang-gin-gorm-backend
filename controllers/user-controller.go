package controllers

import (
	"go-backend/initializers"
	model "go-backend/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(ctx *gin.Context) {

	var user model.User
	if err := BindModel(ctx, &user); err != nil {
		SetResponse(ctx, http.StatusBadRequest)
		return
	}

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

func LoginUser(ctx *gin.Context) {
	var body model.User

	if err := BindModel(ctx, &body); err != nil {
		SetResponse(ctx, http.StatusBadRequest)
		return
	}

	var user model.User
	if err := GetModelBySingleQuery(ctx, "email", body.Email, &user); err != nil {
		SetResponse(ctx, http.StatusBadRequest)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		SetResponse(ctx, http.StatusBadRequest)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"expires": time.Hour * 24,
		"expiration": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		SetResponse(ctx, http.StatusBadRequest)
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600 * 8, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H {})
}

func ValidateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "I'm logged in",
	})
}