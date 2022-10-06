package command

import "social-network/common"

type SaveProfileCommand struct {
}

type SaveProfileHandler = common.Handler[SaveProfileCommand, interface{}]

type saveProfileHandler struct {
}

func NewSaveProfileHandler() SaveProfileHandler {
	return &saveProfileHandler{}
}

func (h *saveProfileHandler) Handle(arg SaveProfileCommand) (interface{}, error) {
	return nil, nil
}
