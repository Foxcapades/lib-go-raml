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

{{define "file-constructor"}}
	out.minLength = rmeta.FileDefaultMinLength
	out.maxLength = rmeta.FileDefaultMaxLength
{{end}}
}

type Bar struct {

	{{define "file-props" -}}
	fileTypes []string
	minLength uint
	maxLength uint
	{{- end}}

}

func (o {{.Name}}Type) marshal(out raml.AnyMap) error {
	{{define "file-marshal" -}}
	out.PutNonNil(rmeta.KeyFileTypes, o.fileTypes)

	if o.minLength != rmeta.FileDefaultMinLength {
		out.Put(rmeta.KeyMinLength, o.minLength)
	}

	if o.maxLength != rmeta.FileDefaultMaxLength {
		out.Put(rmeta.KeyMaxLength, o.maxLength)
	}
	{{- end}}
}

{{define "file-methods" -}}
func (o *{{.Name}}Type) FileTypes() []string {
	return o.fileTypes
}

func (o *{{.Name}}Type) SetFileTypes(val []string) raml.{{.Name}}Type {
	o.fileTypes = val
	return o
}

func (o *{{.Name}}Type) UnsetFileTypes() raml.{{.Name}}Type {
	o.fileTypes = nil
	return o
}

func (o *{{.Name}}Type) MinLength() uint {
	return o.minLength
}

func (o *{{.Name}}Type) SetMinLength(min uint) raml.{{.Name}}Type {
	o.minLength = min
	return o
}

func (o *{{.Name}}Type) MaxLength() uint {
	return o.maxLength
}

func (o *{{.Name}}Type) SetMaxLength(u uint) raml.{{.Name}}Type {
	o.maxLength = u
	return o
}

{{end}}
func (o {{.Name}}Type) assign(key, val interface{}) (err error) {
{{define "file-assign"}}
	switch key.Value {
	case rmeta.KeyFileTypes:
		return assign.AsStringList(val, &o.fileTypes)
	case rmeta.KeyMinLength:
		return assign.ToUint(val, &o.minLength)
	case rmeta.KeyMaxLength:
		return assign.ToUint(val, &o.maxLength)
	}
{{end}}
	return
}
