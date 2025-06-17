package main

import (
	"fmt"
	"net/http"
	"os"
	apitools "stackoverkill/api-tools"

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
	tool := ToolFactory(slug)
	return tool.Handle(c)
}

func ToolFactory(name string) Tool {
	mpg := &apitools.MemorablePasswordGenerator{}
	switch name {
	case mpg.Name():
        return &apitools.MemorablePasswordGenerator{}
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
