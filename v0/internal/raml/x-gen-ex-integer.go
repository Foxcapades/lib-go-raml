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

func NewIntegerExample(log *logrus.Entry) *IntegerExample {
	return &IntegerExample{
		log:         xlog.WithType(log, "internal.IntegerExample"),
		annotations: NewAnnotationMap(log),
		extra:       NewAnyMap(log),
	}
}

type IntegerExample struct {
	log *logrus.Entry

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
	e.annotations = NewAnnotationMap(e.log)
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

func (e *IntegerExample) UnmarshalRAML(val interface{}, log *logrus.Entry) error {
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

	if tmp, ok := val.(int64); ok {
		e.value = tmp
		return nil
	}

	return nil
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

func (e *IntegerExample) assign(key, val interface{}, log *logrus.Entry) error {
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
		if tmp, ok := val.(int64); ok{
			e.value = tmp
			return nil
		}
		return xlog.Errorf(log, "invalid example value for Integer types.  expected \"int64\", got %s", reflect.TypeOf(val))
	}

	e.extra.Put(str, val)
	return nil
}

func (e *IntegerExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}
