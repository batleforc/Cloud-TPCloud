package routes

import (
	"batleforc/tp-cloud/model"
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EditTitleTaskBody struct {
	Label string `json:"label"`
}

// EditTitleTask godoc
// @Summary Edit todoTask Label
// @Param Label body routes.EditTitleTaskBody true "Label"
// @Param id  path string true "Task Id"
// @Success 200 {object} model.Task
// @Router /tache/{id} [put]
func EditTitleTask(c echo.Context) error {
	Client := c.Get("Client").(*mongo.Client)
	DbHandler := c.Get("DB").(model.Db)
	coll := model.GetTaskColl(DbHandler, Client)
	id := c.Param("id")
	idHex, _ := primitive.ObjectIDFromHex(id)
	boudy := new(EditTitleTaskBody)
	if err := c.Bind(boudy); err != nil {
		return c.JSON(http.StatusBadRequest, "Body invalid")
	}
	updateResult, err := coll.UpdateOne(context.TODO(), bson.D{{"_id", idHex}}, bson.D{{"$set", bson.D{{"label", boudy.Label}}}})
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
