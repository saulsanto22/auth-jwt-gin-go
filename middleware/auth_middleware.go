package middleware

import (
	"fmt"
	"go-jwt-auth/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token tidak ditemukan!"})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))

		token, err := utils.ValidateToken(tokenString)

		if err != nil || !token.Valid {
			fmt.Println(token.Valid)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token tidak valid"})

			ctx.Abort()
			return
		}
		ctx.Next()
	}

}
