package handler

import (
	"fmt"
	"goauth/internal/model"
	"goauth/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Ping (ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "PONG",
	})
}

func GetAllUser (ctx *gin.Context) {
	
}

func CreateUser ( db *gorm.DB) gin.HandlerFunc {
	return  func(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	if err := repository.InsertUser(db, &user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}
	fmt.Println(user)
	ctx.JSON(http.StatusCreated, gin.H{
		"status" : 201,
		"data" : user,
	})
}
}

func UpdateUser (ctx *gin.Context) {
	
}

func DellUser (ctx *gin.Context) {
	
}
