package transaction

import (
	"context"
	"sync"

	"github.com/vvijayan-suki/memgraph-notes/source/internal/db"
	"github.com/vvijayan-suki/memgraph-notes/source/internal/query"
)

type Manager struct {
	client db.Client
	mutex  sync.Mutex
}

func NewManager(client db.Client) *Manager {
	return &Manager{
		client: client,
		mutex:  sync.Mutex{},
	}
}

func (m *Manager) Start(ctx context.Context) error {
	return m.client.StartTransaction(ctx)
}

func (m *Manager) Run(ctx context.Context, req query.Request) (any, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	query, err := req.Build()
	if err != nil {
		return nil, err
	}

	// making the below code as critical section as memgraph sessions are not thread-safe
	m.mutex.Lock()
	defer m.mutex.Unlock()

	response, err := m.client.ExecuteQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	return req.EvaluateResponse(ctx, response)
}

func (m *Manager) Commit(ctx context.Context) error {
	return m.client.Commit(ctx)
}

func (m *Manager) Close(ctx context.Context) error {
	return m.client.CloseTransaction(ctx)
}
