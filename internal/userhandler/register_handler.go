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
	var req internal.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to generate password"})
	}
	user := internal.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "register failed U"})
	}

	customer := internal.Customer{
		UserID:      user.UserID,
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
	}
	if err := config.DB.Create(&customer).Error; err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "register failed C"})
	}
	return c.JSON(http.StatusOK, internal.RegisterResponse{
		UserID: user.UserID,
		Email:  user.Email,
	})

}
