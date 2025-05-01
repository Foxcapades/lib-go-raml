package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// NewIncludeType returns a new internal implementation
// of the raml.IncludeType interface.
//
// Generated @ 2025-05-01T01:20:50.297484942-04:00
func NewIncludeType() *IncludeType {
	out := &IncludeType{}

	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeInclude, out)

	return out
}

// IncludeType is a default generated implementation of
// the raml.IncludeType interface
//
// Generated @ 2025-05-01T01:20:50.297484942-04:00
type IncludeType struct {
	*ExtendedDataType

	def  *interface{}
	enum []interface{}
}

func (o *IncludeType) SetType(s string) raml.IncludeType {
	o.schema = s
	return o
}

func (o *IncludeType) Default() option.Untyped {
	return option.NewMaybeUntyped(o.def)

}

func (o *IncludeType) SetDefault(i interface{}) raml.IncludeType {
	o.def = &i
	return o
}

func (o *IncludeType) UnsetDefault() raml.IncludeType {
	o.def = nil
	return o
}

func (o *IncludeType) SetDisplayName(s string) raml.IncludeType {
	o.displayName = &s
	return o
}

func (o *IncludeType) UnsetDisplayName() raml.IncludeType {
	o.displayName = nil
	return o
}

func (o *IncludeType) SetDescription(s string) raml.IncludeType {
	o.description = &s
	return o
}

func (o *IncludeType) UnsetDescription() raml.IncludeType {
	o.description = nil
	return o
}

func (o *IncludeType) SetAnnotations(annotations raml.AnnotationMap) raml.IncludeType {
	if annotations == nil {
		return o.UnsetAnnotations()
	}

	o.hasAnnotations.mp = annotations
	return o
}

func (o *IncludeType) UnsetAnnotations() raml.IncludeType {
	o.hasAnnotations.mp = raml.NewAnnotationMap(0)
	return o
}

func (o *IncludeType) SetFacetDefinitions(facets raml.FacetMap) raml.IncludeType {
	if facets == nil {
		return o.UnsetFacetDefinitions()
	}

	o.facets = facets
	return o
}

func (o *IncludeType) UnsetFacetDefinitions() raml.IncludeType {
	o.facets = raml.NewFacetMap(0)
	return o
}

func (o *IncludeType) SetXML(x raml.XML) raml.IncludeType {
	o.xml = x
	return o
}

func (o *IncludeType) UnsetXML() raml.IncludeType {
	o.xml = nil
	return o
}

func (o *IncludeType) Enum() []interface{} {
	return o.enum
}

func (o *IncludeType) SetEnum(i []interface{}) raml.IncludeType {
	o.enum = i
	return o
}

func (o *IncludeType) UnsetEnum() raml.IncludeType {
	o.enum = nil
	return o
}

func (o *IncludeType) SetExtraFacets(facets raml.AnyMap) raml.IncludeType {
	if facets == nil {
		return o.UnsetExtraFacets()
	}

	o.hasExtra.mp = facets
	return o
}

func (o *IncludeType) UnsetExtraFacets() raml.IncludeType {
	o.hasExtra.mp = raml.NewAnyMap(0)
	return o
}

func (o *IncludeType) SetRequired(b bool) raml.IncludeType {
	o.required = b
	return o
}

func (o *IncludeType) marshal(out raml.AnyMap) error {
	logrus.Trace("internal.IncludeType.marshal")
	out.PutIfNotNil(rmeta.KeyDefault, o.def)

	if err := o.ExtendedDataType.marshal(out); err != nil {
		return err
	}
	out.Put(rmeta.KeyType, &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   "!include",
		Value: o.DataType.schema,
	})
	out.PutIfNotNil(rmeta.KeyEnum, o.enum)

	return nil
}

func (o *IncludeType) assign(key, val *yaml.Node) error {
	switch key.Value {
	case rmeta.KeyEnum:
		return xyml.SequenceForEach(val, func(cur *yaml.Node) error {
			o.enum = append(o.enum, cur)

			return nil
		})
	case rmeta.KeyRequired:
		return assign.AsBool(val, &o.required)
	}

	switch key.Value {
	case rmeta.KeyType, rmeta.KeySchema:
		o.DataType.schema = val.Value
		return nil
	}

	return o.ExtendedDataType.assign(key, val)
}
