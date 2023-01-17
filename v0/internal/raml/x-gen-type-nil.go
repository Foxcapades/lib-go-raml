package raml

import (
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
)

// NewNilType returns a new internal implementation of
// the raml.NilType interface.
//
// Generated @ 2023-01-17T10:02:54.294844187-05:00
func NewNilType() *NilType {
	out := &NilType{}

	out.DataType = NewDataType(rmeta.TypeNil, out)

	return out
}

// NilType is a generated internal implementation of
// the raml.NilType interface.
//
// Generated @ 2023-01-17T10:02:54.294844187-05:00
type NilType struct {
	*DataType
}
