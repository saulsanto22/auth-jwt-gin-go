package model

type RegisterRequest struct {
	Nama     string `json:"nama" validate:"required,min=3,max=50"`
	Email    string `gorm:"uniqueIndex" json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UpdateUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Nama     string `json:"name" validate:"required, min=3,max=50"`
}
