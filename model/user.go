package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role" gorm:"default:user"`
}

type Order struct {
	gorm.Model
	UserID uint
	User   User
}
