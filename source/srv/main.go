package main

import (
	"context"
	"fmt"

	"github.com/vvijayan-suki/memgraph-notes/source/internal/db"
	"github.com/vvijayan-suki/memgraph-notes/source/internal/db/constant/label"
	"github.com/vvijayan-suki/memgraph-notes/source/internal/db/memgraph"
	"github.com/vvijayan-suki/memgraph-notes/source/internal/query/request/getnode"
	"github.com/vvijayan-suki/memgraph-notes/source/internal/transaction"
)

func main() {
	fmt.Println(getMemgraphNoteIDs())
}

func getMemgraphNoteIDs() []string {
	memgraph.New().Initialize()

	request := getnode.New(label.Note)

	transactionManager := transaction.NewManager(memgraph.Get())

	response, err := transactionManager.Run(context.Background(), request)
	if err != nil {
		panic(err)
	}

	nodes, ok := response.([]*db.Node)
	if !ok {
		panic("cannot convert response to nodes")
	}

	ids := make([]string, len(nodes))

	for _, node := range nodes {
		id, ok := node.ID()
		if !ok {
			panic("cannot get id from node")
		}

		ids = append(ids, id)
	}

	return ids
}
