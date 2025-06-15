package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)


func BlogPages(c echo.Context) error {
	slug := c.Param("slug")
	headerContent, err := os.ReadFile("./pages/templates/header.html")
	if (err != nil) {
		fmt.Println(err.Error())
		return err
	}
	mainContent, err := os.ReadFile("./blogs/"+slug+".html")
	if (err != nil) {
		fmt.Println(err.Error())
		return err
	}
	footerContent, err := os.ReadFile("./pages/templates/footer.html")
	if (err != nil) {
		fmt.Println(err.Error())
		return err
	}
	return c.HTML(http.StatusOK, string(headerContent)+string(mainContent)+string(footerContent))
}
