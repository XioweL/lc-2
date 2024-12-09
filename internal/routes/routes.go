package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	//internal "live-code-2-XioweL/internal/middleware"
	handler "live-code-2-XioweL/internal/userhandler"
)

func SetupRoutes() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// User
	e.POST("/users/login", handler.Login)
	e.POST("/users/register", handler.Register)

	//// Booking
	//g := e.Group("/bookings")
	//g.Use(internal.CustomJwtMiddleware)
	//g.GET("", handler.GetBookings)
	//g.GET("/unpaid", handler.GetUnpaidBook)
	//
	//// Tours
	//e.GET("/tours/earning", handler.GetEarningData, internal.CustomJwtMiddleware)
	//
	//// Reports
	//auth := e.Group("/reports")
	//auth.Use(internal.CustomJwtMiddleware)
	//auth.GET("/total-customers", handler.GetTotalCustomers)
	//auth.GET("/bookings-per-tour", handler.GetBookingsPerTour)

	return e
}
