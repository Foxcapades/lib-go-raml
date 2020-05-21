package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
)

// NewAnyType returns a new internal implementation of
// the raml.AnyType interface.
//
// Generated @ 2020-05-20T18:40:13.095690448-04:00
func NewAnyType() *AnyType {
	out := &AnyType{}

	out.DataType = NewDataType(rmeta.TypeAny, out)

	return out
}

// AnyType is a generated internal implementation of
// the raml.AnyType interface.
//
// Generated @ 2020-05-20T18:40:13.095690448-04:00
type AnyType struct {
	*DataType
}