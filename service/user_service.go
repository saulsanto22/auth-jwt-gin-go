package service

import (
	"errors"
	"fmt"
	"go-jwt-auth/model"
	"go-jwt-auth/repository"
	"go-jwt-auth/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}
func (s *UserService) Register(user *model.User) error {

	existingUser, _ := s.userRepository.FindByEmail(user.Email)

	fmt.Println(existingUser)

	if existingUser != nil && existingUser.ID != 0 {
		return errors.New("email sudah terdaftar")
	}
	if user.Role == "" {
		user.Role = utils.RoleUser
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashPassword)
	return s.userRepository.CreateUser(user)
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.userRepository.FindByEmail(email)

	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateToken(user.ID, user.Role, user.Role)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (s *UserService) GetAllUsers(page, limit int, search, role, sortBy, order string) ([]model.User, int64, error) {

	return s.userRepository.FindAll(page, limit, search, role, sortBy, order)
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.userRepository.CreateUser(user)
}

func (s *UserService) GetById(id uint) (*model.User, error) {
	return s.userRepository.FindByID(id)
}

func (s *UserService) UpdateUser(user *model.User) error {
	return s.userRepository.UpdateUser(user)
}

func (s *UserService) Delete(id uint) error {
	return s.userRepository.DeleteUser(id)
}
