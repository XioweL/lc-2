package userhandler

import (
	"github.com/labstack/echo/v4"
	"live-code-2-XioweL/config"
	"net/http"
)

func GetBookings(c echo.Context) error {
	userClaims := c.Get("user").(map[string]interface{})
	userID := userClaims["user_id"].(float64)

	//	bookings := internal.Bookings{}
	//	if err := config.DB.Table("bookings").
	//		Select("bookings.booking_id, tours.tour_name, bookings.booking_date, bookings.booking_status").
	//		Joins("JOIN tour_bookings ON bookings.booking_id = tour_bookings.booking_id").
	//		Joins("JOIN tours ON tour_bookings.tour_id = tours.tour_id").
	//		Where("bookings.user_id = ?", userID).
	//		Scan(&bookings).Error; err != nil {
	//		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch bookings"})
	//	}
	//	return c.JSON(http.StatusOK, internal.Bookings{})
	//}
	var bookings []struct {
		BookingID     int    `json:"booking_id"`
		TourName      string `json:"tour_name"`
		BookingDate   string `json:"booking_date"`
		BookingStatus string `json:"booking_status"`
	}
	if err := config.DB.Table("bookings").
		Select("bookings.booking_id, tours.tour_name, bookings.booking_date, bookings.booking_status").
		Joins("JOIN tour_bookings ON bookings.booking_id = tour_bookings.booking_id").
		Joins("JOIN tours ON tour_bookings.tour_id = tours.tour_id").
		Where("bookings.customer_id = ?", userID).
		Scan(&bookings).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch bookings"})
	}
	return c.JSON(http.StatusOK, bookings)
}
