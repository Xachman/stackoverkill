package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func BlogPages(c echo.Context) error {
	slug := c.Param("slug")
	headerContent, err := os.ReadFile("./web/templates/header.html")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	mainContent, err := os.ReadFile("./blogs/" + slug + ".html")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	footerContent, err := os.ReadFile("./web/templates/footer.html")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return c.HTML(http.StatusOK, string(headerContent)+string(mainContent)+string(footerContent))
}

func BlogPage(c echo.Context) error {
	headerContent, err := os.ReadFile("./web/templates/header.html")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	mainContent, err := os.ReadFile("./web/templates/blog.html")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	footerContent, err := os.ReadFile("./web/templates/footer.html")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	html := string(headerContent) + string(mainContent) + string(footerContent)
	return c.HTML(http.StatusOK, html)
}
