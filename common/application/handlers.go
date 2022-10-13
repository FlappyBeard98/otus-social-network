package application

import "context"

// Handler is common handler for any operations
type Handler[In any, Out any] interface {
	Handle(ctx context.Context, arg In) (result Out, e error)
}
