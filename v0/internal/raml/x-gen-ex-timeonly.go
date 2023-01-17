package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// NewTimeOnlyExample returns a new internal implementation of the
// raml.TimeOnlyExample interface.
//
// Generated @ 2023-01-17T11:42:46.993195814-05:00
func NewTimeOnlyExample() *TimeOnlyExample {
	return &TimeOnlyExample{
		annotations: raml.NewAnnotationMap(0).SerializeOrdered(false),
		extra:       raml.NewAnyMap(0).SerializeOrdered(false),
		strict:      rmeta.ExampleDefaultStrict,
	}
}

// TimeOnlyExample is a generated internal implementation of the
// raml.TimeOnlyExample interface.
type TimeOnlyExample struct {
	displayName *string
	description *string
	annotations raml.AnnotationMap
	value       *string
	strict      bool
	extra       raml.AnyMap
}

func (e *TimeOnlyExample) DisplayName() option.String {
	return option.NewMaybeString(e.displayName)
}

func (e *TimeOnlyExample) SetDisplayName(name string) raml.TimeOnlyExample {
	e.displayName = &name
	return e
}

func (e *TimeOnlyExample) UnsetDisplayName() raml.TimeOnlyExample {
	e.displayName = nil
	return e
}

func (e *TimeOnlyExample) Description() option.String {
	return option.NewMaybeString(e.description)
}

func (e *TimeOnlyExample) SetDescription(desc string) raml.TimeOnlyExample {
	e.description = &desc
	return e
}

func (e *TimeOnlyExample) UnsetDescription() raml.TimeOnlyExample {
	e.description = nil
	return e
}

func (e *TimeOnlyExample) Annotations() raml.AnnotationMap {
	return e.annotations
}

func (e *TimeOnlyExample) SetAnnotations(ann raml.AnnotationMap) raml.TimeOnlyExample {
	if ann == nil {
		return e.UnsetAnnotations()
	}
	e.annotations = ann
	return e
}

func (e *TimeOnlyExample) UnsetAnnotations() raml.TimeOnlyExample {
	e.annotations = raml.NewAnnotationMap(0)
	return e
}

func (e *TimeOnlyExample) Value() option.String {
	return option.NewMaybeString(e.value)
}

func (e *TimeOnlyExample) SetValue(val string) raml.TimeOnlyExample {
	e.value = &val
	return e
}

func (e *TimeOnlyExample) UnsetValue() raml.TimeOnlyExample {
	e.value = nil
	return e
}

func (e *TimeOnlyExample) Strict() bool {
	return e.strict
}

func (e *TimeOnlyExample) SetStrict(b bool) raml.TimeOnlyExample {
	e.strict = b
	return e
}

func (e *TimeOnlyExample) ExtraFacets() raml.AnyMap {
	return e.extra
}

func (e TimeOnlyExample) ToYAML() (*yaml.Node, error) {
	out := raml.NewAnyMap(4 + e.annotations.Len() + e.extra.Len()).SerializeOrdered(false)
	if _, err := e.MarshalRAML(out); err != nil {
		return nil, err
	}
	return out.ToYAML()
}

func (e *TimeOnlyExample) UnmarshalRAML(value *yaml.Node) error {

	if xyml.IsMap(value) {
		return xyml.MapForEach(value, e.assign)
	}

	return e.assignVal(value)
}

func (e *TimeOnlyExample) MarshalRAML(out raml.AnyMap) (bool, error) {
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

func (e *TimeOnlyExample) assign(key, val *yaml.Node) error {
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

func (e *TimeOnlyExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}

func (e *TimeOnlyExample) assignVal(val *yaml.Node) error {
	if err := xyml.RequireString(val); err != nil {
		return err
	}
	e.value = &val.Value

	return nil
}
