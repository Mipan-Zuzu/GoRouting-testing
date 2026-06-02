package main

import (
	"fmt"
	"goauth/internal/handler"
	"goauth/internal/model"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main () {
	godotenv.Load()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
	os.Getenv("DB_HOST"),
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASS"),
	os.Getenv("DB_NAME"),
	os.Getenv("DB_PORT"),
	)
	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		panic("Failed connected to database")
	}

	db.AutoMigrate(&model.User{})

	route := gin.Default()
	route.GET("/ping", handler.Ping)
	SetupRouting(route, db)
	route.Run(":3000")
	fmt.Println("Starting run project...")
	fmt.Println("Succses conected to database")
}

func SetupRouting (route *gin.Engine, db *gorm.DB) {
	route.GET("/api/user", handler.GetAllUser)
	route.POST("/api/user", handler.CreateUser(db))
	route.PATCH("/api/user/:id", handler.UpdateUser)
	route.DELETE("/api/user/:id", handler.DellUser)
}