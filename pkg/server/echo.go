package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

// NewEcho new echo server engine
func NewEcho(c *Config, routes ...func(e *echo.Echo)) (Server, error) {
	e := echo.New()
	e.Server.Addr = c.Addr()

	for _, route := range routes {
		route(e)
	}

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK!")
	})

	return &echoAdapter{Echo: e}, nil
}

type echoAdapter struct {
	*echo.Echo
}

func (server *echoAdapter) Run() error {
	return server.StartServer(server.Server)
}

func (server *echoAdapter) Shutdown(ctx context.Context) error {
	return server.Echo.Shutdown(ctx)
}
