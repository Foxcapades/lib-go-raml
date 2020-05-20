{{- /* gotype: github.com/Foxcapades/lib-go-raml-types/v0/tools/gen/type.extTypeProps */ -}}
package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
)

func Foo() {
	out := &{{ .Name }}Type{}

{{define "string-constructor"}}
	out.minLength = rmeta.StringDefaultMinLength
{{end}}
}

type Bar struct {

	{{define "string-props" -}}
	pattern      *string
	minLength    uint
	maxLength    *uint
	{{- end}}

}

func (o {{.Name}}Type) marshal(out raml.AnyMap) error {
	{{define "string-marshal" -}}
	out.PutNonNil(rmeta.KeyPattern, o.pattern)
	if o.minLength > 0 {
		out.Put(rmeta.KeyMinLength, o.minLength)
	}
	out.PutNonNil(rmeta.KeyMaxLength, o.maxLength)
	{{- end}}
}

{{define "string-methods" -}}
func (o *{{.Name}}Type) Pattern() option.String {
	return option.NewMaybeString(o.pattern)
}

func (o *{{.Name}}Type) SetPattern(pat string) raml.{{.Name}}Type {
	o.pattern = &pat
	return o
}

func (o *{{.Name}}Type) UnsetPattern() raml.{{.Name}}Type {
	o.pattern = nil
	return o
}

func (o *{{.Name}}Type) MinLength() uint {
	return o.minLength
}

func (o *{{.Name}}Type) SetMinLength(min uint) raml.{{.Name}}Type {
	o.minLength = min
	return o
}

func (o *{{.Name}}Type) MaxLength() option.Uint {
	return option.NewMaybeUint(o.maxLength)
}

func (o *{{.Name}}Type) SetMaxLength(u uint) raml.{{.Name}}Type {
	o.maxLength = &u
	return o
}

func (o *{{.Name}}Type) UnsetMaxLength() raml.{{.Name}}Type {
	o.maxLength = nil
	return o
}

func (o {{.Name}}Type) render() bool {
	return true
}
{{end}}
func (o {{.Name}}Type) assign(key, val *yaml.Value) (err error) {
{{define "string-assign"}}
	switch key.Value {
	case rmeta.KeyPattern:
		return assign.AsStringPtr(val, &o.pattern)
	case rmeta.KeyMinLength:
		return assign.ToUint(val, &o.minLength)
	case rmeta.KeyMaxLength:
		return assign.AsUintPtr(val, &o.maxLength)
	}
{{end}}
	return
}
