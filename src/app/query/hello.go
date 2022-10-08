package query

import (
	"context"
	"social-network/common/application"
)

type HelloQuery struct {
	Name string `param:"name"`
}

type HelloHandler = application.Handler[HelloQuery, string]

type helloHandler struct {
}

func NewHelloHandler() HelloHandler {
	return &helloHandler{}
}

func (h *helloHandler) Handle(_ context.Context, arg HelloQuery) (string, error) {
	return "hello " + arg.Name, nil
}
