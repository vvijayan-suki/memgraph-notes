package query

import "context"

type Request interface {
	Validate() error
	Build() (string, error)
	EvaluateResponse(context.Context, any) (any, error)
}
