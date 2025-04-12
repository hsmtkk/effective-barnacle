package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to parse as int: %v", err)
	}
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/index.html")),
	}
	e := echo.New()
	e.Renderer = t
	e.GET("/", root)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

func root(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "World")
}
