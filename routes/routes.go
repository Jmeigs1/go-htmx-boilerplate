package routes

type RouteData struct {
	Name string
	Url  string
}

var AppRoutes = []RouteData{
	{
		Name: "Example",
		Url:  "/example",
	},
	{
		Name: "Page 1",
		Url:  "/1",
	},
	{
		Name: "Page 2",
		Url:  "/2",
	},
	{
		Name: "Page 3",
		Url:  "/3",
	},
}
