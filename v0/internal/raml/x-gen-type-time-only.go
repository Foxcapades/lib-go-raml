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

// NewTimeOnlyType returns a new internal implementation
// of the raml.TimeOnlyType interface.
//
// Generated @ 2020-05-20T21:46:01.015916886-04:00
func NewTimeOnlyType() *TimeOnlyType {
	out := &TimeOnlyType{
		examples: NewTimeOnlyExampleMap(),
	}

	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeTimeOnly, out)

	return out
}

// TimeOnlyType is a default generated implementation of
// the raml.TimeOnlyType interface
//
// Generated @ 2020-05-20T21:46:01.015916886-04:00
type TimeOnlyType struct {
	*ExtendedDataType

	def      *string
	example  raml.TimeOnlyExample
	examples raml.TimeOnlyExampleMap
	enum     []string
}

func (o *TimeOnlyType) SetType(s string) raml.TimeOnlyType {
	o.schema = s
	return o
}

func (o *TimeOnlyType) Default() option.String {
	return option.NewMaybeString(o.def)
}

func (o *TimeOnlyType) SetDefault(i string) raml.TimeOnlyType {
	o.def = &i
	return o
}

func (o *TimeOnlyType) UnsetDefault() raml.TimeOnlyType {
	o.def = nil
	return o
}

func (o *TimeOnlyType) Example() raml.TimeOnlyExample {
	return o.example
}

func (o *TimeOnlyType) SetExample(ex raml.TimeOnlyExample) raml.TimeOnlyType {
	o.example = ex
	return o
}

func (o *TimeOnlyType) UnsetExample() raml.TimeOnlyType {
	o.example = nil
	return o
}

func (o *TimeOnlyType) Examples() raml.TimeOnlyExampleMap {
	return o.examples
}

func (o *TimeOnlyType) SetExamples(examples raml.TimeOnlyExampleMap) raml.TimeOnlyType {
	if examples == nil {
		return o.UnsetExamples()
	}

	o.examples = examples
	return o
}

func (o *TimeOnlyType) UnsetExamples() raml.TimeOnlyType {
	o.examples = NewTimeOnlyExampleMap()
	return o
}

func (o *TimeOnlyType) SetDisplayName(s string) raml.TimeOnlyType {
	o.displayName = &s
	return o
}

func (o *TimeOnlyType) UnsetDisplayName() raml.TimeOnlyType {
	o.displayName = nil
	return o
}

func (o *TimeOnlyType) SetDescription(s string) raml.TimeOnlyType {
	o.description = &s
	return o
}

func (o *TimeOnlyType) UnsetDescription() raml.TimeOnlyType {
	o.description = nil
	return o
}

func (o *TimeOnlyType) SetAnnotations(annotations raml.AnnotationMap) raml.TimeOnlyType {
	if annotations == nil {
		return o.UnsetAnnotations()
	}

	o.hasAnnotations.mp = annotations
	return o
}

func (o *TimeOnlyType) UnsetAnnotations() raml.TimeOnlyType {
	o.hasAnnotations.mp = NewAnnotationMap()
	return o
}

func (o *TimeOnlyType) SetFacetDefinitions(facets raml.FacetMap) raml.TimeOnlyType {
	if facets == nil {
		return o.UnsetFacetDefinitions()
	}

	o.facets = facets
	return o
}

func (o *TimeOnlyType) UnsetFacetDefinitions() raml.TimeOnlyType {
	o.facets = NewFacetMap()
	return o
}

func (o *TimeOnlyType) SetXML(x raml.XML) raml.TimeOnlyType {
	o.xml = x
	return o
}

func (o *TimeOnlyType) UnsetXML() raml.TimeOnlyType {
	o.xml = nil
	return o
}

func (o *TimeOnlyType) Enum() []string {
	return o.enum
}

func (o *TimeOnlyType) SetEnum(i []string) raml.TimeOnlyType {
	o.enum = i
	return o
}

func (o *TimeOnlyType) UnsetEnum() raml.TimeOnlyType {
	o.enum = nil
	return o
}

func (o *TimeOnlyType) SetExtraFacets(facets raml.AnyMap) raml.TimeOnlyType {
	if facets == nil {
		return o.UnsetExtraFacets()
	}

	o.hasExtra.mp = facets
	return o
}

func (o *TimeOnlyType) UnsetExtraFacets() raml.TimeOnlyType {
	o.hasExtra.mp = NewAnyMap()
	return o
}

func (o *TimeOnlyType) SetRequired(b bool) raml.TimeOnlyType {
	o.required = b
	return o
}

func (o *TimeOnlyType) marshal(out raml.AnyMap) error {
	logrus.Trace("internal.TimeOnlyType.marshal")
	out.PutNonNil(rmeta.KeyDefault, o.def)

	if err := o.ExtendedDataType.marshal(out); err != nil {
		return err
	}

	out.PutNonNil(rmeta.KeyEnum, o.enum).
		PutNonNil(rmeta.KeyExample, o.example)

	if o.examples.Len() > 0 {
		out.PutNonNil(rmeta.KeyExamples, o.examples)
	}

	return nil
}

func (o *TimeOnlyType) assign(key, val *yaml.Node) error {
	switch key.Value {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind); err != nil {
			return err
		} else if err := ex.UnmarshalRAML(val); err != nil {
			return err
		} else {
			o.example = ex.(raml.TimeOnlyExample)
		}

		return nil
	case rmeta.KeyExamples:
		return o.examples.UnmarshalRAML(val)
	case rmeta.KeyEnum:
		return xyml.ForEachList(val, func(cur *yaml.Node) error {
			if val, err := xyml.ToString(cur); err != nil {
				return err
			} else {
				o.enum = append(o.enum, val)
			}

			return nil
		})
		return nil
	case rmeta.KeyRequired:
		return assign.AsBool(val, &o.required)
	}

	return o.ExtendedDataType.assign(key, val)
}
