package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// NewUnionExample returns a new internal implementation of the
// raml.UnionExample interface.
//
// Generated @ 2020-05-20T20:54:25.600656868-04:00
func NewUnionExample() *UnionExample {
	return &UnionExample{
		annotations: NewAnnotationMap(),
		extra:       NewAnyMap(),
	}
}

// UnionExample is a generated internal implementation of the
// raml.UnionExample interface.
type UnionExample struct {
	displayName *string
	description *string
	annotations raml.AnnotationMap
	value       interface{}
	strict      bool
	extra       raml.AnyMap
}

func (e *UnionExample) DisplayName() option.String {
	return option.NewMaybeString(e.displayName)
}

func (e *UnionExample) SetDisplayName(name string) raml.UnionExample {
	e.displayName = &name
	return e
}

func (e *UnionExample) UnsetDisplayName() raml.UnionExample {
	e.displayName = nil
	return e
}

func (e *UnionExample) Description() option.String {
	return option.NewMaybeString(e.description)
}

func (e *UnionExample) SetDescription(desc string) raml.UnionExample {
	e.description = &desc
	return e
}

func (e *UnionExample) UnsetDescription() raml.UnionExample {
	e.description = nil
	return e
}

func (e *UnionExample) Annotations() raml.AnnotationMap {
	return e.annotations
}

func (e *UnionExample) SetAnnotations(ann raml.AnnotationMap) raml.UnionExample {
	if ann == nil {
		return e.UnsetAnnotations()
	}
	e.annotations = ann
	return e
}

func (e *UnionExample) UnsetAnnotations() raml.UnionExample {
	e.annotations = NewAnnotationMap()
	return e
}

func (e *UnionExample) Value() interface{} {
	return e.value
}

func (e *UnionExample) SetValue(val interface{}) raml.UnionExample {
	e.value = val
	return e
}

func (e *UnionExample) Strict() bool {
	return e.strict
}

func (e *UnionExample) SetStrict(b bool) raml.UnionExample {
	e.strict = b
	return e
}

func (e *UnionExample) ExtraFacets() raml.AnyMap {
	return e.extra
}

func (e *UnionExample) UnmarshalRAML(value *yaml.Node) error {

	if xyml.IsMap(value) {
		return xyml.ForEachMap(value, e.assign)
	}

	return e.assignVal(value)
}

func (e *UnionExample) MarshalRAML(out raml.AnyMap) (bool, error) {
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

func (e *UnionExample) assign(key, val *yaml.Node) error {
	logrus.Trace("internal.UnionExample.assign")

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

func (e *UnionExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}

func (e *UnionExample) assignVal(val *yaml.Node) error {
	e.value = val

	return nil
}
