package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"gopkg.in/yaml.v3"
	"io"
)

func NewLibrary() *Library {
	return &Library{
		annotations:     NewAnnotationMap(),
		annotationTypes: NewUntypedMap(),
		extra:           NewAnyMap(),
		uses:            NewStringMap(),
		resourceTypes:   NewUntypedMap(),
		securitySchemes: NewUntypedMap(),
		traits:          NewUntypedMap(),
		types:           NewDataTypeMap(),
	}
}

type Library struct {
	annotations     raml.AnnotationMap
	annotationTypes raml.UntypedMap
	extra           raml.AnyMap
	uses            raml.StringMap
	resourceTypes   raml.UntypedMap
	securitySchemes raml.UntypedMap
	traits          raml.UntypedMap
	types           raml.DataTypeMap
	usage           *string
}

func (l *Library) Annotations() raml.AnnotationMap {
	return l.annotations
}

func (l *Library) SetAnnotations(a raml.AnnotationMap) raml.Library {
	if a == nil {
		return l.UnsetAnnotations()
	}
	l.annotations = a
	return l
}

func (l *Library) UnsetAnnotations() raml.Library {
	l.annotations = NewAnnotationMap()
	return l
}

func (l *Library) AnnotationTypes() raml.UntypedMap {
	return l.annotationTypes
}

func (l *Library) ExtraFacets() raml.AnyMap {
	return l.extra
}

func (l *Library) Uses() raml.StringMap {
	return l.uses
}

func (l *Library) SetUses(uses raml.StringMap) raml.Library {
	if uses == nil {
		return l.UnsetUses()
	}
	l.uses = uses
	return l
}

func (l *Library) UnsetUses() raml.Library {
	l.uses = NewStringMap()
	return l
}

func (l *Library) ResourceTypes() raml.UntypedMap {
	return l.resourceTypes
}

func (l *Library) SetResourceTypes(resTypes raml.UntypedMap) raml.Library {
	if resTypes == nil {
		return l.UnsetResourceTypes()
	}
	l.resourceTypes = resTypes
	return l
}

func (l *Library) UnsetResourceTypes() raml.Library {
	l.resourceTypes = NewUntypedMap()
	return l
}

func (l *Library) SecuritySchemes() raml.UntypedMap {
	return l.securitySchemes
}

func (l *Library) Traits() raml.UntypedMap {
	return l.traits
}

func (l *Library) SetTraits(traits raml.UntypedMap) raml.Library {
	if traits == nil {
		return l.UnsetTraits()
	}
	l.traits = traits
	return l
}

func (l *Library) UnsetTraits() raml.Library {
	l.traits = NewUntypedMap()
	return l
}

func (l *Library) Types() raml.DataTypeMap {
	return l.types
}

func (l *Library) SetTypes(types raml.DataTypeMap) raml.Library {
	if types == nil {
		return l.UnsetTypes()
	}
	l.types = types
	return l
}

func (l *Library) UnsetTypes() raml.Library {
	l.types = NewDataTypeMap()
	return l
}

func (l *Library) Schemas() raml.DataTypeMap {
	return l.Types()
}

func (l *Library) SetSchemas(types raml.DataTypeMap) raml.Library {
	return l.SetTypes(types)
}

func (l *Library) UnsetSchemas() raml.Library {
	return l.UnsetTypes()
}

func (l *Library) Usage() option.String {
	return option.NewMaybeString(l.usage)
}

func (l *Library) SetUsage(usage string) raml.Library {
	l.usage = &usage
	return l
}

func (l *Library) UnsetUsage() raml.Library {
	l.usage = nil
	return l
}

func (l *Library) WriteRAML(w io.Writer) error {
	if _, err := w.Write([]byte(rmeta.HeaderLibrary)); err != nil {
		return err
	}
	enc := yaml.NewEncoder(w)
	enc.SetIndent(2)
	return enc.Encode(l)
}

func (l *Library) UnmarshalYAML(root *yaml.Node) error {
	return xyml.ForEachMap(root, l.assign)
}

func (l Library) MarshalYAML() (interface{}, error) {
	out := NewAnyMap().PutNonNil(rmeta.KeyUsage, l.usage)

	if !l.uses.Empty() {
		out.Put(rmeta.KeyUses, l.uses)
	}

	if !l.securitySchemes.Empty() {
		out.Put(rmeta.KeySecuritySchemes, l.securitySchemes)
	}

	if !l.traits.Empty() {
		out.Put(rmeta.KeyTraits, l.traits)
	}

	if !l.annotationTypes.Empty() {
		out.Put(rmeta.KeyAnnotationTypes, l.annotationTypes)
	}

	l.annotations.ForEach(func(k string, v raml.Annotation) { out.Put(k, v) })

	if !l.resourceTypes.Empty() {
		out.Put(rmeta.KeyResourceTypes, l.resourceTypes)
	}

	l.extra.ForEach(func(k, v interface{}) { out.Put(k, v) })

	if !l.types.Empty() {
		out.Put(rmeta.KeyTypes, l.types)
	}

	return out, nil
}

func (l *Library) assign(k, v *yaml.Node) error {
	switch k.Value {
	case rmeta.KeyAnnotationTypes:
		return l.annotationTypes.UnmarshalRAML(v)
	case rmeta.KeyResourceTypes:
		return l.resourceTypes.UnmarshalRAML(v)
	case rmeta.KeyTypes, rmeta.KeySchemas:
		return l.types.UnmarshalRAML(v)
	case rmeta.KeyTraits:
		return l.traits.UnmarshalRAML(v)
	case rmeta.KeyUses:
		return l.uses.UnmarshalRAML(v)
	case rmeta.KeyUsage:
		return assign.AsStringPtr(v, &l.usage)
	case rmeta.KeySecuritySchemes:
		return l.securitySchemes.UnmarshalRAML(v)
	}

	if xyml.IsString(k) && k.Value[0] == '(' {
		tmp := NewUntypedMap()
		if err := tmp.UnmarshalRAML(v); err != nil {
			return err
		}
		l.annotations.Put(k.Value, tmp)
		return nil
	}

	if val, err := xyml.CastYmlTypeToScalar(k); err != nil {
		return err
	} else {
		l.extra.Put(val, v)
	}

	return nil
}