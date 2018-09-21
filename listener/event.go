package listener

import "context"

type function func(param interface{}, ctx context.Context)

// Dispatch ...
func Dispatch(event function, param interface{}, ctx context.Context) {
	//go
	event(param, ctx)
}
