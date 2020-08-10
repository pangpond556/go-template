package controllers

import (
	"appname/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Index root router
func Index(c echo.Context) error {
	models.Result["message"] = "success"
	models.Result["data"] = nil
	return c.JSON(http.StatusOK, models.Result)
}
