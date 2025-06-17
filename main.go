package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", IndexPage)
	e.GET("/blog", BlogPage)
	e.GET("/tool", ToolPage)
	e.GET("/blog/:slug", BlogPages)
	e.GET("/tool/:slug", ToolPages)
	e.GET("/api/tool/:slug", ToolApi)
	e.POST("/time", func(c echo.Context) error {
		return c.String(http.StatusOK, "the time")
	})
	e.Static("/static", "assets")
	e.Logger.Fatal(e.Start(":8080"))
}
