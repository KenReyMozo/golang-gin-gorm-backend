package controllers

import (
	"fmt"
	"go-backend/initializers"
	model "go-backend/models"
	"go-backend/utils"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const USER_KEY = "user"

func RequireAuth(ctx *gin.Context) {

	authorizationHeader := ctx.GetHeader("Authorization")

	if authorizationHeader == "" {

		ctx.AbortWithStatus(http.StatusUnauthorized)
		SetResponse(ctx, http.StatusUnauthorized)
		return
	}

	if !strings.HasPrefix(authorizationHeader, "Bearer ") {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		SetResponse(ctx, http.StatusUnauthorized)
		return
	}

	if strings.Contains(authorizationHeader, "undefined") {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		SetResponse(ctx, http.StatusUnauthorized)
		return
	}

	tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")
	decText, err := utils.Decrypt(tokenString)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	token, err := jwt.Parse(decText, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		if float64(time.Now().Unix()) > claims["expiration"].(float64){
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		var user model.User
		err := initializers.DB.First(&user,"id = ?", claims["id"]).Error;
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		ctx.Set(USER_KEY, user)

		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}