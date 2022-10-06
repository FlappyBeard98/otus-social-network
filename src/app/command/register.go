package command

import "social-network/common"

type RegisterCommand struct {
}

type RegisterHandler = common.Handler[RegisterCommand, interface{}]

type registerHandler struct {
}

func NewRegisterHandler() RegisterHandler {
	return &registerHandler{}
}

func (h *registerHandler) Handle(arg RegisterCommand) (interface{}, error) {
	return nil, nil
}
