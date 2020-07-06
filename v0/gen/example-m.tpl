{{- /*gotype: github.com/Foxcapades/lib-go-raml/v0/tools/gen/example.exampleProps*/ -}}
{{define "example" -}}
package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// New{{.Name}}Example returns a new internal implementation of the
// raml.{{.Name}}Example interface.
//
// Generated @ {{.Time}}
func New{{.Name}}Example() *{{.Name}}Example {
	return &{{.Name}}Example{
		annotations: raml.NewAnnotationMap(0).SerializeOrdered(false),
		extra:       raml.NewAnyMap(0).SerializeOrdered(false),
		strict:      rmeta.ExampleDefaultStrict,
	}
}

// {{.Name}}Example is a generated internal implementation of the
// raml.{{.Name}}Example interface.
type {{.Name}}Example struct {
	displayName *string
	description *string
	annotations raml.AnnotationMap
	value       {{if .UseOption}}*{{end}}{{.Type}}
	strict      bool
	extra       raml.AnyMap
}

func (e *{{.Name}}Example) DisplayName() option.String {
	return option.NewMaybeString(e.displayName)
}

func (e *{{.Name}}Example) SetDisplayName(name string) raml.{{.Name}}Example {
	e.displayName = &name
	return e
}

func (e *{{.Name}}Example) UnsetDisplayName() raml.{{.Name}}Example {
	e.displayName = nil
	return e
}

func (e *{{.Name}}Example) Description() option.String {
	return option.NewMaybeString(e.description)
}

func (e *{{.Name}}Example) SetDescription(desc string) raml.{{.Name}}Example {
	e.description = &desc
	return e
}

func (e *{{.Name}}Example) UnsetDescription() raml.{{.Name}}Example {
	e.description = nil
	return e
}

func (e *{{.Name}}Example) Annotations() raml.AnnotationMap {
	return e.annotations
}

func (e *{{.Name}}Example) SetAnnotations(ann raml.AnnotationMap) raml.{{.Name}}Example {
	if ann == nil {
		return e.UnsetAnnotations()
	}
	e.annotations = ann
	return e
}

func (e *{{.Name}}Example) UnsetAnnotations() raml.{{.Name}}Example {
	e.annotations = raml.NewAnnotationMap(0)
	return e
}

func (e *{{.Name}}Example) Value() {{if .UseOption}}option.{{end}}{{.EType}} {
	{{if .UseOption -}}
	return option.NewMaybe{{.EType}}(e.value)
	{{- else -}}
	return e.value
	{{- end}}
}

func (e *{{.Name}}Example) SetValue(val {{.Type}}) raml.{{.Name}}Example {
	e.value = {{if .UseOption}}&{{end}}val
	return e
}

func (e *{{.Name}}Example) UnsetValue() raml.{{.Name}}Example {
	e.value = nil
	return e
}

func (e *{{.Name}}Example) Strict() bool {
	return e.strict
}

func (e *{{.Name}}Example) SetStrict(b bool) raml.{{.Name}}Example {
	e.strict = b
	return e
}

func (e *{{.Name}}Example) ExtraFacets() raml.AnyMap {
	return e.extra
}

func (e *{{.Name}}Example) UnmarshalRAML(value *yaml.Node) error {

	if xyml.IsMap(value) {
		return xyml.MapForEach(value, e.assign)
	}

	return e.assignVal(value)
}

func (e *{{.Name}}Example) MarshalRAML(out raml.AnyMap) (bool, error) {
	if e.expand() {
		out.PutIfNotNil(rmeta.KeyDisplayName, e.displayName).
			PutIfNotNil(rmeta.KeyDescription, e.description).
			PutIfNotNil(rmeta.KeyValue, e.value)

		if e.strict != rmeta.ExampleDefaultStrict {
			out.Put(rmeta.KeyStrict, e.strict)
		}

		e.annotations.ForEach(func(k string, v raml.Annotation) { out.Put(k, v) })
		e.extra.ForEach(func(k interface{}, v interface{}) { out.Put(k, v) })

		return false, nil
	}

	out.Put("", e.value)
	return true, nil
}

func (e *{{.Name}}Example) assign(key, val *yaml.Node) error {
	if !xyml.IsString(key) {
		if ver, err := xyml.ToScalarValue(key); err != nil {
			return err
		} else {
			e.extra.Put(ver, val)
		}
		return nil
	}

	if key.Value[0] == '(' {
		tmp := NewAnnotation()
		if err := UnmarshalUntypedMapRAML(tmp, val); err != nil {
			return err
		}
		e.annotations.Put(key.Value, tmp)
		return nil
	}

	switch key.Value {
	case rmeta.KeyDisplayName:
		return assign.AsStringPtr(val, &e.displayName)
	case rmeta.KeyDescription:
		return assign.AsStringPtr(val, &e.description)
	case rmeta.KeyStrict:
		return assign.AsBool(val, &e.strict)
	case rmeta.KeyValue:
		return e.assignVal(val)
	}

	if ver, err := xyml.ToScalarValue(key); err != nil {
		return err
	} else {
		e.extra.Put(ver, val)
	}

	return nil
}

func (e *{{.Name}}Example) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}

func (e *{{.Name}}Example) assignVal(val *yaml.Node) error {
	{{if eq .Type "string" -}}
	if err := xyml.RequireString(val); err != nil {
		return err
	}
	e.value = &val.Value
	{{- else if eq .Type "bool" -}}
	if tmp, err := xyml.ToBoolean(val); err != nil {
		return err
	} else {
		e.value = &tmp
	}
	{{- else if eq .Type "int64" -}}
	if tmp, err := xyml.ToInt(val, 10); err != nil {
		return err
	} else {
		e.value = &tmp
	}
	{{- else if eq .Type "float64" -}}
	if tmp, err := xyml.ToFloat(val); err != nil {
		return err
	} else {
		e.value = &tmp
	}
	{{- else if eq .Type "[]interface{}" -}}
	e.value = make([]interface{}, 0, len(val.Content))
	return xyml.SequenceForEach(val, func(v *yaml.Node) error {
		e.value = append(e.value, v)
		return nil
	})
	{{- else -}}
	var tmp interface{} = *val
	e.value = &tmp
	{{- end}}

	return nil
}
{{end}}