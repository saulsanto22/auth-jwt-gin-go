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
	adminHandler := handler.NewUserHandler(userService)

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	auth := r.Group("/auth", middleware.AuthMiddleware())
	auth.GET("/profile")

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, "hai")
	})

	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminOnlyMiddleware())
	{
		// admin.GET("/users", func(ctx *gin.Context) {
		// 	// Akses data user hanya untuk admin
		// 	var users []model.User
		// 	if err := config.DB.Find(&users).Error; err != nil {
		// 		ctx.JSON(500, gin.H{"error": "failed to fetch users"})
		// 		return
		// 	}
		// 	ctx.JSON(200, gin.H{"data": users})
		// })

		admin.GET("users/", adminHandler.GetAllUsers)
		admin.GET("users/:id", adminHandler.GetUserByID)
		admin.POST("user", adminHandler.CreateUser)
		admin.PUT("user/:id", adminHandler.UpdateUser)
		admin.DELETE("user/:id", adminHandler.DeleteUser)
	}

}
