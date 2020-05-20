{{- /* gotype: github.com/Foxcapades/lib-go-raml-types/v0/tools/gen/type.extTypeProps */ -}}
package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
)

func Foo() {
	out := &{{ .Name }}Type{}

{{define "object-constructor"}}
	out.properties = NewPropertyMap(log)
	out.addtlProps = true
{{end}}
}

type Bar struct {

	{{define "object-props" -}}
	properties raml.PropertyMap
	minProps   *uint
	maxProps   *uint
	addtlProps bool
	discrim    *string
	discrimVal *interface{}
	{{- end}}

}

func (o {{.Name}}Type) marshal(out raml.AnyMap) error {
	{{define "object-marshal" -}}
	if !o.addtlProps {
		out.Put(rmeta.KeyAddtlProps, o.addtlProps)
	}

	out.PutNonNil(rmeta.KeyMinProperties, o.minProps).
		PutNonNil(rmeta.KeyMaxProperties, o.maxProps).
		PutNonNil(rmeta.KeyDiscriminator, o.discrim).
		PutNonNil(rmeta.KeyDiscriminatorVal, o.discrimVal)

	if o.properties.Len() > 0 {
		out.Put(rmeta.KeyProperties, o.properties)
	}
	{{- end}}
}

{{define "object-methods" -}}
func (o *{{.Name}}Type) Properties() raml.PropertyMap {
	return o.properties
}

func (o *{{.Name}}Type) SetProperties(props raml.PropertyMap) raml.{{.Name}}Type {
	if props == nil {
		return o.UnsetProperties()
	}

	o.properties = props
	return o
}

func (o *{{.Name}}Type) UnsetProperties() raml.{{.Name}}Type {
	o.properties = NewPropertyMap(o.DataType.log)
	return o
}

func (o *{{.Name}}Type) MinProperties() option.Uint {
	return option.NewMaybeUint(o.minProps)
}

func (o *{{.Name}}Type) SetMinProperties(min uint) raml.{{.Name}}Type {
	o.minProps = &min
	return o
}

func (o *{{.Name}}Type) UnsetMinProperties() raml.{{.Name}}Type {
	o.minProps = nil
	return o
}

func (o *{{.Name}}Type) MaxProperties() option.Uint {
	return option.NewMaybeUint(o.maxProps)
}

func (o *{{.Name}}Type) SetMaxProperties(u uint) raml.{{.Name}}Type {
	o.maxProps = &u
	return o
}

func (o *{{.Name}}Type) UnsetMaxProperties() raml.{{.Name}}Type {
	o.maxProps = nil
	return o
}

func (o *{{.Name}}Type) AdditionalProperties() bool {
	return o.addtlProps
}

func (o *{{.Name}}Type) SetAdditionalProperties(val bool) raml.{{.Name}}Type {
	o.addtlProps = val
	return o
}

func (o *{{.Name}}Type) Discriminator() option.String {
	return option.NewMaybeString(o.discrim)
}

func (o *{{.Name}}Type) SetDiscriminator(facet string) raml.{{.Name}}Type {
	o.discrim = &facet
	return o
}

func (o *{{.Name}}Type) UnsetDiscriminator() raml.{{.Name}}Type {
	o.discrim = nil
	return o
}

func (o *{{.Name}}Type) DiscriminatorValue() option.Untyped {
	return option.NewMaybeUntyped(o.discrimVal)
}

func (o *{{.Name}}Type) SetDiscriminatorValue(val interface{}) raml.{{.Name}}Type {
	o.discrimVal = &val
	return o
}

func (o *{{.Name}}Type) UnsetDiscriminatorValue() raml.{{.Name}}Type {
	o.discrimVal = nil
	return o
}

func (o {{.Name}}Type) render() bool {
	return true
}
{{end}}
func (o {{.Name}}Type) assign(key, val interface{}, log *logrus.Entry) (err error) {
{{define "object-assign"}}
	switch key {
	case rmeta.KeyProperties:
		xlog.OptError(log, o.properties.UnmarshalRAML(val, log))
	case rmeta.KeyMinProperties:
		return assign.AsUintPtr(val, &o.minProps)
	case rmeta.KeyMaxProperties:
		return assign.AsUintPtr(val, &o.maxProps)
	case rmeta.KeyAddtlProps:
		return assign.AsBool(val, &o.addtlProps, log)
	case rmeta.KeyDiscriminator:
		return assign.AsStringPtr(val, &o.discrim, log)
	case rmeta.KeyDiscriminatorVal:
		o.discrimVal = &val
		return nil
	}
{{end}}
}
