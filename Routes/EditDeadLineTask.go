package routes

import (
	"batleforc/tp-cloud/model"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EditDeadLineTaskBody struct {
	DeadLine time.Time `json:"deadline"`
}

// EditTitleTask godoc
// @Summary Edit todoTask DeadLine
// @Param Label body routes.EditDeadLineTaskBody true "Label"
// @Param id  path string true "Task Id"
// @Success 200 {object} model.Task
// @Router /tache/{id}/deadline [put]
func EditDeadLineTask(c echo.Context) error {
	Client := c.Get("Client").(*mongo.Client)
	DbHandler := c.Get("DB").(model.Db)
	coll := model.GetTaskColl(DbHandler, Client)
	id := c.Param("id")
	idHex, _ := primitive.ObjectIDFromHex(id)
	boudy := new(EditDeadLineTaskBody)
	if err := c.Bind(boudy); err != nil {
		return c.JSON(http.StatusBadRequest, "Body invalid")
	}
	updateResult, err := coll.UpdateOne(context.TODO(), bson.D{{"_id", idHex}}, bson.D{{"$set", bson.D{{"deadLine", boudy.DeadLine}}}})
	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, fmt.Sprintf("Task update not working : %s", err))
	}
	if updateResult.ModifiedCount != 1 {
		return c.JSON(http.StatusBadRequest, boudy)
	}
	task := model.Task{}
	task.GetTaskById(DbHandler, Client, idHex)
	return c.JSON(http.StatusOK, task)
}