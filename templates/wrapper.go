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

		// A context can be defined here that will give global access to all rendered content
		ctx := context.Background()

		if len(hxr) > 0 && hxr[0] == "true" {
			Body(route, tabContent).Render(ctx, c.Response())
		} else {
			BodyWithBase(route, tabContent).Render(ctx, c.Response())
		}
		return nil
	}
}
