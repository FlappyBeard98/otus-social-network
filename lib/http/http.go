// Package http contains helper functions for http server
package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// HttpConfig contains configuration for http server
type HttpConfig struct {
	Port int `json:"port" env:"PORT" env-default:"1323"` // Port for http server
}

// GetEchoPort returns formatted port for echo http server
func (o *HttpConfig) GetEchoPort() string {
	return fmt.Sprintf(":%d", o.Port)
}

// StartHttpServer starts http.Echo http with graceful shutdown.
// Copied from https://echo.labstack.com/cookbook/graceful-shutdown/
func StartHttpServer(echo *echo.Echo, cfg HttpConfig) {
	// Start server
	go func() {
		port := cfg.GetEchoPort()
		if err := echo.Start(port); err != nil && err != http.ErrServerClosed {
			echo.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := echo.Shutdown(ctx); err != nil {
		echo.Logger.Fatal(err)
	}
}

// NewKeyMiddleware returns middleware for echo server that checks if request has correct key
func NewKeyMiddleware(accesskey string) echo.MiddlewareFunc {
	return middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == accesskey, nil
	})
}
