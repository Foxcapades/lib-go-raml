package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// NewDatetimeExample returns a new internal implementation of the
// raml.DatetimeExample interface.
//
// Generated @ 2020-07-08T13:31:34.46078727-04:00
func NewDatetimeExample() *DatetimeExample {
	return &DatetimeExample{
		annotations: raml.NewAnnotationMap(0).SerializeOrdered(false),
		extra:       raml.NewAnyMap(0).SerializeOrdered(false),
		strict:      rmeta.ExampleDefaultStrict,
	}
}

// DatetimeExample is a generated internal implementation of the
// raml.DatetimeExample interface.
type DatetimeExample struct {
	displayName *string
	description *string
	annotations raml.AnnotationMap
	value       *string
	strict      bool
	extra       raml.AnyMap
}

func (e *DatetimeExample) DisplayName() option.String {
	return option.NewMaybeString(e.displayName)
}

func (e *DatetimeExample) SetDisplayName(name string) raml.DatetimeExample {
	e.displayName = &name
	return e
}

func (e *DatetimeExample) UnsetDisplayName() raml.DatetimeExample {
	e.displayName = nil
	return e
}

func (e *DatetimeExample) Description() option.String {
	return option.NewMaybeString(e.description)
}

func (e *DatetimeExample) SetDescription(desc string) raml.DatetimeExample {
	e.description = &desc
	return e
}

func (e *DatetimeExample) UnsetDescription() raml.DatetimeExample {
	e.description = nil
	return e
}

func (e *DatetimeExample) Annotations() raml.AnnotationMap {
	return e.annotations
}

func (e *DatetimeExample) SetAnnotations(ann raml.AnnotationMap) raml.DatetimeExample {
	if ann == nil {
		return e.UnsetAnnotations()
	}
	e.annotations = ann
	return e
}

func (e *DatetimeExample) UnsetAnnotations() raml.DatetimeExample {
	e.annotations = raml.NewAnnotationMap(0)
	return e
}

func (e *DatetimeExample) Value() option.String {
	return option.NewMaybeString(e.value)
}

func (e *DatetimeExample) SetValue(val string) raml.DatetimeExample {
	e.value = &val
	return e
}

func (e *DatetimeExample) UnsetValue() raml.DatetimeExample {
	e.value = nil
	return e
}

func (e *DatetimeExample) Strict() bool {
	return e.strict
}

func (e *DatetimeExample) SetStrict(b bool) raml.DatetimeExample {
	e.strict = b
	return e
}

func (e *DatetimeExample) ExtraFacets() raml.AnyMap {
	return e.extra
}

func (e DatetimeExample) ToYAML() (*yaml.Node, error) {
	out := raml.NewAnyMap(4 + e.annotations.Len() + e.extra.Len())
	if _, err := e.MarshalRAML(out); err != nil {
		return nil, err
	}
	return out.ToYAML()
}

func (e *DatetimeExample) UnmarshalRAML(value *yaml.Node) error {

	if xyml.IsMap(value) {
		return xyml.MapForEach(value, e.assign)
	}

	return e.assignVal(value)
}

func (e *DatetimeExample) MarshalRAML(out raml.AnyMap) (bool, error) {
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

func (e *DatetimeExample) assign(key, val *yaml.Node) error {
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

func (e *DatetimeExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}

func (e *DatetimeExample) assignVal(val *yaml.Node) error {
	if err := xyml.RequireString(val); err != nil {
		return err
	}
	e.value = &val.Value

	return nil
}
