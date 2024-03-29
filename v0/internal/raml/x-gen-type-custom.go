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

// NewCustomType returns a new internal implementation
// of the raml.CustomType interface.
//
// Generated @ 2023-01-17T10:02:54.294844187-05:00
func NewCustomType() *CustomType {
	out := &CustomType{
		examples: raml.NewCustomExampleMap(0),
	}

	out.examples.SerializeOrdered(false)

	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeCustom, out)

	return out
}

// CustomType is a default generated implementation of
// the raml.CustomType interface
//
// Generated @ 2023-01-17T10:02:54.294844187-05:00
type CustomType struct {
	*ExtendedDataType

	def      *interface{}
	example  raml.CustomExample
	examples raml.CustomExampleMap
	enum     []interface{}
}

func (o *CustomType) SetType(s string) raml.CustomType {
	o.schema = s
	return o
}

func (o *CustomType) Default() option.Untyped {
	return option.NewMaybeUntyped(o.def)
}

func (o *CustomType) SetDefault(i interface{}) raml.CustomType {
	o.def = &i
	return o
}

func (o *CustomType) UnsetDefault() raml.CustomType {
	o.def = nil
	return o
}

func (o *CustomType) Example() raml.CustomExample {
	return o.example
}

func (o *CustomType) SetExample(ex raml.CustomExample) raml.CustomType {
	o.example = ex
	return o
}

func (o *CustomType) UnsetExample() raml.CustomType {
	o.example = nil
	return o
}

func (o *CustomType) Examples() raml.CustomExampleMap {
	return o.examples
}

func (o *CustomType) SetExamples(examples raml.CustomExampleMap) raml.CustomType {
	if examples == nil {
		return o.UnsetExamples()
	}

	o.examples = examples
	return o
}

func (o *CustomType) UnsetExamples() raml.CustomType {
	o.examples = raml.NewCustomExampleMap(0)
	return o
}

func (o *CustomType) SetDisplayName(s string) raml.CustomType {
	o.displayName = &s
	return o
}

func (o *CustomType) UnsetDisplayName() raml.CustomType {
	o.displayName = nil
	return o
}

func (o *CustomType) SetDescription(s string) raml.CustomType {
	o.description = &s
	return o
}

func (o *CustomType) UnsetDescription() raml.CustomType {
	o.description = nil
	return o
}

func (o *CustomType) SetAnnotations(annotations raml.AnnotationMap) raml.CustomType {
	if annotations == nil {
		return o.UnsetAnnotations()
	}

	o.hasAnnotations.mp = annotations
	return o
}

func (o *CustomType) UnsetAnnotations() raml.CustomType {
	o.hasAnnotations.mp = raml.NewAnnotationMap(0)
	return o
}

func (o *CustomType) SetFacetDefinitions(facets raml.FacetMap) raml.CustomType {
	if facets == nil {
		return o.UnsetFacetDefinitions()
	}

	o.facets = facets
	return o
}

func (o *CustomType) UnsetFacetDefinitions() raml.CustomType {
	o.facets = raml.NewFacetMap(0)
	return o
}

func (o *CustomType) SetXML(x raml.XML) raml.CustomType {
	o.xml = x
	return o
}

func (o *CustomType) UnsetXML() raml.CustomType {
	o.xml = nil
	return o
}

func (o *CustomType) Enum() []interface{} {
	return o.enum
}

func (o *CustomType) SetEnum(i []interface{}) raml.CustomType {
	o.enum = i
	return o
}

func (o *CustomType) UnsetEnum() raml.CustomType {
	o.enum = nil
	return o
}

func (o *CustomType) SetExtraFacets(facets raml.AnyMap) raml.CustomType {
	if facets == nil {
		return o.UnsetExtraFacets()
	}

	o.hasExtra.mp = facets
	return o
}

func (o *CustomType) UnsetExtraFacets() raml.CustomType {
	o.hasExtra.mp = raml.NewAnyMap(0)
	return o
}

func (o *CustomType) SetRequired(b bool) raml.CustomType {
	o.required = b
	return o
}

func (o *CustomType) marshal(out raml.AnyMap) error {
	logrus.Trace("internal.CustomType.marshal")
	out.PutIfNotNil(rmeta.KeyDefault, o.def)

	if err := o.ExtendedDataType.marshal(out); err != nil {
		return err
	}

	out.PutIfNotNil(rmeta.KeyEnum, o.enum).
		PutIfNotNil(rmeta.KeyExample, o.example)

	if o.examples.Len() > 0 {
		out.PutIfNotNil(rmeta.KeyExamples, o.examples)
	}

	return nil
}

func (o *CustomType) assign(key, val *yaml.Node) error {
	switch key.Value {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind); err != nil {
			return err
		} else if err := ex.UnmarshalRAML(val); err != nil {
			return err
		} else {
			o.example = ex.(raml.CustomExample)
		}

		return nil
	case rmeta.KeyExamples:
		return UnmarshalCustomExampleMapRAML(o.examples, val)
	case rmeta.KeyEnum:
		return xyml.SequenceForEach(val, func(cur *yaml.Node) error {
			o.enum = append(o.enum, cur)

			return nil
		})
		return nil
	case rmeta.KeyRequired:
		return assign.AsBool(val, &o.required)
	}

	return o.ExtendedDataType.assign(key, val)
}
