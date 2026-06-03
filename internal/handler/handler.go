package handler

import (
	"fmt"
	"goauth/internal/model"
	"goauth/internal/repository"
	"net/http"
	// "os/user"
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

 func GetUser(db *gorm.DB) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		param := ctx.Param("id")
		id,_ := strconv.Atoi(param)
		var user *model.User
		db.First(&user, id)
		ctx.JSON(http.StatusOK, gin.H{
			"status" : 200,
			"data" : user, 
		})
	}
 }

func DellUser(ctx *gin.Context) {

}

//TODO: Advance

func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		fmt.Println(id)
		result, _ := strconv.Atoi(id)

		var siswa *model.User
		db.First(&siswa, result)
		
		update := map[string]interface{} {

		}
		db.Model(&siswa).Where("id = ?", result).Updates(update)
		ctx.JSON(http.StatusOK, gin.H{
			"id": result,
		})
	}
}
