package common


type Handler[In any, Out any] interface{
	Handle(arg In) (result Out, e error)
}