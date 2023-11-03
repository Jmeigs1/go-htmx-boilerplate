package templates

import (
	"context"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func WrapMainRoute(tabContent templ.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
		hxr := c.Request().Header["Hx-Request"]
		route := c.Path()

		if len(hxr) > 0 && hxr[0] == "true" {
			Body(route, tabContent).Render(context.Background(), c.Response())
		} else {
			Index(route, tabContent).Render(context.Background(), c.Response())

		}
		return nil
	}
}
