package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"gopkg.in/yaml.v3"
)

// NewObjectExample returns a new internal implementation of the
// raml.ObjectExample interface.
//
// Generated @ 2020-05-20T21:46:00.638880955-04:00
func NewObjectExample() *ObjectExample {
	return &ObjectExample{
		annotations: NewAnnotationMap(),
		extra:       NewAnyMap(),
	}
}

// ObjectExample is a generated internal implementation of the
// raml.ObjectExample interface.
type ObjectExample struct {
	displayName *string
	description *string
	annotations raml.AnnotationMap
	value       interface{}
	strict      bool
	extra       raml.AnyMap
}

func (e *ObjectExample) DisplayName() option.String {
	return option.NewMaybeString(e.displayName)
}

func (e *ObjectExample) SetDisplayName(name string) raml.ObjectExample {
	e.displayName = &name
	return e
}

func (e *ObjectExample) UnsetDisplayName() raml.ObjectExample {
	e.displayName = nil
	return e
}

func (e *ObjectExample) Description() option.String {
	return option.NewMaybeString(e.description)
}

func (e *ObjectExample) SetDescription(desc string) raml.ObjectExample {
	e.description = &desc
	return e
}

func (e *ObjectExample) UnsetDescription() raml.ObjectExample {
	e.description = nil
	return e
}

func (e *ObjectExample) Annotations() raml.AnnotationMap {
	return e.annotations
}

func (e *ObjectExample) SetAnnotations(ann raml.AnnotationMap) raml.ObjectExample {
	if ann == nil {
		return e.UnsetAnnotations()
	}
	e.annotations = ann
	return e
}

func (e *ObjectExample) UnsetAnnotations() raml.ObjectExample {
	e.annotations = NewAnnotationMap()
	return e
}

func (e *ObjectExample) Value() interface{} {
	return e.value
}

func (e *ObjectExample) SetValue(val interface{}) raml.ObjectExample {
	e.value = val
	return e
}

func (e *ObjectExample) Strict() bool {
	return e.strict
}

func (e *ObjectExample) SetStrict(b bool) raml.ObjectExample {
	e.strict = b
	return e
}

func (e *ObjectExample) ExtraFacets() raml.AnyMap {
	return e.extra
}

func (e *ObjectExample) UnmarshalRAML(value *yaml.Node) error {

	if xyml.IsMap(value) {
		return xyml.ForEachMap(value, e.assign)
	}

	return e.assignVal(value)
}

func (e *ObjectExample) MarshalRAML(out raml.AnyMap) (bool, error) {
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

func (e *ObjectExample) assign(key, val *yaml.Node) error {
	if !xyml.IsString(key) {
		if ver, err := xyml.CastYmlTypeToScalar(key); err != nil {
			return err
		} else {
			e.extra.Put(ver, val)
		}
		return nil
	}

	if key.Value[0] == '(' {
		tmp := NewAnnotation()
		if err := tmp.UnmarshalRAML(val); err != nil {
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

	if ver, err := xyml.CastYmlTypeToScalar(key); err != nil {
		return err
	} else {
		e.extra.Put(ver, val)
	}

	return nil
}

func (e *ObjectExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}

func (e *ObjectExample) assignVal(val *yaml.Node) error {
	e.value = val

	return nil
}
