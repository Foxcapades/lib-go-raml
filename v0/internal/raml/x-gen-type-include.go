package raml

import (
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
	"gopkg.in/yaml.v3"
)

// NewIncludeType returns a new internal implementation of
// the raml.IncludeType interface.
//
// Generated @ 2025-04-30T18:02:23.857064227-04:00
func NewIncludeType() *IncludeType {
	out := &IncludeType{}

	out.DataType = NewDataType(rmeta.TypeInclude, out)

	return out
}

// IncludeType is a generated internal implementation of
// the raml.IncludeType interface.
//
// Generated @ 2025-04-30T18:02:23.857064227-04:00
type IncludeType struct {
	*DataType
}

func (a IncludeType) marshal(out raml.AnyMap) error {
	out.Put(rmeta.KeyType, &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   "!include",
		Value: a.DataType.schema,
	})
	a.DataType.hasExtra.out(out)

	return nil
}
func (a *IncludeType) assign(key, val *yaml.Node) (err error) {
	switch key.Value {
	case rmeta.KeyType, rmeta.KeySchema:
		a.DataType.schema = val.Value
	}
	return nil
}
