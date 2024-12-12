package middleware

import (
	"live-code-2-XioweL/common/constants"
	"live-code-2-XioweL/config"
	"live-code-2-XioweL/internal/models"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func CustomJwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // Tidak ada "Bearer"
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
			}
			return config.JwtSecret, nil
			//return []byte("Secret_Key"), nil
		})

		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			actor := models.CustomerData{
				ID:          int(claims["id"].(float64)),
				UserID:      int(claims["user_id"].(float64)),
				Name:        claims["name"].(string),
				Email:       claims["email"].(string),
				PhoneNumber: claims["phone_number"].(string),
			}

			c.Set(constants.ActorUserContext, actor)
		}

		return next(c)
	}
}
