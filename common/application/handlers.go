package application

import "context"

type Handler[In any, Out any] interface {
	Handle(ctx context.Context, arg In) (result Out, e error)
}
