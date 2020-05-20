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

func NewDateOnlyExample(log *logrus.Entry) *DateOnlyExample {
	return &DateOnlyExample{
		log:         xlog.WithType(log, "internal.DateOnlyExample"),
		annotations: NewAnnotationMap(log),
		extra:       NewAnyMap(log),
	}
}

type DateOnlyExample struct {
	log *logrus.Entry

	displayName *string
	description *string
	annotations raml.AnnotationMap
	value       string
	strict      bool
	extra       raml.AnyMap
}

func (e *DateOnlyExample) DisplayName() option.String {
	return option.NewMaybeString(e.displayName)
}

func (e *DateOnlyExample) SetDisplayName(name string) raml.DateOnlyExample {
	e.displayName = &name
	return e
}

func (e *DateOnlyExample) UnsetDisplayName() raml.DateOnlyExample {
	e.displayName = nil
	return e
}

func (e *DateOnlyExample) Description() option.String {
	return option.NewMaybeString(e.description)
}

func (e *DateOnlyExample) SetDescription(desc string) raml.DateOnlyExample {
	e.description = &desc
	return e
}

func (e *DateOnlyExample) UnsetDescription() raml.DateOnlyExample {
	e.description = nil
	return e
}

func (e *DateOnlyExample) Annotations() raml.AnnotationMap {
	return e.annotations
}

func (e *DateOnlyExample) SetAnnotations(ann raml.AnnotationMap) raml.DateOnlyExample {
	if ann == nil {
		return e.UnsetAnnotations()
	}
	e.annotations = ann
	return e
}

func (e *DateOnlyExample) UnsetAnnotations() raml.DateOnlyExample {
	e.annotations = NewAnnotationMap(e.log)
	return e
}

func (e *DateOnlyExample) Value() string {
	return e.value
}

func (e *DateOnlyExample) SetValue(val string) raml.DateOnlyExample {
	e.value = val
	return e
}

func (e *DateOnlyExample) Strict() bool {
	return e.strict
}

func (e *DateOnlyExample) SetStrict(b bool) raml.DateOnlyExample {
	e.strict = b
	return e
}

func (e *DateOnlyExample) ExtraFacets() raml.AnyMap {
	return e.extra
}

func (e *DateOnlyExample) UnmarshalRAML(val interface{}, log *logrus.Entry) error {
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

	if tmp, ok := val.(string); ok {
		e.value = tmp
		return nil
	}

	return nil
}

func (e *DateOnlyExample) MarshalRAML(out raml.AnyMap) (bool, error) {
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

func (e *DateOnlyExample) assign(key, val interface{}, log *logrus.Entry) error {
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
		if tmp, ok := val.(string); ok{
			e.value = tmp
			return nil
		}
		return xlog.Errorf(log, "invalid example value for DateOnly types.  expected \"string\", got %s", reflect.TypeOf(val))
	}

	e.extra.Put(str, val)
	return nil
}

func (e *DateOnlyExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}
