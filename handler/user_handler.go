package handler

import (
	"go-jwt-auth/model"
	"go-jwt-auth/service"
	"go-jwt-auth/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func (h *AuthHandler) GetUserByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := h.userService.GetById(uint(id))
	if err != nil {
		utils.Error(ctx, http.StatusNotFound, "User tidak ditemukan")
		return
	}
	utils.Success(ctx, user, "User ditemukan")
}

// CreateUser untuk menambahkan user baru
func (h *AuthHandler) CreateUser(ctx *gin.Context) {
	var input model.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.userService.CreateUser(&input)
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	utils.Created(ctx, input, "User berhasil dibuat")
}

// UpdateUser untuk memperbarui data user
func (h *AuthHandler) UpdateUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var input model.User

	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.userService.GetById(uint(id))
	if err != nil {
		utils.Error(ctx, http.StatusNotFound, "User tidak ditemukan")
		return
	}

	// Update data user
	user.Nama = input.Nama
	user.Email = input.Email
	err = h.userService.UpdateUser(&user)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, "Gagal memperbarui user")
		return
	}

	utils.Success(ctx, user, "User berhasil diperbarui")
}

// DeleteUser untuk menghapus user
func (h *AuthHandler) DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := h.userService.Delete(uint(id))
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, "Gagal menghapus user")
		return
	}

	utils.Success(ctx, nil, "User berhasil dihapus")
}
func NewService(userService *service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	users, err := h.userService.GetAllUsers()

	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, "Gagal mengambil data users")
		return
	}

	utils.Success(ctx, users, "Data Berhasil!")
}
