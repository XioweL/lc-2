package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	internal "live-code-2-XioweL/internal/middleware"
	"live-code-2-XioweL/internal/src/authentication"
	handler "live-code-2-XioweL/internal/userhandler"
)

func SetupRoutes(db *gorm.DB) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	authentication := InitAuthenticationController(db)

	// User
	// todo: ganti handler login ku authentication login customer
	e.POST("/users/login", handler.Login)
	e.POST("/users/register", authentication.RegisterCustomer)

	// Booking
	// todo: booking implementasi keun clean architecture nu kos urang nu isina controller, repostiory, service.
	// todo: implementasikeun oge dependency injection siga InitAuthenticationController
	g := e.Group("/bookings")
	g.Use(internal.CustomJwtMiddleware)
	g.GET("", handler.GetBookings)
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

func InitAuthenticationController(db *gorm.DB) *authentication.AuthenticationController {
	authenticationRepository := authentication.NewAuthenticationRepository(db)
	authenticationService := authentication.NewAuthenticationService(authenticationRepository)
	authenticationController := authentication.NewAuthenticationController(authenticationService)

	return authenticationController
}
