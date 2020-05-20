{{- /*gotype: github.com/Foxcapades/lib-go-raml-types/v0/tools/gen/example.exampleProps*/ -}}
{{define "example" -}}
package raml

import (
	{{if not (eq .Name "Custom" "File" "Object" "Union") -}}
	"reflect"
	{{- end}}

	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func New{{.Name}}Example(log *logrus.Entry) *{{.Name}}Example {
	return &{{.Name}}Example{
		log:         xlog.WithType(log, "internal.{{.Name}}Example"),
		annotations: NewAnnotationMap(log),
		extra:       NewAnyMap(log),
	}
}

type {{.Name}}Example struct {
	log *logrus.Entry

	displayName *string
	description *string
	annotations raml.AnnotationMap
	value       {{.Type}}
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
	e.annotations = NewAnnotationMap(e.log)
	return e
}

func (e *{{.Name}}Example) Value() {{.Type}} {
	return e.value
}

func (e *{{.Name}}Example) SetValue(val {{.Type}}) raml.{{.Name}}Example {
	e.value = val
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

func (e *{{.Name}}Example) UnmarshalRAML(val interface{}, log *logrus.Entry) error {
	if tmp, ok := val.(yaml.MapSlice); ok {
		for i := range tmp {
			row := &tmp[i]
			l2 := xlog.AddPath(e.log, row.Key)

			if err := e.assign(row.Key, row.Value, l2); err != nil {
				return xlog.Error(l2, err)
			}
		}
		return nil
	}

	{{if eq .Type "string" "bool" "int64" "float64" "[]interface{}" -}}
	if tmp, ok := val.({{.Type}}); ok {
		e.value = tmp
		return nil
	}
	{{- else -}}
	e.value = val
	{{- end}}

	return nil
}

func (e *{{.Name}}Example) MarshalRAML(out raml.AnyMap) (bool, error) {
	if e.expand() {
		out.PutNonNil(rmeta.KeyDisplayName, e.displayName).
			PutNonNil(rmeta.KeyDescription, e.description).
			Put(rmeta.KeyValue, e.value)

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

func (e *{{.Name}}Example) assign(key, val interface{}, log *logrus.Entry) error {
	str, ok := key.(string)

	if !ok {
		e.extra.Put(key, val)
		return nil
	}

	if str[0] == '(' {
		tmp := NewAnnotation(log)
		if err := tmp.UnmarshalRAML(val, log); err != nil {
			return xlog.Error(log, err)
		}
		e.annotations.Put(str, tmp)
		return nil
	}

	switch str {
	case rmeta.KeyDisplayName:
		return assign.AsStringPtr(val, &e.displayName, log)
	case rmeta.KeyDescription:
		return assign.AsStringPtr(val, &e.description, log)
	case rmeta.KeyStrict:
		return assign.AsBool(val, &e.strict, log)
	case rmeta.KeyValue:
		{{if eq .Type "interface{}" -}}
		e.value = val
		{{- else -}}
		if tmp, ok := val.({{.Type}}); ok{
			e.value = tmp
			return nil
		}
		return xlog.Errorf(log, "invalid example value for {{.Name}} types.  expected \"{{.Type}}\", got %s", reflect.TypeOf(val))
		{{- end}}
	}

	e.extra.Put(str, val)
	return nil
}

func (e *{{.Name}}Example) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}
{{end}}