package getnode

import (
	"context"
	"errors"
	"fmt"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"

	"github.com/vvijayan-suki/memgraph-notes/source/internal/db"
	query "github.com/vvijayan-suki/memgraph-notes/source/internal/query/request"
)

// Match
// node `n` having specified secondary_id
const parameterizedQuery = "MATCH (n %s %s ) RETURN DISTINCT n"

type GetNode struct {
	Labels     []string
	Properties map[string]any
}

type Option func(*GetNode)

func WithLabel(label string) Option {
	return func(r *GetNode) {
		if len(label) > 0 {
			r.Labels = append(r.Labels, label)
		}
	}
}

func WithProperty(key string, value any) Option {
	return func(r *GetNode) {
		if len(key) > 0 {
			r.Properties[key] = value
		}
	}
}

func New(label string, opts ...Option) *GetNode {
	properties := map[string]any{}

	req := &GetNode{
		Labels:     []string{label},
		Properties: properties,
	}

	for _, opt := range opts {
		opt(req)
	}

	return req
}

func (r *GetNode) Validate() error {
	var _labels []string

	for _, label := range r.Labels {
		if len(label) > 0 {
			_labels = append(_labels, label)
		}
	}

	r.Labels = _labels

	return validation.ValidateStruct(r,
		validation.Field(&r.Labels, validation.Required, validation.Length(1, 0)),
	)
}

func (r *GetNode) Build() (string, error) {
	var propertiesList []string
	for key, value := range r.Properties {
		property, err := query.ToGraphData(value)
		if err != nil {
			return "", err
		}

		propertiesList = append(propertiesList, fmt.Sprintf("%s: %s", key, property))
	}

	labels := ""
	if len(r.Labels) > 0 {
		labels = fmt.Sprintf(":%s", strings.Join(r.Labels, ":"))
	}

	properties := ""
	if len(propertiesList) > 0 {
		properties = fmt.Sprintf("{ %s }", strings.Join(propertiesList, ", "))
	}

	return fmt.Sprintf(parameterizedQuery, labels, properties), nil
}

func (r *GetNode) EvaluateResponse(ctx context.Context, response any) (any, error) {
	switch response := response.(type) {
	case *neo4j.EagerResult:
		return evaluateEagerResult(response)
	case neo4j.ResultWithContext:
		return evaluateResultWithContext(ctx, response)
	default:
		return nil, fmt.Errorf("unsupported response type: %v", response)
	}
}

func evaluateEagerResult(response *neo4j.EagerResult) ([]*db.Node, error) {
	nodes := make([]*db.Node, 0)

	for _, record := range response.Records {
		graphNode, ok := record.AsMap()["n"].(neo4j.Node)
		if !ok {
			return nil, errors.New("invalid conversion to neo4j node")
		}

		node := db.NewNode(graphNode.Labels, graphNode.GetProperties())

		nodes = append(nodes, node)
	}

	return nodes, nil
}

func evaluateResultWithContext(ctx context.Context, response neo4j.ResultWithContext) ([]*db.Node, error) {
	nodes := make([]*db.Node, 0)

	for response.Next(ctx) {
		record := response.Record()

		graphNode, ok := record.AsMap()["n"].(neo4j.Node)
		if !ok {
			return nil, errors.New("invalid conversion to neo4j node")
		}

		node := db.NewNode(graphNode.Labels, graphNode.GetProperties())

		nodes = append(nodes, node)
	}

	return nodes, nil
}
