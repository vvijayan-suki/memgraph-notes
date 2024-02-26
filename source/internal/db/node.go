package db

import (
	"github.com/vvijayan-suki/memgraph-notes/source/internal/db/constant/property"
)

type Node struct {
	labels     []string
	properties map[string]any
}

func (n *Node) Labels() []string {
	return n.labels
}

func (n *Node) FirstLabel() string {
	if len(n.labels) > 0 {
		return n.labels[0]
	}

	return ""
}

func (n *Node) Properties() map[string]any {
	return n.properties
}

func (n *Node) Property(key string) (any, bool) {
	value, ok := n.properties[key]

	return value, ok
}

func (n *Node) SecondaryID() (string, bool) {
	value, _ := n.Property(property.SecondaryID)

	secondaryID, ok := value.(string)

	return secondaryID, ok
}

func (n *Node) ID() (string, bool) {
	value, _ := n.Property(property.ID)

	id, ok := value.(string)

	return id, ok
}

func (n *Node) StartCursorPosition() (int, bool) {
	value, _ := n.Property(property.StartCursorPosition)

	startCursorPosition, ok := value.(int64)

	return int(startCursorPosition), ok
}

func (n *Node) EndCursorPosition() (int, bool) {
	value, _ := n.Property(property.EndCursorPosition)

	endCursorPosition, ok := value.(int64)

	return int(endCursorPosition), ok
}

func (n *Node) Location() (int, bool) {
	value, _ := n.Property(property.Location)

	location, ok := value.(int64)

	return int(location), ok
}

func NewNode(labels []string, properties map[string]any) *Node {
	var _labels []string

	for _, label := range labels {
		if len(label) > 0 {
			_labels = append(_labels, label)
		}
	}

	_properties := make(map[string]any)

	for key, value := range properties {
		if len(key) > 0 {
			_properties[key] = value
		}
	}

	return &Node{
		labels:     _labels,
		properties: _properties,
	}
}

func (n *Node) AddLabels(labels ...string) {
	for _, label := range labels {
		if len(label) > 0 {
			n.labels = append(n.labels, labels...)
		}
	}
}

func (n *Node) SetProperties(properties map[string]any) {
	for key, value := range properties {
		if len(key) > 0 {
			n.properties[key] = value
		}
	}
}
