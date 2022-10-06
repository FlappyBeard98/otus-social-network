package query

import "social-network/common"

type ProfilesByFilterQuery struct {
}

type ProfilesByFilterHandler = common.Handler[ProfilesByFilterQuery, interface{}]

type profilesByFilterHandler struct {
}

func NewProfilesByFilterHandler() ProfilesByFilterHandler {
	return &profilesByFilterHandler{}
}

func (h *profilesByFilterHandler) Handle(arg ProfilesByFilterQuery) (interface{}, error) {
	return nil, nil
}
