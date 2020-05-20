{{- /*gotype: github.com/Foxcapades/lib-go-raml-types/v0/tools/gen/store.row*/ -}}
{{ define "impl" -}}
package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func New{{.Name}}Map(log *logrus.Entry) *{{.Name}}Map {
	return &{{.Name}}Map{
		log:   xlog.WithType(log, "internal.{{.Name}}Map"),
		{{ if .Raml -}}
			index: make(map[{{.KType}}]*raml.{{.Type}}),
		{{- else -}}
			index: make(map[{{.KType}}]*{{.Type}}),
		{{- end }}
	}
}

// {{.Name}}Map generated @ {{.Time}}
type {{.Name}}Map struct {
	log   *logrus.Entry
	slice yaml.MapSlice
	{{ if .Raml -}}
		index map[{{.KType}}]*raml.{{.Type}}
	{{- else -}}
		index map[{{.KType}}]*{{.Type}}
	{{- end }}
}

func (o *{{.Name}}Map) Len() uint {
	return uint(len(o.slice))
}

func (o *{{.Name}}Map) Put(key {{.KType}}, value {{if .Raml}}raml.{{end}}{{.Type}}) raml.{{.Name}}Map {
	o.index[key] = &value
	o.slice = append(o.slice, yaml.MapItem{Key: key, Value: value})
	return o
}

{{ if .Raml -}}
	func (o *{{.Name}}Map) PutNonNil(key {{.KType}}, value raml.{{.Type}}) raml.{{.Name}}Map
{{- else if ne .Kind "Untyped" -}}
	func (o *{{.Name}}Map) PutNonNil(key {{.KType}}, value *{{.Type}}) raml.{{.Name}}Map
{{- else -}}
	func (o *{{.Name}}Map) PutNonNil(key {{.KType}}, value {{.Type}}) raml.{{.Name}}Map
{{- end }} {
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
	o.slice[ind.Get()].Value = value
	return out
}

{{ if .Raml -}}
	func (o *{{.Name}}Map) ReplaceOrPut(key {{.KType}}, value raml.{{.Type}}) raml.{{.Kind}}
{{- else -}}
	func (o *{{.Name}}Map) ReplaceOrPut(key {{.KType}}, value {{.Type}}) option.{{.Kind}}
{{- end }} {
	ind := o.IndexOf(key)

	if ind.IsNil() {
		o.index[key] = &value
		o.slice = append(o.slice, yaml.MapItem{Key: key, Value: value})
		return {{if .Raml}}nil{{else}}option.NewEmpty{{.Kind}}(){{end}}
	}

	out := {{if .Raml}}*o.index[key]{{else}}option.NewMaybe{{.Kind}}(o.index[key]){{end}}
	o.index[key] = &value
	o.slice[ind.Get()].Value = value
	return out
}

func (o *{{.Name}}Map) Get(key {{.KType}}) {{if .Raml}}raml{{else}}option{{end}}.{{.Kind}} {
	if !o.Has(key) {
		return {{if .Raml}}nil{{else}}option.NewEmpty{{.Kind}}(){{end}}
	}

	return {{if .Raml}}*o.index[key]{{else}}option.NewMaybe{{.Kind}}(o.index[key]){{end}}
}

func (o *{{.Name}}Map) At(index uint) (key option.{{.KName}}, value {{if .Raml}}raml{{else}}option{{end}}.{{.Kind}}) {
	tmp := &o.slice[index]
	key = option.New{{.KName}}(tmp.Key.(string))
	{{ if .Raml -}}
	value = tmp.Value.(raml.{{.Name}})
	{{- else }}
	if util.IsNil(tmp.Value) {
		value = option.NewEmpty{{.Kind}}()
	} else {
		value = option.New{{.Kind}}(tmp.Value.({{.Type}}))
	}
	{{- end }}

	return
}

func (o *{{.Name}}Map) IndexOf(key {{.KType}}) option.Uint {
	if !o.Has(key) {
		return option.NewEmptyUint()
	}
	for i := range o.slice {
		if o.slice[i].Key == key {
			return option.NewUint(uint(i))
		}
	}
	panic("invalid map state, index out of sync")
}

func (o *{{.Name}}Map) Has(key {{.KType}}) bool {
	_, ok := o.index[key]
	return ok
}

{{ if .Raml -}}
	func (o *{{.Name}}Map) Delete(key {{.KType}}) raml.{{.Kind}} {
{{- else -}}
	func (o *{{.Name}}Map) Delete(key {{.KType}}) option.{{.Kind}} {
{{- end }}
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
		if o.slice[i].Key == key {
			o.slice = append(o.slice[:i], o.slice[i+1:]...)
			return out
		}
	}
	panic("invalid map state, index out of sync")
}

func (o {{.Name}}Map) ForEach(fn func({{.KType}}, {{if .Raml}}raml.{{end}}{{.Type}})) {
	for k, v := range o.index {
		fn(k, *v)
	}
}

func (o {{.Name}}Map) MarshalYAML() (interface{}, error) {
	return o.slice, nil
}

func (o *{{.Name}}Map) UnmarshalRAML(val interface{}, log *logrus.Entry) (err error) {
	log.Trace("internal.{{.Name}}Map.UnmarshalRAML")
	yml, err := assign.AsMapSlice(val)

	if err != nil {
		return xlog.Error(log, err)
	}

	for i := range yml {
		tmp := &yml[i]
		{{if ne .Name "Any" -}}
			l2 := xlog.AddPath(log, tmp.Key)
		{{- end}}

		{{if ne .KType "interface{}" -}}
		key := ""

		if err = assign.AsString(tmp.Key, &key, l2); err != nil {
			return xlog.Error(l2, err)
		}
		{{- else -}}
		key := tmp.Key
		{{- end}}

		{{ if .Raml -}}

		{{- if eq .Kind "Example" -}}
		tmpVal := ExampleSortingHat(tmp.Value, l2)
		{{- else if eq .Kind "DataType" -}}
		tmpVal, err := TypeSortingHat(tmp.Value, l2)
		if err != nil {
			return xlog.Error(l2, err)
		}
		{{- else if eq .Kind "Property" -}}
		tmpVal, err := PropertySortingHat(tmp.Value, l2)
		if err != nil {
			return xlog.Error(l2, err)
		}
		{{- else -}}
		tmpVal := New{{.Name}}(l2)
		if err = tmpVal.UnmarshalRAML(tmp.Value, l2); err != nil {
			return xlog.Error(l2, err)
		}
		{{- end -}}

		{{- else if eq .Kind "Untyped" -}}
		tmpVal := tmp.Value
		{{- else -}}
		var tmpVal {{ .Type }}
		if err = assign.As{{ .Name }}(tmp.Value, &tmpVal, l2); err != nil {
			return xlog.Error(l2, err)
		}
		{{- end }}

		o.Put(key, tmpVal)
	}

	return nil
}
{{end}}