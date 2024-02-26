package memgraph

import (
	"context"

	"github.com/LearningMotors/platform/service/ctxlog"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Client struct {
	session neo4j.SessionWithContext
	txn     neo4j.ExplicitTransaction
}

func Get() *Client {
	return &Client{}
}

func (c *Client) StartTransaction(ctx context.Context) error {
	if c.session == nil {
		session := driver.NewSession(ctx, neo4j.SessionConfig{
			BookmarkManager: driver.ExecuteQueryBookmarkManager(),
		})

		c.session = session
	}

	txn, err := c.session.BeginTransaction(ctx)
	if err != nil {
		return err
	}

	c.txn = txn

	return nil
}

func (c *Client) ExecuteQuery(ctx context.Context, query string) (any, error) {
	sugar := ctxlog.Sugar(ctx)

	sugar.Infow("query to be executed", "query", query)

	if c.session != nil && c.txn != nil {
		result, err := c.txn.Run(ctx, query, nil)

		sugar.Infow("result of transaction query", "query", query, "result", result, "error", err)

		return result, err
	}

	result, err := neo4j.ExecuteQuery(ctx, driver, query, nil, neo4j.EagerResultTransformer)

	sugar.Infow("result of query", "query", query, "result", result, "error", err)

	return result, err
}

func (c *Client) Commit(ctx context.Context) error {
	sugar := ctxlog.Sugar(ctx)

	sugar.Infow("committing transaction")

	err := c.txn.Commit(ctx)
	if err != nil {
		sugar.Errorw("error committing transaction", "error", err)
		return err
	}

	sugar.Infow("transaction committed")

	return nil
}

func (c *Client) Rollback(ctx context.Context) error {
	sugar := ctxlog.Sugar(ctx)

	sugar.Infow("rolling back transaction")

	err := c.txn.Rollback(ctx)
	if err != nil {
		sugar.Errorw("error rolling back transaction", "error", err)
		return err
	}

	sugar.Infow("transaction rolled back")

	return nil
}

func (c *Client) CloseTransaction(ctx context.Context) error {
	sugar := ctxlog.Sugar(ctx)

	sugar.Infow("closing transaction")

	err := c.txn.Close(ctx)
	if err != nil {
		sugar.Errorw("error closing transaction", "error", err)
	}

	c.txn = nil

	sugar.Infow("transaction closed")

	sugar.Infow("closing session")

	err = c.session.Close(ctx)
	if err != nil {
		sugar.Errorw("error closing session", "error", err)
		return err
	}

	c.session = nil

	sugar.Infow("session closed")

	return nil
}
