package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
)

// NewIncludeType returns a new internal implementation of
// the raml.IncludeType interface.
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
func NewIncludeType(log *logrus.Entry) *IncludeType {
	out := &IncludeType{}

	out.DataType = NewDataType(
		rmeta.TypeInclude,
		xlog.WithType(log, "internal.IncludeType"),
		out)

	return out
}

// IncludeType is a generated internal implementation of
// the raml.IncludeType interface.
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
type IncludeType struct {
	*DataType
}

func (a IncludeType) render() bool {
	return true
}

func (a IncludeType) ExtraFacets() raml.AnyMap {
	return NewAnyMap(a.log)
}

