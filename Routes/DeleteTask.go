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

// DeleteTask godoc
// @Summary Delete todoTask
// @Param id  path string true "Task Id"
// @Success 200 {object} model.Task
// @Router /tache/{id} [delete]
func DeleteTask(c echo.Context) error {
	Client := c.Get("Client").(*mongo.Client)
	DbHandler := c.Get("DB").(model.Db)
	coll := model.GetTaskColl(DbHandler, Client)
	id := c.Param("id")
	idHex, _ := primitive.ObjectIDFromHex(id)

	delResult, err := coll.DeleteOne(context.TODO(), bson.D{{"_id", idHex}})
	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, fmt.Sprintf("Task delete not working : %s", err))
	}
	if delResult.DeletedCount != 1 {
		return c.JSON(http.StatusBadRequest, id)
	}
	return c.JSON(http.StatusOK, "OK")
}
