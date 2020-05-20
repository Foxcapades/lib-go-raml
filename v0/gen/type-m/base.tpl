{{- /* gotype: github.com/Foxcapades/lib-go-raml-types/v0/tools/gen/type.extTypeProps */ -}}
{{define "base" -}}{{if false}}package raml{{end}}
import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
)

// New{{.Name}}Type returns a new internal implementation of
// the raml.{{.Name}}Type interface.
//
// Generated @ {{.Time}}
func New{{.Name}}Type(log *logrus.Entry) *{{.Name}}Type {
	out := &{{.Name}}Type{}

	out.DataType = NewDataType(
		rmeta.Type{{.Name}},
		xlog.WithType(log, "internal.{{.Name}}Type"),
		out)

	return out
}

// {{.Name}}Type is a generated internal implementation of
// the raml.{{.Name}}Type interface.
//
// Generated @ {{.Time}}
type {{.Name}}Type struct {
	*DataType
}

func (a {{.Name}}Type) render() bool {
	return true
}

func (a {{.Name}}Type) ExtraFacets() raml.AnyMap {
	return NewAnyMap(a.log)
}
{{end}}