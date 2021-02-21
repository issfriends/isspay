package server

import (
	"context"
	"fmt"
)

// Config server configuration
type Config struct {
	Host string
	Port int
}

// Addr concate host and port
func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// Server server interface
type Server interface {
	Run() error
	Shutdown(ctx context.Context) error
}
