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

	e.POST("/add", add)
	e.GET("/", getByID)
	e.GET("/all", getAll)

	e.DELETE("/", deleteByID)

	e.PUT("/update", updateByID)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", 8080)))

}

func add(c echo.Context) error {
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

func getByID(c echo.Context) error {
	param := c.QueryParam("id")
	id, _ := strconv.ParseInt(param, 10, 64)
	res, err := service.New(getDbConfig()).GetById(id)

	if err != nil {
		return handleHTTPError(err)
	}

	return c.JSON(http.StatusOK, res)
}

func getAll(c echo.Context) error {
	res, err := service.New(getDbConfig()).GetAll()
	if err != nil {
		return handleHTTPError(err)
	}
	return c.JSON(http.StatusOK, res)
}

func updateByID(c echo.Context) error {
	var req service.UpdateTodoRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err := service.New(getDbConfig()).Update(req)
	if err != nil {
		return handleHTTPError(err)
	}

	return c.NoContent(http.StatusOK)
}

func deleteByID(c echo.Context) error {
	param := c.QueryParam("id")
	id, _ := strconv.ParseInt(param, 10, 64)

	err := service.New(getDbConfig()).DeleteByID(id)

	if err != nil {
		return handleHTTPError(err)
	}

	return c.NoContent(http.StatusOK)
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
