package main

import (
	"./db"
	"./handles"
	"./render"
	"github.com/CloudyKit/jet"
	"github.com/labstack/echo"
)

func main() {
	db.InitMySQL()
	e := echo.New()
	e.Static("/public", "public")

	e.Renderer = &render.Template{
		View: jet.NewHTMLSet("./views").SetDevelopmentMode(true),
	}

	e.GET("/", handles.Main)

	e.Logger.Fatal(e.Start(":1323"))
}
