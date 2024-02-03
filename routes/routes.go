package routes

import "github.com/labstack/echo/v4"

type RouteData struct {
	Name string
	Url  string
}

var AppRoutes = []RouteData{}

func AddRoute(name string, url string, g *echo.Group, handler echo.HandlerFunc) {
	g.GET(url, handler)
	route := RouteData{
		Name: name,
		Url:  url,
	}
	AppRoutes = append(AppRoutes, route)
}
