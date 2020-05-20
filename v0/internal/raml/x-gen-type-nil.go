package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
)

// NewNilType returns a new internal implementation of
// the raml.NilType interface.
//
// Generated @ 2020-05-20T18:40:13.095690448-04:00
func NewNilType() *NilType {
	out := &NilType{}

	out.DataType = NewDataType(rmeta.TypeNil, out)

	return out
}

// NilType is a generated internal implementation of
// the raml.NilType interface.
//
// Generated @ 2020-05-20T18:40:13.095690448-04:00
type NilType struct {
	*DataType
}
