package xyml

import "gopkg.in/yaml.v3"

// IsBool returns whether the given YAML node represents a boolean value.
func IsBool(y *yaml.Node) bool {
	return y.Tag == Bool
}

// IsFloat returns whether the given YAML node represents a float value.
func IsFloat(y *yaml.Node) bool {
	return y.Tag == Float
}

// IsInt returns whether the given YAML node represents an int value.
func IsInt(y *yaml.Node) bool {
	return y.Tag == Int
}

// IsList returns whether the given YAML node represents a sequence value.
func IsList(y *yaml.Node) bool {
	return y.Kind == yaml.SequenceNode
}

// IsMap returns whether the given YAML node represents a mapping value.
func IsMap(y *yaml.Node) bool {
	return y.Kind == yaml.MappingNode
}

// IsString returns whether the given YAML node represents a string value.
func IsString(y *yaml.Node) bool {
	return y.Tag == String
}
