package query

import (
	"social-network/common"
)

type HelloQuery struct {
	Name string `param:"name"`
}

type HelloHandler = common.Handler[HelloQuery, string]

type helloHandler struct {
}

func NewHelloHandler() HelloHandler {
	return &helloHandler{}
}

func (h *helloHandler) Handle(arg HelloQuery) (string, error) {
	return "hello " + arg.Name, nil
}
