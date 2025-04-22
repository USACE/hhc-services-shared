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

	// ~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~
	// basic routes
	// ~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~
	// Public Health Check Route
	e.GET("/health",
		func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]any{"status": "healthy"})
		})

	// API Routes
	e.GET("/routes",
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