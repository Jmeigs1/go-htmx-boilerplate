package example

import (
	"rename-me/templates"

	"github.com/labstack/echo/v4"
)

func AddRoutes(e *echo.Echo) {
	example := templates.WrapMainRoute(example())

	g := e.Group("/example")
	g.GET("", example)
}
