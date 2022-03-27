package routes

import (
	"batleforc/tp-cloud/model"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type AddTaskBody struct {
	Label    string     `json:"label"`
	DeadLine *time.Time `json:"deadline,omitempty"`
}

// AddTask godoc
// @Summary Allow the user to create todotask
// @Description create todotask in mongodb
// @Param Label body routes.AddTaskBody true "Label"
// @Router /tache [post]
func AddTask(c echo.Context) error {
	Client := c.Get("Client").(*mongo.Client)
	DbHandler := c.Get("DB").(model.Db)
	coll := model.GetTaskColl(DbHandler, Client)

	boudy := new(AddTaskBody)
	if err := c.Bind(boudy); err != nil {
		return c.JSON(http.StatusBadRequest, "Body invalid")
	}
	coll.InsertOne(context.TODO(), model.CreateNextTask(boudy.Label, boudy.DeadLine))
	return c.JSON(http.StatusOK, "OK")
}
