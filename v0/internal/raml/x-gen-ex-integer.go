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

func NewIntegerExample() *IntegerExample {
	return &IntegerExample{
		annotations: NewAnnotationMap(),
		extra:       NewAnyMap(),
	}
}

type IntegerExample struct {
	displayName *string
	description *string
	annotations raml.AnnotationMap
	value       int64
	strict      bool
	extra       raml.AnyMap
}

func (e *IntegerExample) DisplayName() option.String {
	return option.NewMaybeString(e.displayName)
}

func (e *IntegerExample) SetDisplayName(name string) raml.IntegerExample {
	e.displayName = &name
	return e
}

func (e *IntegerExample) UnsetDisplayName() raml.IntegerExample {
	e.displayName = nil
	return e
}

func (e *IntegerExample) Description() option.String {
	return option.NewMaybeString(e.description)
}

func (e *IntegerExample) SetDescription(desc string) raml.IntegerExample {
	e.description = &desc
	return e
}

func (e *IntegerExample) UnsetDescription() raml.IntegerExample {
	e.description = nil
	return e
}

func (e *IntegerExample) Annotations() raml.AnnotationMap {
	return e.annotations
}

func (e *IntegerExample) SetAnnotations(ann raml.AnnotationMap) raml.IntegerExample {
	if ann == nil {
		return e.UnsetAnnotations()
	}
	e.annotations = ann
	return e
}

func (e *IntegerExample) UnsetAnnotations() raml.IntegerExample {
	e.annotations = NewAnnotationMap()
	return e
}

func (e *IntegerExample) Value() int64 {
	return e.value
}

func (e *IntegerExample) SetValue(val int64) raml.IntegerExample {
	e.value = val
	return e
}

func (e *IntegerExample) Strict() bool {
	return e.strict
}

func (e *IntegerExample) SetStrict(b bool) raml.IntegerExample {
	e.strict = b
	return e
}

func (e *IntegerExample) ExtraFacets() raml.AnyMap {
	return e.extra
}

func (e *IntegerExample) UnmarshalRAML(value *yaml.Node) error {

	if xyml.IsMap(value) {
		return xyml.ForEachMap(value, e.assign)
	}

	return e.assignVal(value)
}

func (e *IntegerExample) MarshalRAML(out raml.AnyMap) (bool, error) {
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

func (e *IntegerExample) assign(key, val *yaml.Node) error {
	logrus.Trace("internal.IntegerExample.assign")

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

func (e *IntegerExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}

func (e *IntegerExample) assignVal(val *yaml.Node) error {
	if tmp, err := xyml.ToInt64(val); err != nil {
		return err
	} else {
		e.value = tmp
	}
	return nil
}