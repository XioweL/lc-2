package userhandler

import (
	"github.com/labstack/echo/v4"
	internal "live-code-2-XioweL/internal/models"
	"net/http"
)

func Register(c echo.Context) error {
	var req internal.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	}

	hashedPassword, err := internal.HashPassword(req.Password)

}
