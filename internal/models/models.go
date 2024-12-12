package models

type RegisterRequest struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Address     string `json:"address" validate:"required"`
}

// RegisterResponse BISA GINI JUGA
/*
type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}
*/
type RegisterResponse struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
