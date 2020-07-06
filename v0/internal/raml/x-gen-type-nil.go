package raml

import (
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
)

// NewNilType returns a new internal implementation of
// the raml.NilType interface.
//
// Generated @ 2020-07-06T12:49:37.941034901-04:00
func NewNilType() *NilType {
	out := &NilType{}

	out.DataType = NewDataType(rmeta.TypeNil, out)

	return out
}

// NilType is a generated internal implementation of
// the raml.NilType interface.
//
// Generated @ 2020-07-06T12:49:37.941034901-04:00
type NilType struct {
	*DataType
}
