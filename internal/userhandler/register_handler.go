package userhandler

import (
	"github.com/labstack/echo/v4"
	"live-code-2-XioweL/config"
	internal "live-code-2-XioweL/internal/models"
	"net/http"
)

func Register(c echo.Context) error {
	var req internal.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	}

	hashedPassword, err := bycrypt.GenerateFromPassword([]byte(req.Password), bycrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "invalid generate password"})
	}
	user := internal.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}
	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "register failed"})
	}

	customer := internal.Customer{
		UserID:      user.ID,
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
	}
	if err := config.DB.Create(&customer).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "register failed"})
	}
	return c.JSON(http.StatusOK, internal.RegisterResponse{
		UserID: user.UserID,
		Email:  user.Email,
	})

}
