package main

import (
	"rename-me/pages/example"
	"rename-me/templates"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Static("/static", "static")

	e.GET("/", redirectHome)
	example.AddRoutes(e)
	e.GET("/1", printRoute)
	e.GET("/2", printRoute)
	e.GET("/3", printRoute)

	e.Logger.Fatal(e.Start(":1323"))
}

func redirectHome(c echo.Context) error {
	return c.Redirect(302, "/example")
}

func printRoute(c echo.Context) error {
	route := c.Path()
	return templates.WrapMainRoute(templates.Other(route))(c)
}
