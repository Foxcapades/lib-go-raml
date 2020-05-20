{{- /* gotype: github.com/Foxcapades/lib-go-raml-types/v0/tools/gen/type.extTypeProps */ -}}
package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
)

func Foo() {
	out := &{{ .Name }}Type{}

{{define "integer-constructor"}}{{end}}
}

type Bar struct {

	{{define "integer-props" -}}
	minimum    *int64
	maximum    *int64
	format     raml.NumberFormat
	multipleOf *float64
	{{- end}}

}

func (o {{.Name}}Type) marshal(out raml.AnyMap) error {
	{{define "integer-marshal" -}}
	out.PutNonNil(rmeta.KeyFormat, o.format).
		PutNonNil(rmeta.KeyMinimum, o.minimum).
		PutNonNil(rmeta.KeyMaximum, o.maximum).
		PutNonNil(rmeta.KeyMultipleOf, o.multipleOf)
	{{- end}}
}

{{define "integer-methods" -}}
func (o *{{.Name}}Type) Minimum() option.Int64 {
	return option.NewMaybeInt64(o.minimum)
}

func (o *{{.Name}}Type) SetMinimum(pat int64) raml.{{.Name}}Type {
	o.minimum = &pat
	return o
}

func (o *{{.Name}}Type) UnsetMinimum() raml.{{.Name}}Type {
	o.minimum = nil
	return o
}

func (o *{{.Name}}Type) Maximum() option.Int64 {
	return option.NewMaybeInt64(o.maximum)
}

func (o *{{.Name}}Type) SetMaximum(pat int64) raml.{{.Name}}Type {
	o.maximum = &pat
	return o
}

func (o *{{.Name}}Type) UnsetMaximum() raml.{{.Name}}Type {
	o.maximum = nil
	return o
}

func (o *{{.Name}}Type) Format() raml.IntegerFormat {
	return o.format
}

func (o *{{.Name}}Type) SetFormat(f raml.IntegerFormat) raml.{{.Name}}Type {
	o.format = f
	return o
}

func (o *{{.Name}}Type) UnsetFormat() raml.{{.Name}}Type {
	o.format = nil
	return o
}

func (o *{{.Name}}Type) MultipleOf() option.Float64 {
	return option.NewMaybeFloat64(o.multipleOf)
}

func (o *{{.Name}}Type) SetMultipleOf(pat float64) raml.{{.Name}}Type {
	o.multipleOf = &pat
	return o
}

func (o *{{.Name}}Type) UnsetMultipleOf() raml.{{.Name}}Type {
	o.multipleOf = nil
	return o
}

func (o {{.Name}}Type) render() bool {
	return true
}
{{end}}
func (o {{.Name}}Type) assign(key, val interface{}) (err error) {
{{define "integer-assign"}}
	switch key {
	case rmeta.KeyMinimum:
		return assign.AsInt64Ptr(val, &o.minimum, log)
	case rmeta.KeyMaximum:
		return assign.AsInt64Ptr(val, &o.maximum, log)
	case rmeta.KeyFormat:
		if val, err := IntegerFormatSortingHat(val, log); err != nil {
			return err
		} else {
			o.format = val
			return nil
		}
	case rmeta.KeyMultipleOf:
		return assign.AsFloat64Ptr(val, &o.multipleOf, log)
	}
{{end}}
	return
}
