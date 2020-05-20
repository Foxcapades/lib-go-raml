{{- /* gotype: github.com/Foxcapades/lib-go-raml-types/v0/tools/gen/type.extTypeProps */ -}}
package raml

{{define "base" -}}
import (
	{{if eq .Name "Include" -}}
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	{{- end}}
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	{{if not (eq .Name "Any" "Nil") -}}
	"github.com/sirupsen/logrus"
	{{- end}}
)

// New{{.Name}}Type returns a new internal implementation of
// the raml.{{.Name}}Type interface.
//
// Generated @ {{.Time}}
func New{{.Name}}Type() *{{.Name}}Type {
	out := &{{.Name}}Type{}

	out.DataType = NewDataType(rmeta.Type{{.Name}}, out)

	return out
}

// {{.Name}}Type is a generated internal implementation of
// the raml.{{.Name}}Type interface.
//
// Generated @ {{.Time}}
type {{.Name}}Type struct {
	*DataType
}
{{if eq .Name "Include"}}
func (a {{.Name}}Type) marshal(out raml.AnyMap) error {
	logrus.Trace("internal.{{.Name}}Type.marshal")

	out.Put(rmeta.KeyType, "!include " + a.DataType.schema)
	a.DataType.hasExtra.out(out)
	return nil
}
{{end}}
{{end}}