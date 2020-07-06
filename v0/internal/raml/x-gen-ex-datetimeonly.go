package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// NewDatetimeOnlyExample returns a new internal implementation of the
// raml.DatetimeOnlyExample interface.
//
// Generated @ 2020-07-06T13:52:18.264181542-04:00
func NewDatetimeOnlyExample() *DatetimeOnlyExample {
	return &DatetimeOnlyExample{
		annotations: raml.NewAnnotationMap(0).SerializeOrdered(false),
		extra:       raml.NewAnyMap(0).SerializeOrdered(false),
		strict:      rmeta.ExampleDefaultStrict,
	}
}

// DatetimeOnlyExample is a generated internal implementation of the
// raml.DatetimeOnlyExample interface.
type DatetimeOnlyExample struct {
	displayName *string
	description *string
	annotations raml.AnnotationMap
	value       *string
	strict      bool
	extra       raml.AnyMap
}

func (e *DatetimeOnlyExample) DisplayName() option.String {
	return option.NewMaybeString(e.displayName)
}

func (e *DatetimeOnlyExample) SetDisplayName(name string) raml.DatetimeOnlyExample {
	e.displayName = &name
	return e
}

func (e *DatetimeOnlyExample) UnsetDisplayName() raml.DatetimeOnlyExample {
	e.displayName = nil
	return e
}

func (e *DatetimeOnlyExample) Description() option.String {
	return option.NewMaybeString(e.description)
}

func (e *DatetimeOnlyExample) SetDescription(desc string) raml.DatetimeOnlyExample {
	e.description = &desc
	return e
}

func (e *DatetimeOnlyExample) UnsetDescription() raml.DatetimeOnlyExample {
	e.description = nil
	return e
}

func (e *DatetimeOnlyExample) Annotations() raml.AnnotationMap {
	return e.annotations
}

func (e *DatetimeOnlyExample) SetAnnotations(ann raml.AnnotationMap) raml.DatetimeOnlyExample {
	if ann == nil {
		return e.UnsetAnnotations()
	}
	e.annotations = ann
	return e
}

func (e *DatetimeOnlyExample) UnsetAnnotations() raml.DatetimeOnlyExample {
	e.annotations = raml.NewAnnotationMap(0)
	return e
}

func (e *DatetimeOnlyExample) Value() option.String {
	return option.NewMaybeString(e.value)
}

func (e *DatetimeOnlyExample) SetValue(val string) raml.DatetimeOnlyExample {
	e.value = &val
	return e
}

func (e *DatetimeOnlyExample) UnsetValue() raml.DatetimeOnlyExample {
	e.value = nil
	return e
}

func (e *DatetimeOnlyExample) Strict() bool {
	return e.strict
}

func (e *DatetimeOnlyExample) SetStrict(b bool) raml.DatetimeOnlyExample {
	e.strict = b
	return e
}

func (e *DatetimeOnlyExample) ExtraFacets() raml.AnyMap {
	return e.extra
}

func (e *DatetimeOnlyExample) UnmarshalRAML(value *yaml.Node) error {

	if xyml.IsMap(value) {
		return xyml.MapForEach(value, e.assign)
	}

	return e.assignVal(value)
}

func (e *DatetimeOnlyExample) MarshalRAML(out raml.AnyMap) (bool, error) {
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

func (e *DatetimeOnlyExample) assign(key, val *yaml.Node) error {
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

func (e *DatetimeOnlyExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}

func (e *DatetimeOnlyExample) assignVal(val *yaml.Node) error {
	if err := xyml.RequireString(val); err != nil {
		return err
	}
	e.value = &val.Value

	return nil
}
