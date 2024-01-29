package main

import (
	"TodoBackend/db"
	"TodoBackend/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/add", addTodo)
	e.GET("/", getTodo)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", 8080)))

}

func addTodo(c echo.Context) error {
	var req service.AddTodoRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	res, err := service.New(getDbConfig()).Add(req)
	if err != nil {
		return handleHTTPError(err)
	}

	return c.JSON(http.StatusCreated, res)
}

func getTodo(c echo.Context) error {
	param := c.QueryParam("id")
	id, _ := strconv.ParseInt(param, 10, 64)
	res, err := service.New(getDbConfig()).GetById(id)

	if err != nil {
		return handleHTTPError(err)
	}

	return c.JSON(http.StatusOK, res)
}

func handleHTTPError(err error) *echo.HTTPError {
	msg := err.Error()
	if strings.Contains(msg, "not found") {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return echo.NewHTTPError(http.StatusBadRequest)
}

func getDbConfig() db.Config {
	return db.Config{
		Username: "fara",
		Password: "123",
		Port:     3307,
		Host:     "localhost",
		DBName:   "todo_db",
	}
}
