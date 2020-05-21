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

// NewDatetimeType returns a new internal implementation
// of the raml.DatetimeType interface.
//
// Generated @ 2020-05-20T21:46:01.015916886-04:00
func NewDatetimeType() *DatetimeType {
	out := &DatetimeType{
		examples: NewDatetimeExampleMap(),
	}

	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeDatetime, out)

	return out
}

// DatetimeType is a default generated implementation of
// the raml.DatetimeType interface
//
// Generated @ 2020-05-20T21:46:01.015916886-04:00
type DatetimeType struct {
	*ExtendedDataType

	def      *string
	example  raml.DatetimeExample
	examples raml.DatetimeExampleMap
	enum     []string
	format   raml.DateFormat
}

func (o *DatetimeType) SetType(s string) raml.DatetimeType {
	o.schema = s
	return o
}

func (o *DatetimeType) Default() option.String {
	return option.NewMaybeString(o.def)
}

func (o *DatetimeType) SetDefault(i string) raml.DatetimeType {
	o.def = &i
	return o
}

func (o *DatetimeType) UnsetDefault() raml.DatetimeType {
	o.def = nil
	return o
}

func (o *DatetimeType) Example() raml.DatetimeExample {
	return o.example
}

func (o *DatetimeType) SetExample(ex raml.DatetimeExample) raml.DatetimeType {
	o.example = ex
	return o
}

func (o *DatetimeType) UnsetExample() raml.DatetimeType {
	o.example = nil
	return o
}

func (o *DatetimeType) Examples() raml.DatetimeExampleMap {
	return o.examples
}

func (o *DatetimeType) SetExamples(examples raml.DatetimeExampleMap) raml.DatetimeType {
	if examples == nil {
		return o.UnsetExamples()
	}

	o.examples = examples
	return o
}

func (o *DatetimeType) UnsetExamples() raml.DatetimeType {
	o.examples = NewDatetimeExampleMap()
	return o
}

func (o *DatetimeType) SetDisplayName(s string) raml.DatetimeType {
	o.displayName = &s
	return o
}

func (o *DatetimeType) UnsetDisplayName() raml.DatetimeType {
	o.displayName = nil
	return o
}

func (o *DatetimeType) SetDescription(s string) raml.DatetimeType {
	o.description = &s
	return o
}

func (o *DatetimeType) UnsetDescription() raml.DatetimeType {
	o.description = nil
	return o
}

func (o *DatetimeType) SetAnnotations(annotations raml.AnnotationMap) raml.DatetimeType {
	if annotations == nil {
		return o.UnsetAnnotations()
	}

	o.hasAnnotations.mp = annotations
	return o
}

func (o *DatetimeType) UnsetAnnotations() raml.DatetimeType {
	o.hasAnnotations.mp = NewAnnotationMap()
	return o
}

func (o *DatetimeType) SetFacetDefinitions(facets raml.FacetMap) raml.DatetimeType {
	if facets == nil {
		return o.UnsetFacetDefinitions()
	}

	o.facets = facets
	return o
}

func (o *DatetimeType) UnsetFacetDefinitions() raml.DatetimeType {
	o.facets = NewFacetMap()
	return o
}

func (o *DatetimeType) SetXML(x raml.XML) raml.DatetimeType {
	o.xml = x
	return o
}

func (o *DatetimeType) UnsetXML() raml.DatetimeType {
	o.xml = nil
	return o
}

func (o *DatetimeType) Enum() []string {
	return o.enum
}

func (o *DatetimeType) SetEnum(i []string) raml.DatetimeType {
	o.enum = i
	return o
}

func (o *DatetimeType) UnsetEnum() raml.DatetimeType {
	o.enum = nil
	return o
}

func (o *DatetimeType) SetExtraFacets(facets raml.AnyMap) raml.DatetimeType {
	if facets == nil {
		return o.UnsetExtraFacets()
	}

	o.hasExtra.mp = facets
	return o
}

func (o *DatetimeType) UnsetExtraFacets() raml.DatetimeType {
	o.hasExtra.mp = NewAnyMap()
	return o
}

func (o *DatetimeType) SetRequired(b bool) raml.DatetimeType {
	o.required = b
	return o
}

func (o *DatetimeType) Format() raml.DateFormat {
	return o.format
}

func (o *DatetimeType) SetFormat(f raml.DateFormat) raml.DatetimeType {
	o.format = f
	return o
}

func (o *DatetimeType) UnsetFormat() raml.DatetimeType {
	o.format = nil
	return o
}

func (o *DatetimeType) marshal(out raml.AnyMap) error {
	logrus.Trace("internal.DatetimeType.marshal")
	out.PutNonNil(rmeta.KeyDefault, o.def)

	if err := o.ExtendedDataType.marshal(out); err != nil {
		return err
	}
	out.PutNonNil(rmeta.KeyFormat, o.format)
	out.PutNonNil(rmeta.KeyEnum, o.enum).
		PutNonNil(rmeta.KeyExample, o.example)

	if o.examples.Len() > 0 {
		out.PutNonNil(rmeta.KeyExamples, o.examples)
	}

	return nil
}

func (o *DatetimeType) assign(key, val *yaml.Node) error {
	switch key.Value {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind); err != nil {
			return err
		} else if err := ex.UnmarshalRAML(val); err != nil {
			return err
		} else {
			o.example = ex.(raml.DatetimeExample)
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

	if key.Value == rmeta.KeyFormat {
		if val, err := DateFormatSortingHat(val); err != nil {
			return err
		} else {
			o.format = val
		}

		return nil
	}

	return o.ExtendedDataType.assign(key, val)
}
