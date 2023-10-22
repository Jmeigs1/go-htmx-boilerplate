package main

import (
	"context"
	"rename-me/templates"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		templates.Index("").Render(context.Background(), c.Response())
		return nil
	})

	e.GET("/:route", func(c echo.Context) error {
		hxr := c.Request().Header["Hx-Request"]
		if len(hxr) > 0 && hxr[0] == "true" {
			templates.Body(c.Param("route")).Render(context.Background(), c.Response())
		} else {
			templates.Index(c.Param("route")).Render(context.Background(), c.Response())

		}
		return nil
	})
	e.Logger.Fatal(e.Start(":1323"))
}
