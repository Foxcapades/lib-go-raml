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

// NewObjectType returns a new internal implementation
// of the raml.ObjectType interface.
//
// Generated @ 2020-05-20T18:40:13.095690448-04:00
func NewObjectType() *ObjectType {
	out := &ObjectType{
		examples: NewObjectExampleMap(),
	}

	out.properties = NewPropertyMap()
	out.addtlProps = true

	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeObject, out)

	return out
}

// ObjectType is a default generated implementation of
// the raml.ObjectType interface
//
// Generated @ 2020-05-20T18:40:13.095690448-04:00
type ObjectType struct {
	*ExtendedDataType

	def        *interface{}
	example    raml.ObjectExample
	examples   raml.ObjectExampleMap
	enum       []interface{}
	properties raml.PropertyMap
	minProps   *uint
	maxProps   *uint
	addtlProps bool
	discrim    *string
	discrimVal *interface{}
}

func (o *ObjectType) SetType(s string) raml.ObjectType {
	o.schema = s
	return o
}

func (o *ObjectType) Default() option.Untyped {
	return option.NewMaybeUntyped(o.def)
}

func (o *ObjectType) SetDefault(i interface{}) raml.ObjectType {
	o.def = &i
	return o
}

func (o *ObjectType) UnsetDefault() raml.ObjectType {
	o.def = nil
	return o
}

func (o *ObjectType) Example() raml.ObjectExample {
	return o.example
}

func (o *ObjectType) SetExample(ex raml.ObjectExample) raml.ObjectType {
	o.example = ex
	return o
}

func (o *ObjectType) UnsetExample() raml.ObjectType {
	o.example = nil
	return o
}

func (o *ObjectType) Examples() raml.ObjectExampleMap {
	return o.examples
}

func (o *ObjectType) SetExamples(examples raml.ObjectExampleMap) raml.ObjectType {
	if examples == nil {
		return o.UnsetExamples()
	}
	o.examples = examples
	return o
}

func (o *ObjectType) UnsetExamples() raml.ObjectType {
	o.examples = NewObjectExampleMap()
	return o
}

func (o *ObjectType) SetDisplayName(s string) raml.ObjectType {
	o.displayName = &s
	return o
}

func (o *ObjectType) UnsetDisplayName() raml.ObjectType {
	o.displayName = nil
	return o
}

func (o *ObjectType) SetDescription(s string) raml.ObjectType {
	o.description = &s
	return o
}

func (o *ObjectType) UnsetDescription() raml.ObjectType {
	o.description = nil
	return o
}

func (o *ObjectType) SetAnnotations(annotations raml.AnnotationMap) raml.ObjectType {
	if annotations == nil {
		return o.UnsetAnnotations()
	}
	o.hasAnnotations.mp = annotations
	return o
}

func (o *ObjectType) UnsetAnnotations() raml.ObjectType {
	o.hasAnnotations.mp = NewAnnotationMap()
	return o
}

func (o *ObjectType) SetFacetDefinitions(facets raml.FacetMap) raml.ObjectType {
	if facets == nil {
		return o.UnsetFacetDefinitions()
	}
	o.facets = facets
	return o
}

func (o *ObjectType) UnsetFacetDefinitions() raml.ObjectType {
	o.facets = NewFacetMap()
	return o
}

func (o *ObjectType) SetXml(x raml.Xml) raml.ObjectType {
	o.xml = x
	return o
}

func (o *ObjectType) UnsetXml() raml.ObjectType {
	o.xml = nil
	return o
}

func (o *ObjectType) Enum() []interface{} {
	return o.enum
}

func (o *ObjectType) SetEnum(i []interface{}) raml.ObjectType {
	o.enum = i
	return o
}

func (o *ObjectType) UnsetEnum() raml.ObjectType {
	o.enum = nil
	return o
}

func (o *ObjectType) SetExtraFacets(facets raml.AnyMap) raml.ObjectType {
	if facets == nil {
		return o.UnsetExtraFacets()
	}
	o.hasExtra.mp = facets
	return o
}

func (o *ObjectType) UnsetExtraFacets() raml.ObjectType {
	o.hasExtra.mp = NewAnyMap()
	return o
}

func (o *ObjectType) SetRequired(b bool) raml.ObjectType {
	o.required = b
	return o
}

func (o *ObjectType) Properties() raml.PropertyMap {
	return o.properties
}

func (o *ObjectType) SetProperties(props raml.PropertyMap) raml.ObjectType {
	if props == nil {
		return o.UnsetProperties()
	}

	o.properties = props
	return o
}

func (o *ObjectType) UnsetProperties() raml.ObjectType {
	o.properties = NewPropertyMap()
	return o
}

func (o *ObjectType) MinProperties() option.Uint {
	return option.NewMaybeUint(o.minProps)
}

func (o *ObjectType) SetMinProperties(min uint) raml.ObjectType {
	o.minProps = &min
	return o
}

func (o *ObjectType) UnsetMinProperties() raml.ObjectType {
	o.minProps = nil
	return o
}

func (o *ObjectType) MaxProperties() option.Uint {
	return option.NewMaybeUint(o.maxProps)
}

func (o *ObjectType) SetMaxProperties(u uint) raml.ObjectType {
	o.maxProps = &u
	return o
}

func (o *ObjectType) UnsetMaxProperties() raml.ObjectType {
	o.maxProps = nil
	return o
}

func (o *ObjectType) AdditionalProperties() bool {
	return o.addtlProps
}

func (o *ObjectType) SetAdditionalProperties(val bool) raml.ObjectType {
	o.addtlProps = val
	return o
}

func (o *ObjectType) Discriminator() option.String {
	return option.NewMaybeString(o.discrim)
}

func (o *ObjectType) SetDiscriminator(facet string) raml.ObjectType {
	o.discrim = &facet
	return o
}

func (o *ObjectType) UnsetDiscriminator() raml.ObjectType {
	o.discrim = nil
	return o
}

func (o *ObjectType) DiscriminatorValue() option.Untyped {
	return option.NewMaybeUntyped(o.discrimVal)
}

func (o *ObjectType) SetDiscriminatorValue(val interface{}) raml.ObjectType {
	o.discrimVal = &val
	return o
}

func (o *ObjectType) UnsetDiscriminatorValue() raml.ObjectType {
	o.discrimVal = nil
	return o
}

func (o ObjectType) render() bool {
	return true
}
func (o *ObjectType) marshal(out raml.AnyMap) error {
	logrus.Trace("internal.ObjectType.marshal")
	out.PutNonNil(rmeta.KeyDefault, o.def)

	if err := o.ExtendedDataType.marshal(out); err != nil {
		return err
	}
	if !o.addtlProps {
		out.Put(rmeta.KeyAddtlProps, o.addtlProps)
	}

	out.PutNonNil(rmeta.KeyMinProperties, o.minProps).
		PutNonNil(rmeta.KeyMaxProperties, o.maxProps).
		PutNonNil(rmeta.KeyDiscriminator, o.discrim).
		PutNonNil(rmeta.KeyDiscriminatorVal, o.discrimVal)

	if o.properties.Len() > 0 {
		out.Put(rmeta.KeyProperties, o.properties)
	}
	out.PutNonNil(rmeta.KeyEnum, o.enum).
		PutNonNil(rmeta.KeyExample, o.example)

	if o.examples.Len() > 0 {
		out.PutNonNil(rmeta.KeyExamples, o.examples)
	}

	return nil
}

func (o *ObjectType) assign(key, val *yaml.Node) error {
	logrus.Trace("internal.ObjectType.assign")
	switch key.Value {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind); err != nil {
			return err
		} else if err := ex.UnmarshalRAML(val); err != nil {
			return err
		} else {
			o.example = ex.(raml.ObjectExample)
		}
		return nil
	case rmeta.KeyExamples:
		return o.examples.UnmarshalRAML(val)
	case rmeta.KeyEnum:
		return xyml.ForEachList(val, func(cur *yaml.Node) error {
			o.enum = append(o.enum, val)

			return nil
		})
		return nil
	case rmeta.KeyRequired:
		return assign.AsBool(val, &o.required)
	}

	switch key.Value {
	case rmeta.KeyProperties:
		return o.properties.UnmarshalRAML(val)
	case rmeta.KeyMinProperties:
		return assign.AsUintPtr(val, &o.minProps)
	case rmeta.KeyMaxProperties:
		return assign.AsUintPtr(val, &o.maxProps)
	case rmeta.KeyAddtlProps:
		return assign.AsBool(val, &o.addtlProps)
	case rmeta.KeyDiscriminator:
		return assign.AsStringPtr(val, &o.discrim)
	case rmeta.KeyDiscriminatorVal:
		var foo interface{} = val
		o.discrimVal = &foo
		return nil
	}

	return o.ExtendedDataType.assign(key, val)
}
