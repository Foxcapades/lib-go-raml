package raml

import (
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
	"gopkg.in/yaml.v3"
)

// NewIncludeType returns a new internal implementation of
// the raml.IncludeType interface.
//
// Generated @ 2020-07-06T12:49:37.941034901-04:00
func NewIncludeType() *IncludeType {
	out := &IncludeType{}

	out.DataType = NewDataType(rmeta.TypeInclude, out)

	return out
}

// IncludeType is a generated internal implementation of
// the raml.IncludeType interface.
//
// Generated @ 2020-07-06T12:49:37.941034901-04:00
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
