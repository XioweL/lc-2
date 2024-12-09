package models

type User struct {
	UserID       int      `json:"user_id" gorm:"primaryKey;column:user_id"`
	Email        string   `json:"email" gorm:"unique;not null"`
	PasswordHash string   `json:"password_hash" gorm:"not null"`
	Customer     Customer `gorm:"foreignKey:UserID" json:"customer,omitempty"`
}
type Customer struct {
	ID     int    `json:"customer_id" gorm:"primaryKey;column:customer_id"`
	UserID int    `json:"user_id" gorm:"not null"`
	Name   string `json:"name" validate:"required"`
	Email  string `json:"email" validate:"required,email"`
	//Password    string `json:"password" validate:"required,min=8"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Address     string `json:"address" validate:"required"`
}
