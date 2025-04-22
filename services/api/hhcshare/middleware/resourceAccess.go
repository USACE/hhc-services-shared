package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type (
	// ResourceAccessConfig struct defining needed fields to validate and authorize
	ResourceAccessConfig struct {
		// Skipper defines a function to skip middleware.
		// Returning true skips processing the middleware.
		Skipper func(c echo.Context) bool

		// Roles is the list of roles to authorize
		Roles []string

		// Role seperator.
		// Optional.  Default value ""
		RoleSeperator string

		// Typically the office of the user defined in routes
		Scope string

		// ScopeFromContext func to get the scope
		// Typically from the route and is most likely an office symbol
		ScopeFromContext func(c echo.Context) error

		// Context key
		// Optional.  Default value "user"
		ContextKey string

		// Claims.
		Claims AuthorizeCustomClaims
	}
)

var (
	DefaultResourceAccessConfig = ResourceAccessConfig{
		Skipper: DefaultResourceAcessSkipper,
	}
)

// DefaultResourceAcessSkipper returns false which processes the middleware.
func DefaultResourceAcessSkipper(echo.Context) bool {
	return false
}

// DefaultScopeFromContext
func (r *ResourceAccessConfig) DefaultScopeFromContext(c echo.Context) error {
	paramNames := c.ParamNames()
	for _, name := range paramNames {
		if name == r.ContextKey {
			r.Scope = c.Param(name)
		}
	}
	return nil
}

// AuthResourceAccessWithConfig
func ResourceAccessWithConfig(accessConfig ResourceAccessConfig) echo.MiddlewareFunc {
	if accessConfig.Skipper == nil {
		accessConfig.Skipper = DefaultResourceAccessConfig.Skipper
	}

	if accessConfig.ScopeFromContext == nil {
		accessConfig.ScopeFromContext = func(c echo.Context) error { return nil }
	}

	if accessConfig.ContextKey == "" {
		accessConfig.ContextKey = "user"
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			// skip if true
			if accessConfig.Skipper(c) {
				return next(c)
			}

			// define Scope from context func
			if err = accessConfig.ScopeFromContext(c); err != nil {
				return
			}

			user := c.Get(accessConfig.ContextKey).(*jwt.Token)
			
			claims, ok := user.Claims.(*AuthorizeCustomClaims)
			if !ok {
				return fmt.Errorf("error cast claims as jwt.MapClaims")
			}

			resource_access := claims.ResourceAccess
			// get the resources for the app and error if app not in map
			app_roles, ok := resource_access[claims.AuthrorizedParty].(map[string]interface{})
			if !ok {
				return fmt.Errorf("failed to cast resource_access[app] as map[string]interface{}")
			}

			tokenRoles, ok := app_roles["roles"].([]interface{})
			if !ok {
				return fmt.Errorf("failed to cast app roles")
			}

			b := []bool{false}
			for _, tokenRole := range tokenRoles {
				tokenRoleString := tokenRole.(string)
				for _, authRole := range accessConfig.Roles {
					if accessConfig.RoleSeperator != "" {
						scopeRole := strings.Split(tokenRoleString, accessConfig.RoleSeperator)
						scope := scopeRole[0]
						role := scopeRole[1]
						if (strings.EqualFold(accessConfig.Scope, scope) || scope == "application") && strings.EqualFold(authRole, role) {
							b = append(b, true)
						}
					}
					if strings.EqualFold(tokenRoleString, authRole) {
						b = append(b, true)
					}
				}
			}
			// find a true and return next(c)
			// else return error
			for _, v := range b {
				if v {
					return next(c)
				}
			}

			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "resource access not allowed with current roles"})
		}
	}
}
