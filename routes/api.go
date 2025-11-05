package routes

import (
	"go-jwt-auth/handler"
	"go-jwt-auth/middleware"
	"go-jwt-auth/repository"
	"go-jwt-auth/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	userRepository := repository.NewRepository()
	userService := service.NewUserService(userRepository)
	authHandler := handler.NewAuthHandler(userService)

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	auth := r.Group("/auth", middleware.AuthMiddleware())
	auth.GET("/profile", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "ini setelah login!"})
	})

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, "hai")
	})
}
