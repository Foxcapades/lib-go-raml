package xyml

import (
	"gopkg.in/yaml.v3"
	"strconv"
)

func BoolNode(v bool) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   Bool,
		Value: strconv.FormatBool(v),
	}
}

func FloatNode(v float64) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   Float,
		Value: strconv.FormatFloat(v, 'f', -1, 64),
	}
}

func IntNode(v int64) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   Int,
		Value: strconv.FormatInt(v, 10),
	}
}

func MapNode(size int) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.MappingNode,
		Tag:     Map,
		Content: make([]*yaml.Node, 0, size*MapSkip),
	}
}

func SequenceNode(size int) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.SequenceNode,
		Tag:     Sequence,
		Content: make([]*yaml.Node, 0, size),
	}
}

func StringNode(v string) *yaml.Node {
	return &yaml.Node{Kind: yaml.ScalarNode, Tag: String, Value: v}
}
