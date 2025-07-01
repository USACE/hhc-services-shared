package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"hhcshare/config"
	"hhcshare/handler"
)

func main() {
	// parse configuration from environment variables
	// Environment Variable Config
	var cfg config.Config
	if err := cfg.ParseEnvVars(); err != nil {
		log.Fatalf("fail on parse env vars %s", err)
	}

	// LoadDefaultAwsConfig is a wrapper for LoadDefaultConfig to handle using MINIO locally
	if err := cfg.LoadDefaultAwsConfig(); err != nil {
		log.Fatalf("failed to load AWS configuration %s", err)
	}

	e := echo.New()

	// Middleware for *echo.Echo
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())

	// Groups
	grp_api := e.Group("/api")
	public := grp_api.Group("")

	// private routes with auth
	// adminAuth := _middleware.JwtAuthenticateConfig{
	// 	Skipper:        _middleware.DefaultJwtAuthSkipper,
	// 	PublicKeyRaw:   cfg.AuthPublicKey,
	// 	ApplicationKey: cfg.ApplicationKey,
	// }

	// private := grp_api.Group("", adminAuth.AuthParseWithClaims(),
	// 	_middleware.ResourceAccessWithConfig(_middleware.ResourceAccessConfig{
	// 		Roles:         cfg.ResourceAccessRoles,
	// 		RoleSeperator: ".",
	// 	}))

	// ~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~
	// basic routes
	// ~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~
	// Public Health Check Route
	public.GET("/health",
		func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]any{"status": "healthy"})
		})

	// API Routes
	public.GET("/routes",
		func(c echo.Context) error { return c.JSON(http.StatusOK, e.Routes()) },
	)

	// Database
	pgStore, err := config.NewPgStore(&cfg)
	if err != nil {
		e.Logger.Fatal(err.Error())
	}
	pgH := handler.HandlerStore{Connection: pgStore.Connection, Config: &cfg}

	// ~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~
	// routes
	// ~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~
	// offices
	public.GET("/offices", pgH.ListOffices)
	public.GET("/offices/:office/geometry", pgH.OfficeGeometry)

	// ~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~
	// start server
	// ~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~
	port, hasPort := os.LookupEnv("API_PORT")
	if !hasPort || port == "" {
		port = "8080"
	}

	addressPort := fmt.Sprintf(":%s", port)

	e.Logger.Printf("Starting main server on :%s...", addressPort)
	e.Logger.Fatal(e.Start(addressPort))
}
