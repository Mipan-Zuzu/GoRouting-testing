package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goauth/internal/handler"
	"goauth/internal/model"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func main () {
	db, error := gorm.Open(postgres.Open("usergo"), &gorm.Config{})
	if error != nil {
		panic("Failed connected to database")
	}

	db.AutoMigrate(&model.User{})
	
	route := gin.Default()
	route.GET("/ping", handler.Ping)
	SetupRouting(route)
	route.Run(":3000")
	fmt.Println("Starting run project...")
	fmt.Println("Succses conected to database")
}

func SetupRouting (route *gin.Engine) {
	route.GET("/api/user", handler.GetAllUser)
	route.POST("/api/user", handler.CreateUser)
	route.PATCH("/api/user/:id", handler.UpdateUser)
	route.DELETE("/api/user/:id", handler.DellUser)
}