package middleware

import (
	"go-jwt-auth/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			utils.Error(ctx, http.StatusUnauthorized, "token tidak valid")
			ctx.Abort()
			return
		}

		tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))

		token, err := utils.ValidateToken(tokenString)

		if err != nil {

			utils.Error(ctx, http.StatusUnauthorized, "token tidak valid")

			ctx.Abort()
			return
		}

		ctx.Set("user_id", token.UserID)
		ctx.Set("email", token.Email)
		ctx.Set("user_role", token.Role)

		ctx.Next()

	}
}

func AdminOnlyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.GetString("user_role")

		if role != utils.RoleAdmin {
			utils.Error(ctx, 401, "error")
			ctx.Abort()

			return
		}

		ctx.Next()

	}
}
