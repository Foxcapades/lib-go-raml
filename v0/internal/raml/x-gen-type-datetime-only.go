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

// NewDatetimeOnlyType returns a new internal implementation
// of the raml.DatetimeOnlyType interface.
//
// Generated @ 2020-05-20T21:46:01.015916886-04:00
func NewDatetimeOnlyType() *DatetimeOnlyType {
	out := &DatetimeOnlyType{
		examples: NewDatetimeOnlyExampleMap(),
	}

	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeDatetimeOnly, out)

	return out
}

// DatetimeOnlyType is a default generated implementation of
// the raml.DatetimeOnlyType interface
//
// Generated @ 2020-05-20T21:46:01.015916886-04:00
type DatetimeOnlyType struct {
	*ExtendedDataType

	def      *string
	example  raml.DatetimeOnlyExample
	examples raml.DatetimeOnlyExampleMap
	enum     []string
}

func (o *DatetimeOnlyType) SetType(s string) raml.DatetimeOnlyType {
	o.schema = s
	return o
}

func (o *DatetimeOnlyType) Default() option.String {
	return option.NewMaybeString(o.def)
}

func (o *DatetimeOnlyType) SetDefault(i string) raml.DatetimeOnlyType {
	o.def = &i
	return o
}

func (o *DatetimeOnlyType) UnsetDefault() raml.DatetimeOnlyType {
	o.def = nil
	return o
}

func (o *DatetimeOnlyType) Example() raml.DatetimeOnlyExample {
	return o.example
}

func (o *DatetimeOnlyType) SetExample(ex raml.DatetimeOnlyExample) raml.DatetimeOnlyType {
	o.example = ex
	return o
}

func (o *DatetimeOnlyType) UnsetExample() raml.DatetimeOnlyType {
	o.example = nil
	return o
}

func (o *DatetimeOnlyType) Examples() raml.DatetimeOnlyExampleMap {
	return o.examples
}

func (o *DatetimeOnlyType) SetExamples(examples raml.DatetimeOnlyExampleMap) raml.DatetimeOnlyType {
	if examples == nil {
		return o.UnsetExamples()
	}

	o.examples = examples
	return o
}

func (o *DatetimeOnlyType) UnsetExamples() raml.DatetimeOnlyType {
	o.examples = NewDatetimeOnlyExampleMap()
	return o
}

func (o *DatetimeOnlyType) SetDisplayName(s string) raml.DatetimeOnlyType {
	o.displayName = &s
	return o
}

func (o *DatetimeOnlyType) UnsetDisplayName() raml.DatetimeOnlyType {
	o.displayName = nil
	return o
}

func (o *DatetimeOnlyType) SetDescription(s string) raml.DatetimeOnlyType {
	o.description = &s
	return o
}

func (o *DatetimeOnlyType) UnsetDescription() raml.DatetimeOnlyType {
	o.description = nil
	return o
}

func (o *DatetimeOnlyType) SetAnnotations(annotations raml.AnnotationMap) raml.DatetimeOnlyType {
	if annotations == nil {
		return o.UnsetAnnotations()
	}

	o.hasAnnotations.mp = annotations
	return o
}

func (o *DatetimeOnlyType) UnsetAnnotations() raml.DatetimeOnlyType {
	o.hasAnnotations.mp = NewAnnotationMap()
	return o
}

func (o *DatetimeOnlyType) SetFacetDefinitions(facets raml.FacetMap) raml.DatetimeOnlyType {
	if facets == nil {
		return o.UnsetFacetDefinitions()
	}

	o.facets = facets
	return o
}

func (o *DatetimeOnlyType) UnsetFacetDefinitions() raml.DatetimeOnlyType {
	o.facets = NewFacetMap()
	return o
}

func (o *DatetimeOnlyType) SetXML(x raml.XML) raml.DatetimeOnlyType {
	o.xml = x
	return o
}

func (o *DatetimeOnlyType) UnsetXML() raml.DatetimeOnlyType {
	o.xml = nil
	return o
}

func (o *DatetimeOnlyType) Enum() []string {
	return o.enum
}

func (o *DatetimeOnlyType) SetEnum(i []string) raml.DatetimeOnlyType {
	o.enum = i
	return o
}

func (o *DatetimeOnlyType) UnsetEnum() raml.DatetimeOnlyType {
	o.enum = nil
	return o
}

func (o *DatetimeOnlyType) SetExtraFacets(facets raml.AnyMap) raml.DatetimeOnlyType {
	if facets == nil {
		return o.UnsetExtraFacets()
	}

	o.hasExtra.mp = facets
	return o
}

func (o *DatetimeOnlyType) UnsetExtraFacets() raml.DatetimeOnlyType {
	o.hasExtra.mp = NewAnyMap()
	return o
}

func (o *DatetimeOnlyType) SetRequired(b bool) raml.DatetimeOnlyType {
	o.required = b
	return o
}

func (o *DatetimeOnlyType) marshal(out raml.AnyMap) error {
	logrus.Trace("internal.DatetimeOnlyType.marshal")
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

func (o *DatetimeOnlyType) assign(key, val *yaml.Node) error {
	switch key.Value {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind); err != nil {
			return err
		} else if err := ex.UnmarshalRAML(val); err != nil {
			return err
		} else {
			o.example = ex.(raml.DatetimeOnlyExample)
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
