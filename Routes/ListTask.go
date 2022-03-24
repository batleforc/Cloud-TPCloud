package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ListTask godoc
// @Summary Allow the user to get all todotask
// @Description get all todotask in mongodb
// @Router /tache [get]
func ListTask(c echo.Context) error {
	return c.String(http.StatusOK, "HELLO WORLD")
}
