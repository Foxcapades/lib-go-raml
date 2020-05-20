{{- /*gotype: github.com/Foxcapades/lib-go-raml-types/v0/tools/gen/store.row*/ -}}
{{ define "impl" -}}
package raml

import (
	"fmt"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util"
	"strings"


{{if eq .Name "String" -}}
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	{{- end}}
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func New{{.Name}}Map() *{{.Name}}Map {
	return &{{.Name}}Map{
		{{ if .Raml -}}
			index: make(map[{{.KType}}]*raml.{{.Type}}),
		{{- else -}}
			index: make(map[{{.KType}}]*{{.Type}}),
		{{- end }}
	}
}

// {{.Name}}Map generated @ {{.Time}}
type {{.Name}}Map struct {
	slice []mapPair
	{{ if .Raml -}}
		index map[{{.KType}}]*raml.{{.Type}}
	{{- else -}}
		index map[{{.KType}}]*{{.Type}}
	{{- end }}
}

func (o *{{.Name}}Map) Len() uint {
	logrus.Trace("internal.{{.Name}}Map.Len")
	return uint(len(o.slice))
}

func (o *{{.Name}}Map) Put(key {{.KType}}, value {{if .Raml}}raml.{{end}}{{.Type}}) raml.{{.Name}}Map {
	logrus.Trace("internal.{{.Name}}Map.Put")
	o.index[key] = &value
	o.slice = append(o.slice, mapPair{key: key, val: value})
	return o
}

{{ if .Raml -}}
	func (o *{{.Name}}Map) PutNonNil(key {{.KType}}, value raml.{{.Type}}) raml.{{.Name}}Map
{{- else if ne .Kind "Untyped" -}}
	func (o *{{.Name}}Map) PutNonNil(key {{.KType}}, value *{{.Type}}) raml.{{.Name}}Map
{{- else -}}
	func (o *{{.Name}}Map) PutNonNil(key {{.KType}}, value {{.Type}}) raml.{{.Name}}Map
{{- end }} {
	logrus.Trace("internal.{{.Name}}Map.PutNonNil")

	if !util.IsNil(value) {
		return o.Put(key, {{if (and (not .Raml) (ne .Kind "Untyped"))}}*{{end}}value)
	}

	return o
}

{{ if .Raml -}}
	func (o *{{.Name}}Map) Replace(key {{.KType}}, value raml.{{.Type}}) raml.{{.Kind}}
{{- else -}}
	func (o *{{.Name}}Map) Replace(key {{.KType}}, value {{.Type}}) option.{{.Kind}}
{{- end }} {
	logrus.Trace("internal.{{.Name}}Map.Replace")

	ind := o.IndexOf(key)

	if ind.IsNil() {
		{{ if .Raml -}}
			return nil
		{{- else -}}
			return option.NewEmpty{{.Kind}}()
		{{- end }}
	}

	{{ if .Raml -}}
		out := *o.index[key]
	{{- else -}}
		out := option.NewMaybe{{.Kind}}(o.index[key])
	{{- end }}

	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

{{ if .Raml -}}
	func (o *{{.Name}}Map) ReplaceOrPut(key {{.KType}}, value raml.{{.Type}}) raml.{{.Kind}}
{{- else -}}
	func (o *{{.Name}}Map) ReplaceOrPut(key {{.KType}}, value {{.Type}}) option.{{.Kind}}
{{- end }} {
	logrus.Trace("internal.{{.Name}}Map.ReplaceOrPut")

	ind := o.IndexOf(key)

	if ind.IsNil() {
		o.index[key] = &value
		o.slice = append(o.slice, mapPair{key: key, val: value})
		return {{if .Raml}}nil{{else}}option.NewEmpty{{.Kind}}(){{end}}
	}

	out := {{if .Raml}}*o.index[key]{{else}}option.NewMaybe{{.Kind}}(o.index[key]){{end}}
	o.index[key] = &value
	o.slice[ind.Get()].val = value
	return out
}

func (o *{{.Name}}Map) Get(key {{.KType}}) {{if .Raml}}raml{{else}}option{{end}}.{{.Kind}} {
	logrus.Trace("internal.{{.Name}}Map.Get")

	if !o.Has(key) {
		return {{if .Raml}}nil{{else}}option.NewEmpty{{.Kind}}(){{end}}
	}

	return {{if .Raml}}*o.index[key]{{else}}option.NewMaybe{{.Kind}}(o.index[key]){{end}}
}

{{if .Raml}}
func (o *{{.Name}}Map) At(index uint) (key option.{{.KName}}, value raml.{{.Kind}}) {
{{else}}
func (o *{{.Name}}Map) At(index uint) (key option.{{.KName}}, value option.{{.Kind}}) {
{{end}}
	logrus.Trace("internal.{{.Name}}Map.At")

	tmp := &o.slice[index]
	key = option.New{{.KName}}(tmp.key.(string))

	{{ if .Raml -}}
	value = tmp.val.(raml.{{.Name}})
	{{- else }}
	if util.IsNil(tmp.val) {
		value = option.NewEmpty{{.Kind}}()
	} else {
		value = option.New{{.Kind}}(tmp.val.({{.Type}}))
	}
	{{- end }}

	return
}

func (o *{{.Name}}Map) IndexOf(key {{.KType}}) option.Uint {
	logrus.Trace("internal.{{.Name}}Map.IndexOf")
	if !o.Has(key) {
		return option.NewEmptyUint()
	}
	for i := range o.slice {
		if o.slice[i].key == key {
			return option.NewUint(uint(i))
		}
	}
	panic("invalid map state, index out of sync")
}

func (o *{{.Name}}Map) Has(key {{.KType}}) bool {
	logrus.Trace("internal.{{.Name}}Map.Has")

	_, ok := o.index[key]
	return ok
}

{{ if .Raml -}}
	func (o *{{.Name}}Map) Delete(key {{.KType}}) raml.{{.Kind}} {
{{- else -}}
	func (o *{{.Name}}Map) Delete(key {{.KType}}) option.{{.Kind}} {
{{- end }}
	logrus.Trace("internal.{{.Name}}Map.Delete")

	if !o.Has(key) {
		return {{if .Raml}}nil{{else}}option.NewEmpty{{.Kind}}(){{end}}
	}

	{{ if .Raml -}}
		out := *o.index[key]
	{{- else -}}
		out := option.NewMaybe{{.Kind}}(o.index[key])
	{{- end }}
	delete(o.index, key)

	for i := range o.slice {
		if o.slice[i].key == key {
			o.slice = append(o.slice[:i], o.slice[i+1:]...)
			return out
		}
	}
	panic("invalid map state, index out of sync")
}

func (o {{.Name}}Map) ForEach(fn func({{.KType}}, {{if .Raml}}raml.{{end}}{{.Type}})) {
	logrus.Trace("internal.{{.Name}}Map.ForEach")

	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o {{.Name}}Map) MarshalYAML() (interface{}, error) {
	logrus.Trace("internal.{{.Name}}Map.MarshalYAML")

	out := xyml.MapNode(len(o.slice)*2)
	for i := range o.slice {
		if err := xyml.AppendToMap(out, o.slice[i].key, o.slice[i].val); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (o *{{.Name}}Map) UnmarshalRAML(val *yaml.Node) (err error) {
	logrus.Trace("internal.{{.Name}}Map.UnmarshalRAML")

	if err := xyml.RequireMapping(val); err != nil {
		return err
	}

	for i := 0; i < len(val.Content); i += 2 {
		key := val.Content[i]
		val := val.Content[i+1]

		{{if eq .KType "interface{}" -}}
			altKey, err := xyml.CastScalarToYmlType(key)
			if err != nil {
				return err
			}
		{{- else -}}
			altKey := key.Value
		{{- end}}

		{{if .Raml}}
			{{if eq .Kind "Example"}}
				tmpVal := ExampleSortingHat(val)
			{{else if eq .Kind "DataType"}}
				tmpVal, err := TypeSortingHat(val)
				if err != nil {
					return err
				}
			{{else if eq .Kind "Property"}}
				tmpVal, err := PropertySortingHat(val)
				if err != nil {
					return err
				}
			{{else}}
				tmpVal := New{{.Name}}()
				if err = tmpVal.UnmarshalRAML(val); err != nil {
					return err
				}
			{{end}}
		{{else if eq .Kind "Untyped"}}
			tmpVal := val
		{{else}}
			var tmpVal {{ .Type }}
			if err = assign.As{{ .Name }}(val, &tmpVal); err != nil {
				return err
			}
		{{end}}

		o.Put(altKey, tmpVal)
	}

	return nil
}

func (o *{{.Name}}Map) String() string {
	tmp := strings.Builder{}
	enc := yaml.NewEncoder(&tmp)
	enc.SetIndent(2)
	if err := enc.Encode(o.index); err != nil {
		return fmt.Sprint(o.index)
	} else {
		return tmp.String()
	}
}
{{end}}