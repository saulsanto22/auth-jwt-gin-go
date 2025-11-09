package repository

import (
	"fmt"
	"go-jwt-auth/config"
	"go-jwt-auth/model"
)

type UserRepository struct{}

func NewRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindAll(page, limit int, search, role, sortBy, order string) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	q := config.DB.Model(&model.User{})

	if search != "" {
		q = q.Where("nama LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if role != "" {
		q = q.Where("role = ?", role)
	}

	//validasi col sort
	allowedSort := map[string]bool{
		"id":         true,
		"nama":       true,
		"email":      true,
		"created_at": true,
	}

	if !allowedSort[sortBy] {
		sortBy = "id"
	}

	if order != "asc" && order != "desc" {
		order = "asc"
	}
	q.Count(&total)

	offset := (page - 1) * limit
	err := q.Order(fmt.Sprintf("%s %s", sortBy, order)).
		Limit(limit).
		Offset(offset).
		Find(&config.DB.ClauseBuilders).Error

	if err != nil {
		return nil, 0, err
	}

	return users, total, nil

}

func (r *UserRepository) FindByID(id uint) (*model.User, error) {

	var user model.User

	err := config.DB.First(&user, id).Error
	return &user, err

}

func (r *UserRepository) CreateUser(user *model.User) error {
	return config.DB.Create(user).Error

}

func (r *UserRepository) UpdateUser(user *model.User) error {
	return config.DB.Save(user).Error

}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := config.DB.Where("email =?", email).First(&user).Error

	return &user, err
}

func (r *UserRepository) DeleteUser(id uint) error {

	var user model.User
	return config.DB.Delete(&user, id).Error

}
