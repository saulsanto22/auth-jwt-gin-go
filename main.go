package main

import (
	"go-jwt-auth/config"
	"go-jwt-auth/handler"
	"go-jwt-auth/middleware"
	"go-jwt-auth/model"
	"go-jwt-auth/repository"
	"go-jwt-auth/service"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&model.User{})

	userRepository := repository.NewRepository()
	userService := service.NewUserService(userRepository)
	authHandler := handler.NewAuthHandler(userService)
	r := gin.Default()

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	auth := r.Group("/auth", middleware.AuthMiddleware())
	auth.GET("/profile", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "ini setelah login!"})
	})

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, "hai")
	})
	r.Run(":8080")
}
