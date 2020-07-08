package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// NewCustomExample returns a new internal implementation of the
// raml.CustomExample interface.
//
// Generated @ 2020-07-08T13:31:34.46078727-04:00
func NewCustomExample() *CustomExample {
	return &CustomExample{
		annotations: raml.NewAnnotationMap(0).SerializeOrdered(false),
		extra:       raml.NewAnyMap(0).SerializeOrdered(false),
		strict:      rmeta.ExampleDefaultStrict,
	}
}

// CustomExample is a generated internal implementation of the
// raml.CustomExample interface.
type CustomExample struct {
	displayName *string
	description *string
	annotations raml.AnnotationMap
	value       *interface{}
	strict      bool
	extra       raml.AnyMap
}

func (e *CustomExample) DisplayName() option.String {
	return option.NewMaybeString(e.displayName)
}

func (e *CustomExample) SetDisplayName(name string) raml.CustomExample {
	e.displayName = &name
	return e
}

func (e *CustomExample) UnsetDisplayName() raml.CustomExample {
	e.displayName = nil
	return e
}

func (e *CustomExample) Description() option.String {
	return option.NewMaybeString(e.description)
}

func (e *CustomExample) SetDescription(desc string) raml.CustomExample {
	e.description = &desc
	return e
}

func (e *CustomExample) UnsetDescription() raml.CustomExample {
	e.description = nil
	return e
}

func (e *CustomExample) Annotations() raml.AnnotationMap {
	return e.annotations
}

func (e *CustomExample) SetAnnotations(ann raml.AnnotationMap) raml.CustomExample {
	if ann == nil {
		return e.UnsetAnnotations()
	}
	e.annotations = ann
	return e
}

func (e *CustomExample) UnsetAnnotations() raml.CustomExample {
	e.annotations = raml.NewAnnotationMap(0)
	return e
}

func (e *CustomExample) Value() option.Untyped {
	return option.NewMaybeUntyped(e.value)
}

func (e *CustomExample) SetValue(val interface{}) raml.CustomExample {
	e.value = &val
	return e
}

func (e *CustomExample) UnsetValue() raml.CustomExample {
	e.value = nil
	return e
}

func (e *CustomExample) Strict() bool {
	return e.strict
}

func (e *CustomExample) SetStrict(b bool) raml.CustomExample {
	e.strict = b
	return e
}

func (e *CustomExample) ExtraFacets() raml.AnyMap {
	return e.extra
}

func (e CustomExample) ToYAML() (*yaml.Node, error) {
	out := raml.NewAnyMap(4 + e.annotations.Len() + e.extra.Len())
	if _, err := e.MarshalRAML(out); err != nil {
		return nil, err
	}
	return out.ToYAML()
}

func (e *CustomExample) UnmarshalRAML(value *yaml.Node) error {

	if xyml.IsMap(value) {
		return xyml.MapForEach(value, e.assign)
	}

	return e.assignVal(value)
}

func (e *CustomExample) MarshalRAML(out raml.AnyMap) (bool, error) {
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

func (e *CustomExample) assign(key, val *yaml.Node) error {
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

func (e *CustomExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}

func (e *CustomExample) assignVal(val *yaml.Node) error {
	var tmp interface{} = *val
	e.value = &tmp

	return nil
}
