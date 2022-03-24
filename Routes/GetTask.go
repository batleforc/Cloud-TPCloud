package routes

import (
	"batleforc/tp-cloud/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetTask godoc
// @Summary Allow the user to get one todotask
// @Description get one todotask in mongodb
// @Success 200 {array} model.Task
// @Param id  path string true "Task Id"
// @Router /tache/{id} [get]
func GetTask(c echo.Context) error {
	Client := c.Get("Client").(*mongo.Client)
	DbHandler := c.Get("DB").(model.Db)
	//coll := model.GetTaskColl(DbHandler, Client)
	id := c.Param("id")
	idHex, _ := primitive.ObjectIDFromHex(id)
	task := model.Task{}
	err := task.GetTaskById(DbHandler, Client, idHex)
	//err := coll.FindOne(context.TODO(), bson.M{"_id": idHex}).Decode(&task)
	if err != nil {
		return c.String(404, fmt.Sprintf("data not found, %s", err))
	}
	return c.JSON(http.StatusOK, task)
}
