package raml
import (
	
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
)

// NewNilType returns a new internal implementation of
// the raml.NilType interface.
//
// Generated @ 2020-07-02T14:31:30.98374873-04:00
func NewNilType() *NilType {
	out := &NilType{}

	out.DataType = NewDataType(rmeta.TypeNil, out)

	return out
}

// NilType is a generated internal implementation of
// the raml.NilType interface.
//
// Generated @ 2020-07-02T14:31:30.98374873-04:00
type NilType struct {
	*DataType
}


