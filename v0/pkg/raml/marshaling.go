package raml

import (
	"gopkg.in/yaml.v3"
)

type Unmarshaler interface {
	UnmarshalRAML(value *yaml.Node) error
}

type Marshaler interface {
	MarshalRAML(out AnyMap) (simple bool, err error)
}
