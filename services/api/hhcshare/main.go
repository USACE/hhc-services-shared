package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	// Middleware for *echo.Echo
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())

	// Groups
	grp_api:=e.Group("api")

	// ~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~
	// basic routes
	// ~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~
	// Public Health Check Route
	grp_api.GET("/health",
		func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]any{"status": "healthy"})
		})

	// API Routes
	grp_api.GET("/routes",
		func(c echo.Context) error { return c.JSON(http.StatusOK, e.Routes()) },
	)

	// ~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~
	// routes
	// ~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~

	// ~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~
	// start server
	// ~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~
	port, hasPort := os.LookupEnv("ADDRESS_PORT")
	if !hasPort || port == "" {
		port = "8080"
	}

	addressPort := fmt.Sprintf(":%s", port)

	e.Logger.Printf("Starting main server on :%s...", addressPort)
	e.Logger.Fatal(e.Start(addressPort))
}