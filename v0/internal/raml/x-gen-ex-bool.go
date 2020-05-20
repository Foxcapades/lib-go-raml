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

func NewBoolExample() *BoolExample {
	return &BoolExample{
		annotations: NewAnnotationMap(),
		extra:       NewAnyMap(),
	}
}

type BoolExample struct {
	displayName *string
	description *string
	annotations raml.AnnotationMap
	value       bool
	strict      bool
	extra       raml.AnyMap
}

func (e *BoolExample) DisplayName() option.String {
	return option.NewMaybeString(e.displayName)
}

func (e *BoolExample) SetDisplayName(name string) raml.BoolExample {
	e.displayName = &name
	return e
}

func (e *BoolExample) UnsetDisplayName() raml.BoolExample {
	e.displayName = nil
	return e
}

func (e *BoolExample) Description() option.String {
	return option.NewMaybeString(e.description)
}

func (e *BoolExample) SetDescription(desc string) raml.BoolExample {
	e.description = &desc
	return e
}

func (e *BoolExample) UnsetDescription() raml.BoolExample {
	e.description = nil
	return e
}

func (e *BoolExample) Annotations() raml.AnnotationMap {
	return e.annotations
}

func (e *BoolExample) SetAnnotations(ann raml.AnnotationMap) raml.BoolExample {
	if ann == nil {
		return e.UnsetAnnotations()
	}
	e.annotations = ann
	return e
}

func (e *BoolExample) UnsetAnnotations() raml.BoolExample {
	e.annotations = NewAnnotationMap()
	return e
}

func (e *BoolExample) Value() bool {
	return e.value
}

func (e *BoolExample) SetValue(val bool) raml.BoolExample {
	e.value = val
	return e
}

func (e *BoolExample) Strict() bool {
	return e.strict
}

func (e *BoolExample) SetStrict(b bool) raml.BoolExample {
	e.strict = b
	return e
}

func (e *BoolExample) ExtraFacets() raml.AnyMap {
	return e.extra
}

func (e *BoolExample) UnmarshalRAML(value *yaml.Node) error {

	if xyml.IsMap(value) {
		return xyml.ForEachMap(value, e.assign)
	}

	return e.assignVal(value)
}

func (e *BoolExample) MarshalRAML(out raml.AnyMap) (bool, error) {
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

func (e *BoolExample) assign(key, val *yaml.Node) error {
	logrus.Trace("internal.BoolExample.assign")

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

func (e *BoolExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}

func (e *BoolExample) assignVal(val *yaml.Node) error {
	if tmp, err := xyml.ToBool(val); err != nil {
		return err
	} else {
		e.value = tmp
	}
	return nil
}
