package routes

import (
	"appname/controllers"

	"github.com/labstack/echo/v4"
)

// HomeRoutes - root router
func HomeRoutes(e *echo.Echo) {
	g := e.Group("/")
	{
		g.GET("", controllers.Index)
	}
}
