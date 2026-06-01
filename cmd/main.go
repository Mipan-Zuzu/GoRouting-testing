package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goauth/internal/handler"
)

func main () {
	fmt.Println("Starting run project...")
	route := gin.Default()
	route.GET("/ping", handler.Ping)
	route.Run(":3000")
}

func Routing () {
	
}