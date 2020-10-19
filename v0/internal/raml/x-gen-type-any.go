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

// NewAnyType returns a new internal implementation
// of the raml.AnyType interface.
//
// Generated @ 2020-10-19T13:48:24.9771134-04:00
func NewAnyType() *AnyType {
	out := &AnyType{
		examples: raml.NewAnyExampleMap(0),
	}

	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeAny, out)

	return out
}

// AnyType is a default generated implementation of
// the raml.AnyType interface
//
// Generated @ 2020-10-19T13:48:24.9771134-04:00
type AnyType struct {
	*ExtendedDataType

	def      *interface{}
	example  raml.AnyExample
	examples raml.AnyExampleMap
	enum     []interface{}
}

func (o *AnyType) SetType(s string) raml.AnyType {
	o.schema = s
	return o
}

func (o *AnyType) Default() option.Untyped {
	return option.NewMaybeUntyped(o.def)
}

func (o *AnyType) SetDefault(i interface{}) raml.AnyType {
	o.def = &i
	return o
}

func (o *AnyType) UnsetDefault() raml.AnyType {
	o.def = nil
	return o
}

func (o *AnyType) Example() raml.AnyExample {
	return o.example
}

func (o *AnyType) SetExample(ex raml.AnyExample) raml.AnyType {
	o.example = ex
	return o
}

func (o *AnyType) UnsetExample() raml.AnyType {
	o.example = nil
	return o
}

func (o *AnyType) Examples() raml.AnyExampleMap {
	return o.examples
}

func (o *AnyType) SetExamples(examples raml.AnyExampleMap) raml.AnyType {
	if examples == nil {
		return o.UnsetExamples()
	}

	o.examples = examples
	return o
}

func (o *AnyType) UnsetExamples() raml.AnyType {
	o.examples = raml.NewAnyExampleMap(0)
	return o
}

func (o *AnyType) SetDisplayName(s string) raml.AnyType {
	o.displayName = &s
	return o
}

func (o *AnyType) UnsetDisplayName() raml.AnyType {
	o.displayName = nil
	return o
}

func (o *AnyType) SetDescription(s string) raml.AnyType {
	o.description = &s
	return o
}

func (o *AnyType) UnsetDescription() raml.AnyType {
	o.description = nil
	return o
}

func (o *AnyType) SetAnnotations(annotations raml.AnnotationMap) raml.AnyType {
	if annotations == nil {
		return o.UnsetAnnotations()
	}

	o.hasAnnotations.mp = annotations
	return o
}

func (o *AnyType) UnsetAnnotations() raml.AnyType {
	o.hasAnnotations.mp = raml.NewAnnotationMap(0)
	return o
}

func (o *AnyType) SetFacetDefinitions(facets raml.FacetMap) raml.AnyType {
	if facets == nil {
		return o.UnsetFacetDefinitions()
	}

	o.facets = facets
	return o
}

func (o *AnyType) UnsetFacetDefinitions() raml.AnyType {
	o.facets = raml.NewFacetMap(0)
	return o
}

func (o *AnyType) SetXML(x raml.XML) raml.AnyType {
	o.xml = x
	return o
}

func (o *AnyType) UnsetXML() raml.AnyType {
	o.xml = nil
	return o
}

func (o *AnyType) Enum() []interface{} {
	return o.enum
}

func (o *AnyType) SetEnum(i []interface{}) raml.AnyType {
	o.enum = i
	return o
}

func (o *AnyType) UnsetEnum() raml.AnyType {
	o.enum = nil
	return o
}

func (o *AnyType) SetExtraFacets(facets raml.AnyMap) raml.AnyType {
	if facets == nil {
		return o.UnsetExtraFacets()
	}

	o.hasExtra.mp = facets
	return o
}

func (o *AnyType) UnsetExtraFacets() raml.AnyType {
	o.hasExtra.mp = raml.NewAnyMap(0)
	return o
}

func (o *AnyType) SetRequired(b bool) raml.AnyType {
	o.required = b
	return o
}

func (o *AnyType) marshal(out raml.AnyMap) error {
	logrus.Trace("internal.AnyType.marshal")
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

func (o *AnyType) assign(key, val *yaml.Node) error {
	switch key.Value {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind); err != nil {
			return err
		} else if err := ex.UnmarshalRAML(val); err != nil {
			return err
		} else {
			o.example = ex.(raml.AnyExample)
		}

		return nil
	case rmeta.KeyExamples:
		return UnmarshalAnyExampleMapRAML(o.examples, val)
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
