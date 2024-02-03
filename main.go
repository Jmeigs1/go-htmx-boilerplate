package main

import (
	"fmt"
	"rename-me/pages"
	"rename-me/routes"
	"rename-me/templates"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	PORT = 1323
	HOME = "/example"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Static("/static", "static")

	// Can add a prefix here if app needs one
	g := e.Group("")

	// Order of adding determines the order in sidebar
	pages.AddAllRoutes(g)
	routes.AddRoute("1", "/1", g, printRoute)
	routes.AddRoute("2", "/2", g, printRoute)
	routes.AddRoute("3", "/3", g, printRoute)

	e.Any("/*", redirectHome)
	fmt.Printf("Server hosted at http://localhost:%d\n", PORT)
	e.Logger.Fatal(e.Start(fmt.Sprint(":", PORT)))
}

func redirectHome(c echo.Context) error {
	return c.Redirect(302, HOME)
}

func printRoute(c echo.Context) error {
	route := c.Path()
	return templates.WrapMainRoute(templates.Other(route))(c)
}
