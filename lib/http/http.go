package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func NewKeyMiddleware(accesskey string) echo.MiddlewareFunc {
	return middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == accesskey, nil
	})
}

type RequestFactory interface {
	CreateRequest(host string) (*http.Request, error)
}

func GetHttpResponse[T any](requestFactory RequestFactory, host string, setup func(*http.Request) *http.Request) (*T, error) {

	request, err := requestFactory.CreateRequest(host)
	if err != nil {
		return nil, err
	}

	request = setup(request)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(response.Body)

	defer func() { _ = response.Body.Close() }()
	if response.StatusCode > 299 {
		err = fmt.Errorf("%v, body: %s", *response, string(data))
	}

	if err != nil {
		return nil, err
	}

	result := new(T)
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func SetBasicAuth(username string, password string) func(*http.Request) *http.Request {
	return func(request *http.Request) *http.Request {
		request.SetBasicAuth(username, password)
		return request
	}
}
