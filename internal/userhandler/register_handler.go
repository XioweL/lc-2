package userhandler

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"live-code-2-XioweL/config"
	internal "live-code-2-XioweL/internal/models"
	"log"
	"net/http"
)

func Register(c echo.Context) error {
	//? // Step 1: Bind input JSON ke struct
	var req internal.RegisterRequest
	// Bind request data to struct
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	}

	//! HASH NYA JUGA BISA KAYAK GINI
	/*
		func hashPassword(password string) (string, error) {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			return string(hashedPassword), err
		}

		passwordHash, err := hashPassword(input.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
		}

		user := models.User{
			Email:        input.Email,
			PasswordHash: passwordHash,
		}
	*/
	//? Step 2: Hash password menggunakan bcrypt
	//* Hash password before saving to database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to generate password"})
	}
	//* Save data to table users
	//* KENAPA CUMAN EMAIL DAN PASSWORD ? KARENA DI TABLE USERS HANYA ADA EMAIL DAN PASSWORD YANG KITA BUTUHKAN
	//* SISANYA ADA DI TABLE CUSTOMER
	/*
		-- Table: Users
		CREATE TABLE Users (
		    user_id SERIAL PRIMARY KEY,
		    email VARCHAR(255) NOT NULL UNIQUE,
		    password_hash VARCHAR(255) NOT NULL,
		    last_login_date TIMESTAMP,
		    jwt_token TEXT
		);

		-- Table: Customers
		CREATE TABLE Customers (
		    customer_id SERIAL PRIMARY KEY,
		    user_id INT REFERENCES Users(user_id),
		    name VARCHAR(255) NOT NULL,
		    email VARCHAR(255) NOT NULL,
		    phone_number VARCHAR(15),
		    address TEXT
		);
	*/
	//? Step 3: Simpan User ke tabel Users
	user := internal.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	/*
		?config.DB.Create(&user).Error:
		Menyimpan data pengguna baru ke dalam tabel users.
		?log.Println(err):
		Mencetak pesan kesalahan ke log jika terjadi kesalahan.
		?return c.JSON(http.StatusInternalServerError, map[string]string{"message": "register failed U"}):
		Mengembalikan respons JSON dengan status 500 Internal Server Error dan pesan kesalahan "register failed U" jika penyimpanan data gagal.
	*/
	if err := config.DB.Create(&user).Error; err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "register failed U"})
	}
	//? Step 4: Simpan Customer ke tabel Customers (gunakan user_id yang baru dibuat)
	//* Save data to table customers
	customer := internal.Customer{
		UserID:      user.UserID, //? Relasikan dengan User yang baru saja dibuat
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
	}
	/*
			-- Table: Customers
			CREATE TABLE Customers (
		    customer_id SERIAL PRIMARY KEY,
		    user_id INT REFERENCES Users(user_id),
		    name VARCHAR(255) NOT NULL,
		    email VARCHAR(255) NOT NULL,
		    phone_number VARCHAR(15),
		    address TEXT
		);
	*/

	//* Save data to table customers
	if err := config.DB.Create(&customer).Error; err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "register failed C"})
	}
	//? Step 5: Response sukses
	// Custom response for register success
	return c.JSON(http.StatusOK, internal.RegisterResponse{
		UserID: user.UserID,
		Email:  user.Email,
	})

	// Atau Response nya gini
	/*
		return c.JSON(http.StatusCreated, map[string]string{
		"message": "User and customer created successfully"
		})
	*/

}

/*
!Urutan Alur Kode:
!1. Dapatkan input dari pengguna — melalui struct RegisterInput.
!2. Hash password dan buat User — simpan ke tabel Users.
!3. Gunakan user_id untuk membuat entri Customer — simpan data ke tabel Customers.
*/
