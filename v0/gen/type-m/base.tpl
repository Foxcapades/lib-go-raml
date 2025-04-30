{{- /* gotype: github.com/Foxcapades/lib-go-raml/v0/tools/gen/type.extTypeProps */ -}}
package raml

{{define "base" -}}
import (
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
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
{{end}}