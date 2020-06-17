{{- /* gotype: github.com/Foxcapades/lib-go-raml/v0/tools/gen/type.extTypeProps */ -}}
package raml

{{define "base" -}}
import (
	{{if eq .Name "Include" -}}
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"gopkg.in/yaml.v3"

{{- end}}
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
{{if eq .Name "Include"}}
func (a {{.Name}}Type) marshal(out raml.AnyMap) error {
	out.Put(rmeta.KeyType, &yaml.Node{
		Kind: yaml.ScalarNode,
		Tag: "!include",
		Value: a.DataType.schema,
	})
	a.DataType.hasExtra.out(out)

	return nil
}
{{end}}
{{end}}