package raml

import (
	

	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func NewUnionExample(log *logrus.Entry) *UnionExample {
	return &UnionExample{
		log:         xlog.WithType(log, "internal.UnionExample"),
		annotations: NewAnnotationMap(log),
		extra:       NewAnyMap(log),
	}
}

type UnionExample struct {
	log *logrus.Entry

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
	e.annotations = NewAnnotationMap(e.log)
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

func (e *UnionExample) UnmarshalRAML(val interface{}, log *logrus.Entry) error {
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

	e.value = val

	return nil
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

func (e *UnionExample) assign(key, val interface{}, log *logrus.Entry) error {
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
		e.value = val
	}

	e.extra.Put(str, val)
	return nil
}

func (e *UnionExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}
