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

func NewArrayExample(log *logrus.Entry) *ArrayExample {
	return &ArrayExample{
		log:         xlog.WithType(log, "internal.ArrayExample"),
		annotations: NewAnnotationMap(log),
		extra:       NewAnyMap(log),
	}
}

type ArrayExample struct {
	log *logrus.Entry

	displayName *string
	description *string
	annotations raml.AnnotationMap
	value       []interface{}
	strict      bool
	extra       raml.AnyMap
}

func (e *ArrayExample) DisplayName() option.String {
	return option.NewMaybeString(e.displayName)
}

func (e *ArrayExample) SetDisplayName(name string) raml.ArrayExample {
	e.displayName = &name
	return e
}

func (e *ArrayExample) UnsetDisplayName() raml.ArrayExample {
	e.displayName = nil
	return e
}

func (e *ArrayExample) Description() option.String {
	return option.NewMaybeString(e.description)
}

func (e *ArrayExample) SetDescription(desc string) raml.ArrayExample {
	e.description = &desc
	return e
}

func (e *ArrayExample) UnsetDescription() raml.ArrayExample {
	e.description = nil
	return e
}

func (e *ArrayExample) Annotations() raml.AnnotationMap {
	return e.annotations
}

func (e *ArrayExample) SetAnnotations(ann raml.AnnotationMap) raml.ArrayExample {
	if ann == nil {
		return e.UnsetAnnotations()
	}
	e.annotations = ann
	return e
}

func (e *ArrayExample) UnsetAnnotations() raml.ArrayExample {
	e.annotations = NewAnnotationMap(e.log)
	return e
}

func (e *ArrayExample) Value() []interface{} {
	return e.value
}

func (e *ArrayExample) SetValue(val []interface{}) raml.ArrayExample {
	e.value = val
	return e
}

func (e *ArrayExample) Strict() bool {
	return e.strict
}

func (e *ArrayExample) SetStrict(b bool) raml.ArrayExample {
	e.strict = b
	return e
}

func (e *ArrayExample) ExtraFacets() raml.AnyMap {
	return e.extra
}

func (e *ArrayExample) UnmarshalRAML(val interface{}, log *logrus.Entry) error {
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

	if tmp, ok := val.([]interface{}); ok {
		e.value = tmp
		return nil
	}

	return nil
}

func (e *ArrayExample) MarshalRAML(out raml.AnyMap) (bool, error) {
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

func (e *ArrayExample) assign(key, val interface{}, log *logrus.Entry) error {
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
		if tmp, ok := val.([]interface{}); ok{
			e.value = tmp
			return nil
		}
		return xlog.Errorf(log, "invalid example value for Array types.  expected \"[]interface{}\", got %s", reflect.TypeOf(val))
	}

	e.extra.Put(str, val)
	return nil
}

func (e *ArrayExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}
