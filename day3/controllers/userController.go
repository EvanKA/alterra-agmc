package controllers

import (
	"net/http"
	"day3/lib/database"
	"day3/models"

	"github.com/labstack/echo"
)

func GetUserControllers (c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewtHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{} {
		"status": "success",
		"users": users
	})
}