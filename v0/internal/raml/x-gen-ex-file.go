package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"gopkg.in/yaml.v3"
)

// NewFileExample returns a new internal implementation of the
// raml.FileExample interface.
//
// Generated @ 2020-05-20T21:46:00.638880955-04:00
func NewFileExample() *FileExample {
	return &FileExample{
		annotations: NewAnnotationMap(),
		extra:       NewAnyMap(),
	}
}

// FileExample is a generated internal implementation of the
// raml.FileExample interface.
type FileExample struct {
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
	e.annotations = NewAnnotationMap()
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

func (e *FileExample) UnmarshalRAML(value *yaml.Node) error {

	if xyml.IsMap(value) {
		return xyml.ForEachMap(value, e.assign)
	}

	return e.assignVal(value)
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

func (e *FileExample) assign(key, val *yaml.Node) error {
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

func (e *FileExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}

func (e *FileExample) assignVal(val *yaml.Node) error {
	e.value = val

	return nil
}
