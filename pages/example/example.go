package example

import (
	"rename-me/templates"

	"github.com/labstack/echo/v4"
)

var Example echo.HandlerFunc = templates.WrapMainRoute(example())
