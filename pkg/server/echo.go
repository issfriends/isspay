package server

import "github.com/labstack/echo/v4"

type Config struct {
	Host string
	Port int
}

func New(c *Config) (*echo.Echo, error) {
	e := echo.New()

	return e, nil
}
