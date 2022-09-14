package controllers

import (
	"net/http"
	"day2/lib/database"
	"day2/models"

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