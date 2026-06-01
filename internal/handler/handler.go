package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Ping (ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "PONG",
	})
}

func GetAllUser (ctx *gin.Context) {
	
}

func CreateUser (ctx *gin.Context) {
	
}

func UpdateUser (ctx *gin.Context) {
	
}

func DellUser (ctx *gin.Context) {
	
}
