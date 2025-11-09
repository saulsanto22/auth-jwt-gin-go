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

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{s}
}

func (h *UserHandler) GetUserByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := h.userService.GetById(uint(id))
	if err != nil {
		utils.Error(ctx, http.StatusNotFound, "User tidak ditemukan")
		return
	}
	utils.Success(ctx, user, "User ditemukan")
}

// CreateUser untuk menambahkan user baru
func (h *UserHandler) CreateUser(ctx *gin.Context) {
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
func (h *UserHandler) UpdateUser(ctx *gin.Context) {
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
	err = h.userService.UpdateUser(user)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, "Gagal memperbarui user")
		return
	}

	utils.Success(ctx, user, "User berhasil diperbarui")
}

// DeleteUser untuk menghapus user
func (h *UserHandler) DeleteUser(ctx *gin.Context) {
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

	pageString := ctx.DefaultQuery("page", "1")
	limitString := ctx.DefaultQuery("limit", "10")
	search := ctx.DefaultQuery("search", "")
	role := ctx.DefaultQuery("role", "")
	sortBy := ctx.DefaultQuery("sort_by", "id")
	order := ctx.DefaultQuery("order", "asc")

	page, err := strconv.Atoi(pageString)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitString)

	if err != nil || limit < 1 {
		limit = 10
	}
	users, total, err := h.userService.GetAllUsers(page, limit, search, role, sortBy, order)

	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, "Gagal mengambil data users")
		return
	}

	utils.Success(ctx, gin.H{
		"users": users,
		"meta": gin.H{
			"total": total,
			"page":  page,
			"limit": limit,
		},
	}, "Data users berhasil diambil")
}
