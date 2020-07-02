package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// NewAnyExample returns a new internal implementation of the
// raml.AnyExample interface.
//
// Generated @ 2020-07-02T14:19:59.953300998-04:00
func NewAnyExample() *AnyExample {
	return &AnyExample{
		annotations: raml.NewAnnotationMap(0),
		extra:       raml.NewAnyMap(0),
		strict:      rmeta.ExampleDefaultStrict,
	}
}

// AnyExample is a generated internal implementation of the
// raml.AnyExample interface.
type AnyExample struct {
	displayName *string
	description *string
	annotations raml.AnnotationMap
	value       *interface{}
	strict      bool
	extra       raml.AnyMap
}

func (e *AnyExample) DisplayName() option.String {
	return option.NewMaybeString(e.displayName)
}

func (e *AnyExample) SetDisplayName(name string) raml.AnyExample {
	e.displayName = &name
	return e
}

func (e *AnyExample) UnsetDisplayName() raml.AnyExample {
	e.displayName = nil
	return e
}

func (e *AnyExample) Description() option.String {
	return option.NewMaybeString(e.description)
}

func (e *AnyExample) SetDescription(desc string) raml.AnyExample {
	e.description = &desc
	return e
}

func (e *AnyExample) UnsetDescription() raml.AnyExample {
	e.description = nil
	return e
}

func (e *AnyExample) Annotations() raml.AnnotationMap {
	return e.annotations
}

func (e *AnyExample) SetAnnotations(ann raml.AnnotationMap) raml.AnyExample {
	if ann == nil {
		return e.UnsetAnnotations()
	}
	e.annotations = ann
	return e
}

func (e *AnyExample) UnsetAnnotations() raml.AnyExample {
	e.annotations = raml.NewAnnotationMap(0)
	return e
}

func (e *AnyExample) Value() option.Untyped {
	return option.NewMaybeUntyped(e.value)
}

func (e *AnyExample) SetValue(val interface{}) raml.AnyExample {
	e.value = &val
	return e
}

func (e *AnyExample) UnsetValue() raml.AnyExample {
	e.value = nil
	return e
}

func (e *AnyExample) Strict() bool {
	return e.strict
}

func (e *AnyExample) SetStrict(b bool) raml.AnyExample {
	e.strict = b
	return e
}

func (e *AnyExample) ExtraFacets() raml.AnyMap {
	return e.extra
}

func (e *AnyExample) UnmarshalRAML(value *yaml.Node) error {

	if xyml.IsMap(value) {
		return xyml.MapForEach(value, e.assign)
	}

	return e.assignVal(value)
}

func (e *AnyExample) MarshalRAML(out raml.AnyMap) (bool, error) {
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

func (e *AnyExample) assign(key, val *yaml.Node) error {
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

func (e *AnyExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}

func (e *AnyExample) assignVal(val *yaml.Node) error {
	var tmp interface{} = *val
	e.value = &tmp

	return nil
}