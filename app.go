package main

import (
	"appname/routers"
	"appname/utils"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	err := utils.Config()
	if err != nil {
		log.Fatal(err)
	}

	// Register Routes
	routers.HomeRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
