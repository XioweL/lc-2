package authentication

import (
	"github.com/labstack/echo/v4"
	internal "live-code-2-XioweL/internal/models"
	"net/http"
)

type AuthenticationController struct {
	AuthenticationService *AuthenticationService
}

func NewAuthenticationController(authenticationService *AuthenticationService) *AuthenticationController {
	return &AuthenticationController{
		AuthenticationService: authenticationService,
	}
}

func (controller *AuthenticationController) RegisterCustomer(ctx echo.Context) error {
	var req internal.RegisterRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	}

	if err := controller.AuthenticationService.RegisterUser(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to register user"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "user registered successfully"})
}

// Todo: Nyieun Login Didieu saruakeun kos anu register
