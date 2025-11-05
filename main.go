package main

import (
	"go-jwt-auth/config"
	"go-jwt-auth/model"
	"go-jwt-auth/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&model.User{})

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
