package models

type User struct {
	UserID       int    `json:"user_id" gorm:"primaryKey;column:user_id"`
	Email        string `json:"email" gorm:"unique;not null"`
	PasswordHash string `json:"password_hash" gorm:"not null"`
	//Customer     *Customer `gorm:"foreignKey:UserID" json:"customer,omitempty"`  // Relasi ke Customer
}

type Customer struct {
	CustomerID  int    `json:"customer_id" gorm:"primaryKey;column:customer_id"`
	UserID      int    `json:"user_id" gorm:"not null"` // Ini adalah foreign key ke User
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Address     string `json:"address" validate:"required"`
	//User        User   `gorm:"foreignKey:UserID" json:"user,omitempty"`  // Relasi ke User
}

/*?
type User struct {
	UserID       uint      `json:"user_id" gorm:"primaryKey"`
	// Kolom lainnya...
}

func (User) TableName() string {
	return "Users"  // Menentukan nama tabel yang berbeda jika tidak sesuai dengan konvensi GORM
}
?*/
// Region
/*
sekarang kan kita membuat struct dengan nama User dan Customer,
dari mana kita tau kalo itu bind ke table Users dan Customers,  dari nama saja sudah berbeda ?


Itu pertanyaan yang bagus! Meskipun nama struct (User dan Customer)
di kode Go mungkin berbeda dengan nama tabel di database (Users dan Customers),
hubungan antara struct dan tabel diatur oleh tag GORM yang ada pada setiap field di struct.
Tag ini memberi tahu GORM bagaimana menghubungkan field di struct dengan kolom di tabel database.

GORM dan Tagging
GORM memungkinkan kita untuk menentukan nama tabel, kolom, dan pengaturan lainnya menggunakan tag di struct.
Tag gorm inilah yang memastikan bahwa field dalam struct kita akan terkait dengan kolom yang tepat di tabel database.
Tag ini juga bisa mengatur nama tabel (jika berbeda) dan beberapa pengaturan lainnya, seperti primary key, unique, dan foreign key.

*/
// Region

//type Bookings struct {
//	BookingID     int    `json:"booking_id" gorm:"primaryKey;column:booking_id"`
//	TourName      string `json:"tour_name" gorm:"not null"`
//	BookingDate   string `json:"booking_date" gorm:"not null"`
//	BookingStatus string `json:"booking_status" gorm:"not null"`
//}
