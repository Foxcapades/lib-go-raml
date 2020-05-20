package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
)

// NewAnyType returns a new internal implementation of
// the raml.AnyType interface.
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
func NewAnyType(log *logrus.Entry) *AnyType {
	out := &AnyType{}

	out.DataType = NewDataType(
		rmeta.TypeAny,
		xlog.WithType(log, "internal.AnyType"),
		out)

	return out
}

// AnyType is a generated internal implementation of
// the raml.AnyType interface.
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
type AnyType struct {
	*DataType
}

func (a AnyType) render() bool {
	return true
}

func (a AnyType) ExtraFacets() raml.AnyMap {
	return NewAnyMap(a.log)
}

