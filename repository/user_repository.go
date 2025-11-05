package repository

import (
	"go-jwt-auth/config"
	"go-jwt-auth/model"
)

type UserRepository struct{}

func NewRepository() *UserRepository {
	return &UserRepository{}
}
func (r *UserRepository) CreateUser(user *model.User) error {
	return config.DB.Create(user).Error

}

func (r *UserRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := config.DB.Where("email =?", email).First(&user).Error

	return user, err
}
