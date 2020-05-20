package raml

import (
	"reflect"

	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func NewBoolExample(log *logrus.Entry) *BoolExample {
	return &BoolExample{
		log:         xlog.WithType(log, "internal.BoolExample"),
		annotations: NewAnnotationMap(log),
		extra:       NewAnyMap(log),
	}
}

type BoolExample struct {
	log *logrus.Entry

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
	e.annotations = NewAnnotationMap(e.log)
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

func (e *BoolExample) UnmarshalRAML(val interface{}, log *logrus.Entry) error {
	if tmp, ok := val.(yaml.MapSlice); ok {
		for i := range tmp {
			row := &tmp[i]
			l2 := xlog.AddPath(e.log, row.Key)

			if err := e.assign(row.Key, row.Value, l2); err != nil {
				return xlog.Error(l2, err)
			}
		}
		return nil
	}

	if tmp, ok := val.(bool); ok {
		e.value = tmp
		return nil
	}

	return nil
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

func (e *BoolExample) assign(key, val interface{}, log *logrus.Entry) error {
	str, ok := key.(string)

	if !ok {
		e.extra.Put(key, val)
		return nil
	}

	if str[0] == '(' {
		tmp := NewAnnotation(log)
		if err := tmp.UnmarshalRAML(val, log); err != nil {
			return xlog.Error(log, err)
		}
		e.annotations.Put(str, tmp)
		return nil
	}

	switch str {
	case rmeta.KeyDisplayName:
		return assign.AsStringPtr(val, &e.displayName, log)
	case rmeta.KeyDescription:
		return assign.AsStringPtr(val, &e.description, log)
	case rmeta.KeyStrict:
		return assign.AsBool(val, &e.strict, log)
	case rmeta.KeyValue:
		if tmp, ok := val.(bool); ok{
			e.value = tmp
			return nil
		}
		return xlog.Errorf(log, "invalid example value for Bool types.  expected \"bool\", got %s", reflect.TypeOf(val))
	}

	e.extra.Put(str, val)
	return nil
}

func (e *BoolExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}
