package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// NewStringExample returns a new internal implementation of the
// raml.StringExample interface.
//
// Generated @ 2023-01-17T11:42:46.993195814-05:00
func NewStringExample() *StringExample {
	return &StringExample{
		annotations: raml.NewAnnotationMap(0).SerializeOrdered(false),
		extra:       raml.NewAnyMap(0).SerializeOrdered(false),
		strict:      rmeta.ExampleDefaultStrict,
	}
}

// StringExample is a generated internal implementation of the
// raml.StringExample interface.
type StringExample struct {
	displayName *string
	description *string
	annotations raml.AnnotationMap
	value       *string
	strict      bool
	extra       raml.AnyMap
}

func (e *StringExample) DisplayName() option.String {
	return option.NewMaybeString(e.displayName)
}

func (e *StringExample) SetDisplayName(name string) raml.StringExample {
	e.displayName = &name
	return e
}

func (e *StringExample) UnsetDisplayName() raml.StringExample {
	e.displayName = nil
	return e
}

func (e *StringExample) Description() option.String {
	return option.NewMaybeString(e.description)
}

func (e *StringExample) SetDescription(desc string) raml.StringExample {
	e.description = &desc
	return e
}

func (e *StringExample) UnsetDescription() raml.StringExample {
	e.description = nil
	return e
}

func (e *StringExample) Annotations() raml.AnnotationMap {
	return e.annotations
}

func (e *StringExample) SetAnnotations(ann raml.AnnotationMap) raml.StringExample {
	if ann == nil {
		return e.UnsetAnnotations()
	}
	e.annotations = ann
	return e
}

func (e *StringExample) UnsetAnnotations() raml.StringExample {
	e.annotations = raml.NewAnnotationMap(0)
	return e
}

func (e *StringExample) Value() option.String {
	return option.NewMaybeString(e.value)
}

func (e *StringExample) SetValue(val string) raml.StringExample {
	e.value = &val
	return e
}

func (e *StringExample) UnsetValue() raml.StringExample {
	e.value = nil
	return e
}

func (e *StringExample) Strict() bool {
	return e.strict
}

func (e *StringExample) SetStrict(b bool) raml.StringExample {
	e.strict = b
	return e
}

func (e *StringExample) ExtraFacets() raml.AnyMap {
	return e.extra
}

func (e StringExample) ToYAML() (*yaml.Node, error) {
	out := raml.NewAnyMap(4 + e.annotations.Len() + e.extra.Len()).SerializeOrdered(false)
	if _, err := e.MarshalRAML(out); err != nil {
		return nil, err
	}
	return out.ToYAML()
}

func (e *StringExample) UnmarshalRAML(value *yaml.Node) error {

	if xyml.IsMap(value) {
		return xyml.MapForEach(value, e.assign)
	}

	return e.assignVal(value)
}

func (e *StringExample) MarshalRAML(out raml.AnyMap) (bool, error) {
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

func (e *StringExample) assign(key, val *yaml.Node) error {
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

func (e *StringExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}

func (e *StringExample) assignVal(val *yaml.Node) error {
	if err := xyml.RequireString(val); err != nil {
		return err
	}
	e.value = &val.Value

	return nil
}
