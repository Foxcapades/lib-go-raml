{{- /* gotype: github.com/Foxcapades/lib-go-raml/v0/tools/gen/type.extTypeProps */ -}}
package raml

import (
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
	"gopkg.in/yaml.v3"
)

func Foo() {
	out := &{{ .Name }}Type{}
{{define "include-constructor"}}{{end}}
}

type Bar struct {
	{{define "include-props"}}{{end}}
}

func (o {{.Name}}Type) marshal(out raml.AnyMap) error {
	{{define "include-marshal" -}}
	out.Put(rmeta.KeyType, &yaml.Node{
		Kind: yaml.ScalarNode,
		Tag: "!include",
		Value: o.DataType.schema,
	})
	{{- end}}
}

{{define "include-methods" -}}{{end}}

func (o {{.Name}}Type) assign(key, val *yaml.Value) (err error) {
{{define "include-assign"}}
	switch key.Value {
	case rmeta.KeyType, rmeta.KeySchema:
		o.DataType.schema = val.Value
		return nil
	}
{{end}}
	return
}
