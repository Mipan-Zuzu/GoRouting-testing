package handler

import (
	"fmt"
	"goauth/internal/model"
	"goauth/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "PONG",
	})
}

func GetAllUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user []model.User
		result := db.Find(&user)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error,
			})
			return
		}

		if len(user) == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Empty data",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   user,
		})
	}
}

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := repository.InsertUser(db, &user); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(user)
		ctx.JSON(http.StatusCreated, gin.H{
			"status": 201,
			"data":   user,
		})
	}
}

func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result := ctx.Param("id")
		id, _ := strconv.Atoi(result)
		var payload map[string]interface{}
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			ctx.JSON(400, gin.H{
				"status":  400,
				"message": payload,
			})
			fmt.Println("payload :", payload)
			return
		}
		db.Model(&model.User{}).Where("id = ?", id).Updates(payload)
		ctx.JSON(200, gin.H{
			"status":  200,
			"message": payload,
		})
	}
}

func DellUser(db *gorm.DB) gin.HandlerFunc{
	return  func(ctx *gin.Context) {
		result := ctx.Param("id")
		id,errors := strconv.Atoi(result)
		var user *model.User
		db.Find(&user, id)
		if user.ID == 0 {
			ctx.JSON(200, gin.H{
				"status" : 200,
				"data" : fmt.Sprintf("Cannot delete user with id %s user was empty", result),
			})
			return
		}
		if errors != nil {
			ctx.JSON(400, gin.H{
				"status" : 400,
				"data" : "invalid id ",
			})
			return
		}
		db.Delete(&user, id)
		ctx.JSON(200, gin.H{
			"status" : 200,
			"data" : user,
		})
	}
}

// TODO: Advance
func GetUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Param("id")
		id, _ := strconv.Atoi(param)
		var user *model.User
		db.First(&user, id)
		if user.ID == 0 {
			ctx.JSON(200, gin.H{
				"status" : 200,
				"message" : "succses with no data",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   user,
		})
	}
}
