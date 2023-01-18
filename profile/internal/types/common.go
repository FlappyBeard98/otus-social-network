package types

import "errors"

var (
	ErrInvalidOperation = errors.New("invalid operation")
	ErrInvalidInput     = errors.New("invalid input")
	ErrInternal         = errors.New("internal error")
)
