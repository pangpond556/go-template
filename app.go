package main

import (
	"appname/routes"
	"appname/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	err := utils.Init()
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Register Routes
	routes.HomeRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
