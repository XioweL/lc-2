package middleware

import (
	"github.com/labstack/echo/v4"
	"live-code-2-XioweL/common/constants"
	"live-code-2-XioweL/internal/models"
)

func GetActor(ctx echo.Context) models.CustomerData {
	actor := ctx.Get(constants.ActorUserContext).(models.CustomerData)

	return models.CustomerData{
		ID:          actor.ID,
		UserID:      actor.UserID,
		Name:        actor.Name,
		Email:       actor.Email,
		PhoneNumber: actor.PhoneNumber,
	}
}
