package userhandler

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"live-code-2-XioweL/config"
	internal "live-code-2-XioweL/internal/models"
	"log"
	"net/http"
	"time"
)

// todo: ke ieu delete pindahkeun ka src/authentication
func Login(c echo.Context) error {
	var req internal.LoginRequest
	if err := c.Bind(&req); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Request"})
	}

	var user internal.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Email or Password"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Email or Password"})
	}

	var customer internal.Customer
	if err := config.DB.Where("user_id = ?", user.UserID).First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Email or Password"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":           customer.ID,
		"user_id":      user.UserID,
		"name":         customer.Name,
		"email":        customer.Email,
		"phone_number": customer.PhoneNumber,
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(config.JwtSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Generate Token"})
	}

	return c.JSON(http.StatusOK, internal.LoginResponse{Token: tokenString})
}
