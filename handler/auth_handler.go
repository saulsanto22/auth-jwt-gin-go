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
	var input model.RegisterRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if !utils.ValidateStruct(ctx, &input) {
		return
	}

	user := model.User{
		Nama:     input.Nama,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := h.userService.Register(&user); err != nil {
		utils.Error(ctx, http.StatusBadRequest, err.Error())

		return
	}

	utils.Success(ctx, gin.H{"message": "Registrasi berhasil"}, "")
}
func (h *AuthHandler) Login(ctx *gin.Context) {
	var input model.LoginRequest
	fmt.Println(input)
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if !utils.ValidateStruct(ctx, &input) {
		return
	}
	token, err := h.userService.Login(input.Email, input.Password)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(ctx, gin.H{"token": token}, "Berhasil!")

}

func (h *AuthHandler) GetProfile(ctx *gin.Context) {
	UserID, err := utils.CheckContext[uint](ctx, "user_id")
	if err != nil {
		utils.Error(ctx, http.StatusUnauthorized, "User tidak ditemukan!")
		return
	}

	user, err := h.userService.GetById(UserID)
	if err != nil {
		utils.Error(ctx, http.StatusUnauthorized, "User tidak ditemukan!")
		return
	}

	utils.Success(ctx, user, "profile")
}
func (h *AuthHandler) UpdateProfile(ctx *gin.Context) {
	UserID, err := utils.CheckContext[uint](ctx, "user_id")
	if err != nil {
		utils.Error(ctx, http.StatusUnauthorized, "User tidak ditemukan!")
		return
	}
	var input model.UpdateUser
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.userService.GetById(UserID)

	if err != nil {
		utils.Error(ctx, http.StatusNotFound, "User tidak ditemukan")
		return
	}

	if !utils.ValidateStruct(ctx, input) {
		return
	}
	err = h.userService.UpdateUser(user)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, "Gagal memperbarui profil")
		return
	}

	utils.Success(ctx, user, "data behasil diubah!")
}
