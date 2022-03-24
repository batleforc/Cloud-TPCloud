package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// AddTask godoc
// @Summary Allow the user to create todotask
// @Description create todotask in mongodb
// @Router /tache [get]
func AddTask(c echo.Context) error {
	return c.String(http.StatusOK, "truc")
}
