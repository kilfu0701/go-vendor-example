package app

import (
	"fmt"
	"net/http"
	"helpers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func RunApp() {
	// test import local package
	result := helpers.Sum("123", "456")
	fmt.Printf("123 + 456 = %d\n", result)

	// test import vendor package
	port := "8081"
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	fmt.Printf("Start running ... localhost:%s\n", port)

	e.Run(standard.New(":" + port))
}
