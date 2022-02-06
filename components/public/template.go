package public

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type Template struct {
	Template *template.Template
}

type Deps struct {
	Template *Template
}

func (d *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return d.Template.ExecuteTemplate(w, name, data)
}

//func Index(c echo.Context) error {
//	return c.Render(200, "index", echo.Map{
//		"title": "Index",
//		"body":  "Hello, World!",
//	})
//}
