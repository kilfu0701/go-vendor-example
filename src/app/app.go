package app

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func RunApp() {
	port := "8081"

	fmt.Printf("Start running ... localhost:%s", port)

    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

    e.Run(standard.New(":" + port))
}
