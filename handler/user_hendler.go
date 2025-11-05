package handler

import (
	"fmt"
	"go-jwt-auth/model"
	"go-jwt-auth/service"
	"go-jwt-auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{userService}

}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var input model.User

	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userService.Register(&input); err != nil {
		utils.Error(ctx, http.StatusBadRequest, err.Error())

		return
	}

	utils.Success(ctx, gin.H{"message": "Registrasi berhasil"}, "")
}
func (h *AuthHandler) Login(ctx *gin.Context) {
	var input model.User
	fmt.Println(input)
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.userService.Login(input.Email, input.Password)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(ctx, gin.H{"token": token}, "Berhasil!")

}
