package db

import "context"

type Client interface {
	ExecuteQuery(context.Context, string) (any, error)

	StartTransaction(ctx context.Context) error
	Commit(ctx context.Context) error
	CloseTransaction(ctx context.Context) error
}
