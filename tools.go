package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type Tool interface {
	Name() string
	Handle(c echo.Context) error
}

func ToolPages(c echo.Context) error {
	slug := c.Param("slug")
	headerContent, err := os.ReadFile("./web/templates/header.html")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	mainContent, err := os.ReadFile("./web/templates/tool/" + slug + ".html")
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

func ToolApi(c echo.Context) error {
	slug := c.Param("slug")
	apiContent, err := os.ReadFile("./tool/" + slug + ".json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return c.JSON(http.StatusOK, apiContent)
}

func ToolFactory(name string) Tool {
	switch name {
	default:
		return nil
	}
}

func ToolPage(c echo.Context) error {
	headerContent, err := os.ReadFile("./web/templates/header.html")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	mainContent, err := os.ReadFile("./web/templates/tool.html")
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
