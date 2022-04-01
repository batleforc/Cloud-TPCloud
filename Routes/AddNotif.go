package routes

import (
	"net/http"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/labstack/echo/v4"
)

type AddNotifBody struct {
	Subscription webpush.Subscription `json:"subscription"`
}

// AddNotif godoc
// @Summary Add notification to the array of notification
// @Description Add notification to the array of notification
// @Param Notification body routes.AddNotifBody true "notification"
// @Router /notif [post]
func AddNotif(c echo.Context) error {
	NotifBuffer := c.Get("NotifBuffer").(*[]webpush.Subscription)
	boudy := new(AddNotifBody)
	if err := c.Bind(boudy); err != nil {
		return c.JSON(http.StatusBadRequest, "Body invalid")
	}
	*NotifBuffer = append(*NotifBuffer, boudy.Subscription)
	return c.JSON(http.StatusOK, "OK")
}
