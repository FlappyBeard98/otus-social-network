package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

type HttpConfig struct {
	Port int `json:"port" env:"PORT" env-default:"1323"`
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
