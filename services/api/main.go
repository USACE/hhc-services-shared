package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/hhc-services-shared/services/api/handler"

	_config "github.com/hhc-services-shared/services/api/config"
	_middleware "github.com/hhc-services-shared/services/api/middleware"
)

func main() {
	// parse configuration from environment variables
	// Environment Variable Config
	var cfg _config.Config
	err := cfg.ParseEnvVars()
	if err != nil {
		log.Fatalf("fail on parse env vars %s", err)
	}

	// Populate AWS Config with the values from the configurations
	cfg.AwsConfig, err = config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("failed to load AWS configuration %s", err)
	}

	e := echo.New()

	// Middleware for *echo.Echo
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())

	// Middleware to serve static content from s3
	// Routes use a base url and those have to be ignored here

	e.Use(_middleware.S3StaticWithConfig(_middleware.S3StaticConfig{
		Skipper: func(c echo.Context) bool {
			path := c.Request().URL.Path
			return strings.HasPrefix(path, "/api")
		},
		AwsConfig:    cfg.AwsConfig,
		Bucket:       cfg.S3Bucket,
		UsePathStyle: cfg.UsePathStyle,
		Prefix:       cfg.S3PrefixStatic,
	}))

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
	pgStore, err := _config.NewPgStore(&cfg)
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
