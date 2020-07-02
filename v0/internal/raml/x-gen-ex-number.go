package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// NewNumberExample returns a new internal implementation of the
// raml.NumberExample interface.
//
// Generated @ 2020-07-02T14:19:59.953300998-04:00
func NewNumberExample() *NumberExample {
	return &NumberExample{
		annotations: raml.NewAnnotationMap(0),
		extra:       raml.NewAnyMap(0),
		strict:      rmeta.ExampleDefaultStrict,
	}
}

// NumberExample is a generated internal implementation of the
// raml.NumberExample interface.
type NumberExample struct {
	displayName *string
	description *string
	annotations raml.AnnotationMap
	value       *float64
	strict      bool
	extra       raml.AnyMap
}

func (e *NumberExample) DisplayName() option.String {
	return option.NewMaybeString(e.displayName)
}

func (e *NumberExample) SetDisplayName(name string) raml.NumberExample {
	e.displayName = &name
	return e
}

func (e *NumberExample) UnsetDisplayName() raml.NumberExample {
	e.displayName = nil
	return e
}

func (e *NumberExample) Description() option.String {
	return option.NewMaybeString(e.description)
}

func (e *NumberExample) SetDescription(desc string) raml.NumberExample {
	e.description = &desc
	return e
}

func (e *NumberExample) UnsetDescription() raml.NumberExample {
	e.description = nil
	return e
}

func (e *NumberExample) Annotations() raml.AnnotationMap {
	return e.annotations
}

func (e *NumberExample) SetAnnotations(ann raml.AnnotationMap) raml.NumberExample {
	if ann == nil {
		return e.UnsetAnnotations()
	}
	e.annotations = ann
	return e
}

func (e *NumberExample) UnsetAnnotations() raml.NumberExample {
	e.annotations = raml.NewAnnotationMap(0)
	return e
}

func (e *NumberExample) Value() option.Float64 {
	return option.NewMaybeFloat64(e.value)
}

func (e *NumberExample) SetValue(val float64) raml.NumberExample {
	e.value = &val
	return e
}

func (e *NumberExample) UnsetValue() raml.NumberExample {
	e.value = nil
	return e
}

func (e *NumberExample) Strict() bool {
	return e.strict
}

func (e *NumberExample) SetStrict(b bool) raml.NumberExample {
	e.strict = b
	return e
}

func (e *NumberExample) ExtraFacets() raml.AnyMap {
	return e.extra
}

func (e *NumberExample) UnmarshalRAML(value *yaml.Node) error {

	if xyml.IsMap(value) {
		return xyml.MapForEach(value, e.assign)
	}

	return e.assignVal(value)
}

func (e *NumberExample) MarshalRAML(out raml.AnyMap) (bool, error) {
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

func (e *NumberExample) assign(key, val *yaml.Node) error {
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

func (e *NumberExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}

func (e *NumberExample) assignVal(val *yaml.Node) error {
	if tmp, err := xyml.ToFloat(val); err != nil {
		return err
	} else {
		e.value = &tmp
	}

	return nil
}
