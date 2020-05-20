package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
)

// NewNilType returns a new internal implementation of
// the raml.NilType interface.
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
func NewNilType(log *logrus.Entry) *NilType {
	out := &NilType{}

	out.DataType = NewDataType(
		rmeta.TypeNil,
		xlog.WithType(log, "internal.NilType"),
		out)

	return out
}

// NilType is a generated internal implementation of
// the raml.NilType interface.
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
type NilType struct {
	*DataType
}

func (a NilType) render() bool {
	return true
}

func (a NilType) ExtraFacets() raml.AnyMap {
	return NewAnyMap(a.log)
}

