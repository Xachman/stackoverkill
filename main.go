package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"zacharyironside.com/pages"
)

func main() {
	e := echo.New()
	e.GET("/", pages.IndexPage)
	e.GET("/blog/:slug", BlogPages)
	e.POST("/time", func (c echo.Context) error {
		return c.String(http.StatusOK, "the time")
	})
	e.Static("/static", "assets")
	e.Logger.Fatal(e.Start(":8080"))
}