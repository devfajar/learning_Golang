package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	//"github.com/labstack/echo/v4/middleware"
)

type M map[string]interface{}

const CSRFTokenHeader = "X-Csrf-Token"
const CSRFKey = "csrf_token"

func main() {
	tmpl := template.Must(template.ParseGlob("./*.html"))

	e := echo.New()

	// ...

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:" + CSRFTokenHeader,
		ContextKey:  CSRFKey,
	}))

	e.GET("/index", func(c echo.Context) error {
		data := make(M)
		data[CSRFKey] = c.Get(CSRFKey)
		return tmpl.Execute(c.Response(), data)
	})

	e.POST("/sayhello", func(c echo.Context) error {
		data := make(M)
		if err := c.Bind(&data); err != nil {
			return err
		}

		message := fmt.Sprintf("hello %s", data["name"])
		return c.JSON(http.StatusOK, message)
	})

	e.Logger.Fatal(e.Start(":9000"))
}
