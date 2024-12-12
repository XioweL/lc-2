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

type Bookings struct {
	BookingID     int    `json:"booking_id" gorm:"primaryKey;column:booking_id"`
	TourName      string `json:"tour_name" gorm:"not null"`
	BookingDate   string `json:"booking_date" gorm:"not null"`
	BookingStatus string `json:"booking_status" gorm:"not null"`
}
