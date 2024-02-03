package pages

import (
	"rename-me/pages/example"
	"rename-me/routes"

	"github.com/labstack/echo/v4"
)

func AddAllRoutes(g *echo.Group) {
	routes.AddRoute("Example", "/example", g, example.Example)
}
