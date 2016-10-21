package app

import (
	"fmt"
	"net/http"

	// local packages
	"helpers"
	"app/config"

	// vendor packages
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func RunApp() {
	result := helpers.Sum("123", "456")
	fmt.Printf("123 + 456 = %d\n", result)

	cfg := config.Load()
	port, _ := cfg["port"].(string)
	host, _ := cfg["host"].(string)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	fmt.Printf("Start running ... %s:%s\n", host, port)

	e.Run(standard.New(host + ":" + port))
}
