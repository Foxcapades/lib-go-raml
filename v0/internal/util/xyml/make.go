package xyml

import (
	"gopkg.in/yaml.v3"
	"strconv"
)

func StringNode(v string) *yaml.Node {
	return &yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: v}
}

func IntNode(v int64) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   "!!int",
		Value: strconv.FormatInt(v, 10),
	}
}

func FloatNode(v float64) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   "!!float",
		Value: strconv.FormatFloat(v, 'f', -1, 64),
	}
}

func MapNode(size int) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.MappingNode,
		Tag:     "!!map",
		Content: make([]*yaml.Node, size),
	}
}

func SequenceNode(size int) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.SequenceNode,
		Tag:     "!!seq",
		Content: make([]*yaml.Node, size),
	}
}
