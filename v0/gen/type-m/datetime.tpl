{{- /* gotype: github.com/Foxcapades/lib-go-raml/v0/tools/gen/type.extTypeProps */ -}}
package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
)

func Foo() {
	out := &{{ .Name }}Type{}

{{define "datetime-constructor"}}{{end}}
}

type Bar struct {

	{{define "datetime-props" -}}
	format     raml.DateFormat
	{{- end}}

}

{{define "datetime-methods" -}}
func (o *{{.Name}}Type) Format() raml.DateFormat {
	return o.format
}

func (o *{{.Name}}Type) SetFormat(f raml.DateFormat) raml.{{.Name}}Type {
	o.format = f
	return o
}

func (o *{{.Name}}Type) UnsetFormat() raml.{{.Name}}Type {
	o.format = nil
	return o
}

{{end}}

func (o {{.Name}}Type) marshal(out raml.AnyMap) error {
	{{define "datetime-marshal" -}}
	out.PutIfNotNil(rmeta.KeyFormat, o.format)
	{{- end}}
}

func (o {{.Name}}Type) assign(key, val interface{}) (err error) {
{{define "datetime-assign"}}
	if key.Value == rmeta.KeyFormat {
		if val, err := DateFormatSortingHat(val); err != nil {
			return err
		} else {
			o.format = val
		}

		return nil
	}
{{end}}
}
