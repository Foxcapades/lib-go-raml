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

func NewFileExample(log *logrus.Entry) *FileExample {
	return &FileExample{
		log:         xlog.WithType(log, "internal.FileExample"),
		annotations: NewAnnotationMap(log),
		extra:       NewAnyMap(log),
	}
}

type FileExample struct {
	log *logrus.Entry

	displayName *string
	description *string
	annotations raml.AnnotationMap
	value       interface{}
	strict      bool
	extra       raml.AnyMap
}

func (e *FileExample) DisplayName() option.String {
	return option.NewMaybeString(e.displayName)
}

func (e *FileExample) SetDisplayName(name string) raml.FileExample {
	e.displayName = &name
	return e
}

func (e *FileExample) UnsetDisplayName() raml.FileExample {
	e.displayName = nil
	return e
}

func (e *FileExample) Description() option.String {
	return option.NewMaybeString(e.description)
}

func (e *FileExample) SetDescription(desc string) raml.FileExample {
	e.description = &desc
	return e
}

func (e *FileExample) UnsetDescription() raml.FileExample {
	e.description = nil
	return e
}

func (e *FileExample) Annotations() raml.AnnotationMap {
	return e.annotations
}

func (e *FileExample) SetAnnotations(ann raml.AnnotationMap) raml.FileExample {
	if ann == nil {
		return e.UnsetAnnotations()
	}
	e.annotations = ann
	return e
}

func (e *FileExample) UnsetAnnotations() raml.FileExample {
	e.annotations = NewAnnotationMap(e.log)
	return e
}

func (e *FileExample) Value() interface{} {
	return e.value
}

func (e *FileExample) SetValue(val interface{}) raml.FileExample {
	e.value = val
	return e
}

func (e *FileExample) Strict() bool {
	return e.strict
}

func (e *FileExample) SetStrict(b bool) raml.FileExample {
	e.strict = b
	return e
}

func (e *FileExample) ExtraFacets() raml.AnyMap {
	return e.extra
}

func (e *FileExample) UnmarshalRAML(val interface{}, log *logrus.Entry) error {
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

func (e *FileExample) MarshalRAML(out raml.AnyMap) (bool, error) {
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

func (e *FileExample) assign(key, val interface{}, log *logrus.Entry) error {
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

func (e *FileExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}
