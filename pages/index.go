package pages

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func IndexPage(c echo.Context) error {
	headerContent, err := os.ReadFile("./pages/templates/header.html")
	if (err != nil) {
		fmt.Println(err.Error())
		return err
	}
	mainContent, err := os.ReadFile("./pages/templates/index.html")
	if (err != nil) {
		fmt.Println(err.Error())
		return err
	}
	footerContent, err := os.ReadFile("./pages/templates/footer.html")
	if (err != nil) {
		fmt.Println(err.Error())
		return err
	}
	html := string(headerContent)+string(mainContent)+string(footerContent)
	return c.HTML(http.StatusOK, html)
}