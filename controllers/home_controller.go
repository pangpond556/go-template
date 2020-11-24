package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Index - index router
func Index(c echo.Context) error {
	result := map[string]interface{}{
		"message": "",
		"payload": nil,
	}
	return c.JSON(http.StatusOK, result)
}

// LargeJSON - steam large json
func LargeJSON(c echo.Context) error {
	var d []string
	for i := 0; i < 100; i++ {
		d = append(d, "data"+strconv.Itoa(i))
	}
	result := map[string]interface{}{
		"message": "",
		"payload": map[string]interface{}{
			"data": d,
		},
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(result)
}
