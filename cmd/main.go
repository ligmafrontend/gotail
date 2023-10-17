package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"ligmafrontend/gotail/pkg/db"
	"ligmafrontend/gotail/pkg/pages"
	"log"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	db.Init()

	templates, err := template.New("").ParseGlob("templates/*.html")

	if err != nil {
		log.Fatal("Error loading templates: ", err)
	}

	e := echo.New()
	e.Renderer = &TemplateRenderer{
		templates: templates,
	}

	e.Use(middleware.Logger())
	e.Static("/static", "static")
	e.Static("/css", "css")
	e.Static("/htmx", "htmx")

	e.GET("/", pages.IndexPage)

	e.Logger.Fatal(e.Start(":42069"))
}
