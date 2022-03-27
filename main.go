package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	routes "batleforc/tp-cloud/Routes"
	_ "batleforc/tp-cloud/docs"
	"batleforc/tp-cloud/model"
)

var DB = "mongodb+srv://tp-cloud:RhtX1ZRUcXik5ZLk@db1.pg9nx.mongodb.net/?retryWrites=true&w=majority"

// @title Todolist Cloud Deployment
// @version 1.0
// @description Api TODO List du cours de virtualisation

// @BasePath /api
func main() {
	e := echo.New()
	DbHandler := model.Db{}
	if val, exist := os.LookupEnv("DB"); exist {
		DbHandler.Init(val, os.Getenv("ENV"))
	} else {
		DbHandler.Init(DB, os.Getenv("ENV"))
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${remote_ip} : ${time_rfc3339_nano}] ${status} : ${method} => ${uri}\n",
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: strings.Split("http://localhost:3000,http://localhost:3001,https://niort-tpcloud-front.herokuapp.com", ","),
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAuthorization, echo.HeaderAccept},
	}))
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	task := e.Group("/api/tache")

	task.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			client, err := DbHandler.GetClient()
			if err != nil {
				return echo.NewHTTPError(http.StatusServiceUnavailable, fmt.Sprintf("Database init Err %s", err))
			}
			err = client.Connect(context.TODO())
			defer client.Disconnect(context.TODO())
			c.Set("DB", DbHandler)
			c.Set("Client", client)
			return next(c)
		}
	})

	task.POST("", routes.AddTask)
	task.GET("", routes.ListTask)
	task.GET("/:id", routes.GetTask)

	task.PUT("/:id", routes.EditTitleTask)
	task.PUT("/:id/change-statut", routes.ChangeStatusTask)
	task.PUT("/:id/deadline", routes.EditDeadLineTask)
	task.DELETE("/:id", routes.DeleteTask)

	if val, exist := os.LookupEnv("PORT"); exist {
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", val)))
	} else {
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", "3001")))
	}

}
