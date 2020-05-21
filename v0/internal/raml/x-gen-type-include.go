package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
)

// NewIncludeType returns a new internal implementation of
// the raml.IncludeType interface.
//
// Generated @ 2020-05-20T20:54:26.833516016-04:00
func NewIncludeType() *IncludeType {
	out := &IncludeType{}

	out.DataType = NewDataType(rmeta.TypeInclude, out)

	return out
}

// IncludeType is a generated internal implementation of
// the raml.IncludeType interface.
//
// Generated @ 2020-05-20T20:54:26.833516016-04:00
type IncludeType struct {
	*DataType
}

func (a IncludeType) marshal(out raml.AnyMap) error {
	out.Put(rmeta.KeyType, "!include "+a.DataType.schema)
	a.DataType.hasExtra.out(out)

	return nil
}
