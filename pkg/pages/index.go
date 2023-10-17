package pages

import (
	"github.com/labstack/echo/v4"
)

type Page struct {
	Error string
}

func IndexPage(c echo.Context) error {
	return c.Render(200, "index.html", Page{
		Error: "",
	})
}
