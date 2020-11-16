package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Index root router
func Index(c echo.Context) error {
	result := map[string]interface{}{
		"status":  "success",
		"message": "",
		"payload": nil,
	}
	return c.JSON(http.StatusOK, result)
}
