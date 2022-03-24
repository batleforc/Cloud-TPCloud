package routes

import (
	"batleforc/tp-cloud/model"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ListTask godoc
// @Summary Allow the user to get all todotask
// @Description get all todotask in mongodb
// @Success 200 {array} model.Task
// @Router /tache [get]
func ListTask(c echo.Context) error {
	Client := c.Get("Client").(*mongo.Client)
	DbHandler := c.Get("DB").(model.Db)
	coll := model.GetTaskColl(DbHandler, Client)
	Tasks := []model.Task{}
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return c.String(http.StatusInternalServerError, "Fetch data dont work")
	}
	for cursor.Next(context.TODO()) {
		var tache model.Task
		err := cursor.Decode(&tache)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		Tasks = append(Tasks, tache)
	}
	return c.JSON(http.StatusOK, Tasks)
}
