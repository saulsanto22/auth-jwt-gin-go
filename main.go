package main

import (
	"go-jwt-auth/config"
	"go-jwt-auth/model"
	"go-jwt-auth/routes"

	_ "go-jwt-auth/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go JWT Auth API
// @version 1.0
// @description This is a simple JWT authentication API built with Gin and GORM.
// @host localhost:8080
// @BasePath /
func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&model.User{})

	r := gin.Default()

	routes.SetupRoutes(r)
	//Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
