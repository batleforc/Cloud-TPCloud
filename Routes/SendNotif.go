package routes

import (
	"batleforc/tp-cloud/model"
	"net/http"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/labstack/echo/v4"
)

type SendNotifBody struct {
	Notif model.Notification `json:"notif"`
}

// SendNotif godoc
// @Summary Send notification to the array of notification
// @Description Send notification to the array of notification
// @Param Sub body routes.SendNotifBody true "subscription"
// @Router /notif/send [post]
func SendNotif(c echo.Context) error {
	NotifBuffer := c.Get("NotifBuffer").(*[]webpush.Subscription)
	PushNotif := c.Get("PushNotif").(*model.Push)
	boudy := new(SendNotifBody)
	if err := c.Bind(boudy); err != nil {
		return c.JSON(http.StatusBadRequest, "Body invalid")
	}
	PushNotif.SendToallsub(*NotifBuffer, "", boudy.Notif)
	return c.JSON(http.StatusOK, "OK")
}
