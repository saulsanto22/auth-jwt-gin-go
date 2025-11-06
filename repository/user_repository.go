package repository

import (
	"go-jwt-auth/config"
	"go-jwt-auth/model"
)

type UserRepository struct{}

func NewRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindAll(user *[]model.User) error {
	return config.DB.Find(user).Error

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

func (r *UserRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := config.DB.Where("email =?", email).First(&user).Error

	return user, err
}

func (r *UserRepository) DeleteUser(id uint) error {

	var user model.User
	return config.DB.Delete(&user, id).Error

}
