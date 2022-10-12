package http

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// StartHttpServer starts http.Echo http with graceful shutdown.
// Copied from https://echo.labstack.com/cookbook/graceful-shutdown/
func StartHttpServer(echo *echo.Echo, echoPort string) {
	// Start server
	go func() {
		if err := echo.Start(echoPort); err != nil && err != http.ErrServerClosed {
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
