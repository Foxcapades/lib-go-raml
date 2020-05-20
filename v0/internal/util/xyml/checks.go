package xyml

import "gopkg.in/yaml.v3"

func IsBool(y *yaml.Node) bool {
	return y.Tag == Bool
}

func IsFloat(y *yaml.Node) bool {
	return y.Tag == Float
}

func IsInt(y *yaml.Node) bool {
	return y.Tag == Int
}

func IsList(y *yaml.Node) bool {
	return y.Kind == yaml.SequenceNode
}

func IsMap(y *yaml.Node) bool {
	return y.Kind == yaml.MappingNode
}

func IsString(y *yaml.Node) bool {
	return y.Tag == String
}
