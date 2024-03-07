package controllers

import (
	"go-backend/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartTransation(ctx *gin.Context) *gorm.DB {
	tx := initializers.DB.Begin()

	if tx.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		return nil
	}
	return tx
}

func UpdateModelByID(ctx *gin.Context, tx *gorm.DB, entity interface{}) error {
	id := ctx.Param("id")
	var err = tx.Model(entity).Where("id = ?", id).Updates(entity).Error;
	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update"})
		return err
	}

	tx.Commit()
	return nil
}

func DeleteModelByID(ctx *gin.Context, tx *gorm.DB, entity interface{}) error {
	id := ctx.Param("id")
	if err := tx.Where("id = ?", id).Delete(&entity).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return err
	}

	tx.Commit()
	return nil
}

func GetModels(ctx *gin.Context, entity interface{}) error {
	err := initializers.DB.Find(entity).Error;
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return err
	}
	return nil
}

func GetModelByID(ctx *gin.Context, entity interface{}) error {
	id := ctx.Param("id")
	err := initializers.DB.First(&entity, id).Error;
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return err
	}
	return nil
}

func BindModel(ctx *gin.Context, entity interface{}) error {
	err := ctx.ShouldBindJSON(&entity)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return err
	}
	return nil
}

func SetResponse(ctx *gin.Context, code int) {
	ctx.JSON(code, gin.H {
		"error": "Status Failed",
	})
}