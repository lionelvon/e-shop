package main

import (
	"github.com/CloudyKit/jet"
	"github.com/labstack/echo"
	"github.com/lionelvon/e-shop/handles"
	"github.com/lionelvon/e-shop/jetengine"
)

func main() {
	e := echo.New()
	e.Static("/public", "public")

	e.Renderer = &jetengine.Template{
		View: jet.NewHTMLSet("./views").SetDevelopmentMode(true),
	}

	e.GET("/", handles.Main)

	e.Logger.Fatal(e.Start(":1323"))
}
