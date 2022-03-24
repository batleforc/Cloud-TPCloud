package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	routes "batleforc/tp-cloud/Routes"
	_ "batleforc/tp-cloud/docs"
)

var DB = "mongodb"

// @title Todolist Cloud Deployment
// @version 1.0
// @description Api TODO List du cours de virtualisation

// @BasePath /api
func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${remote_ip} : ${time_rfc3339_nano}] ${status} : ${method} => ${uri}\n",
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: strings.Split("http://localhost:3000,http://localhost:3001", ","),
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAuthorization, echo.HeaderAccept},
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	task := e.Group("/api/tache")

	task.POST("", routes.AddTask)
	task.GET("", routes.ListTask)
	task.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})

	task.PUT("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})
	task.PUT("/:id/change-statut", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})
	task.DELETE("/:id", func(c echo.Context) error {
		type test struct {
			Val []string
			Key []string
			Grr string
		}
		return c.JSON(http.StatusOK, test{c.ParamValues(), c.ParamNames(), c.Param("id.json")})
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
