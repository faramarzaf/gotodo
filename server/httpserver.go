package server

import (
	"TodoBackend/db"
	"TodoBackend/dto"
	"TodoBackend/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
	"strings"
)

func Serve(port int) {
	e := echo.New()

	//more config at middleware.LoggerConfig
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "{remoteIp=${remote_ip}, method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}}\n",
	}))

	e.Use(middleware.Recover())

	e.POST("/add", add)
	e.GET("/", getByID)
	e.GET("/all", getAll)
	e.DELETE("/", deleteByID)
	e.PUT("/update", updateByID)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

func add(c echo.Context) error {
	var req dto.AddTodoRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	res, err := service.New(db.GetDbConfig()).Add(req)
	if err != nil {
		return handleHTTPError(err)
	}

	return c.JSON(http.StatusCreated, res)
}

func getByID(c echo.Context) error {
	param := c.QueryParam("id")
	id, _ := strconv.ParseInt(param, 10, 64)
	res, err := service.New(db.GetDbConfig()).GetById(id)

	if err != nil {
		return handleHTTPError(err)
	}

	return c.JSON(http.StatusOK, res)
}

func getAll(c echo.Context) error {
	res, err := service.New(db.GetDbConfig()).GetAll()
	if err != nil {
		return handleHTTPError(err)
	}
	return c.JSON(http.StatusOK, res)
}

func updateByID(c echo.Context) error {
	var req dto.UpdateTodoRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err := service.New(db.GetDbConfig()).Update(req)
	if err != nil {
		return handleHTTPError(err)
	}

	return c.NoContent(http.StatusOK)
}

func deleteByID(c echo.Context) error {
	param := c.QueryParam("id")
	id, _ := strconv.ParseInt(param, 10, 64)

	err := service.New(db.GetDbConfig()).DeleteByID(id)

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
