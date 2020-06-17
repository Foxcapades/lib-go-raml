{{- /*gotype: github.com/Foxcapades/lib-go-raml/v0/tools/gen./type.extTypeProps*/ -}}
package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
	"gopkg.in/yaml.v3"
)

func Foo() {
	out := &{{ .Name }}Type{}

{{define "array-constructor"}}
	out.minItems    = rmeta.ArrayDefaultMinItems
	out.maxItems    = rmeta.ArrayDefaultMaxItems
	out.uniqueItems = rmeta.ArrayDefaultUniqueItems
{{end}}
}

type Bar struct {

	{{define "array-props" -}}
	uniqueItems bool
	minItems    uint
	maxItems    uint
	items       raml.DataType
	{{- end}}

}

{{define "array-methods" -}}
func (o *{{.Name}}Type) UniqueItems() bool {
	return o.uniqueItems
}

func (o *{{.Name}}Type) SetUniqueItems(val bool) raml.{{.Name}}Type {
	o.uniqueItems = val
	return o
}

func (o *{{.Name}}Type) MinItems() uint {
	return o.minItems
}

func (o *{{.Name}}Type) SetMinItems(min uint) raml.{{.Name}}Type {
	o.minItems = min
	return o
}

func (o *{{.Name}}Type) MaxItems() uint {
	return o.maxItems
}

func (o *{{.Name}}Type) SetMaxItems(u uint) raml.{{.Name}}Type {
	o.maxItems = u
	return o
}

func (o *{{.Name}}Type) Items() raml.DataType {
	return o.items
}

func (o *{{.Name}}Type) SetItems(val raml.DataType) raml.{{.Name}}Type {
	o.items = val
	return o
}

func (o *{{.Name}}Type) UnsetItems() raml.{{.Name}}Type {
	o.items = nil
	return o
}

{{end}}

func (o {{.Name}}Type) marshal(out raml.AnyMap) error {
{{define "array-marshal"}}
	if o.uniqueItems != rmeta.ArrayDefaultUniqueItems {
		out.Put(rmeta.KeyUniqueItems, o.uniqueItems)
	}

	if o.minItems != rmeta.ArrayDefaultMinItems {
		out.Put(rmeta.KeyMinItems, o.minItems)
	}

	if o.maxItems != rmeta.ArrayDefaultMaxItems {
		out.Put(rmeta.KeyMaxItems, o.maxItems)
	}

	out.PutIfNotNil(rmeta.KeyItems, o.items)
{{end}}
}

func (o {{.Name}}Type) assign(key, val *yaml.Node) (err error) {
{{define "array-assign"}}
	switch key.Value {
	case rmeta.KeyUniqueItems:
		return assign.AsBool(val, &o.uniqueItems)
	case rmeta.KeyMinItems:
		return assign.ToUint(val, &o.minItems)
	case rmeta.KeyMaxItems:
		return assign.ToUint(val, &o.maxItems)
	case rmeta.KeyItems:
		if val, err := TypeSortingHat(val); err == nil {
			o.items = val
			return nil
		} else {
			return err
		}
	}
{{end}}
}
