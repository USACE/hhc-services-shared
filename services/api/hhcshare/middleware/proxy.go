package middleware

import (
	"net/url"
	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomProxyTarget struct {
	URL          *url.URL
	Rewrite      map[string]string
	RegexRewrite map[*regexp.Regexp]string
}

func (p CustomProxyTarget) ProxyWithCustomConfig() echo.MiddlewareFunc {
	return middleware.ProxyWithConfig(
		middleware.ProxyConfig{
			Rewrite:      p.Rewrite,
			RegexRewrite: p.RegexRewrite,
			Balancer: middleware.NewRoundRobinBalancer(
				[]*middleware.ProxyTarget{
					{
						URL: p.URL,
					},
				}),
		})
}
