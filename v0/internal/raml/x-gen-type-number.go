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

// NewNumberType returns a new internal implementation
// of the raml.NumberType interface.
//
// Generated @ 2020-05-20T20:54:26.833516016-04:00
func NewNumberType() *NumberType {
	out := &NumberType{
		examples: NewNumberExampleMap(),
	}

	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeNumber, out)

	return out
}

// NumberType is a default generated implementation of
// the raml.NumberType interface
//
// Generated @ 2020-05-20T20:54:26.833516016-04:00
type NumberType struct {
	*ExtendedDataType

	def        *float64
	example    raml.NumberExample
	examples   raml.NumberExampleMap
	enum       []float64
	minimum    *float64
	maximum    *float64
	format     raml.NumberFormat
	multipleOf *float64
}

func (o *NumberType) SetType(s string) raml.NumberType {
	o.schema = s
	return o
}

func (o *NumberType) Default() option.Float64 {
	return option.NewMaybeFloat64(o.def)
}

func (o *NumberType) SetDefault(i float64) raml.NumberType {
	o.def = &i
	return o
}

func (o *NumberType) UnsetDefault() raml.NumberType {
	o.def = nil
	return o
}

func (o *NumberType) Example() raml.NumberExample {
	return o.example
}

func (o *NumberType) SetExample(ex raml.NumberExample) raml.NumberType {
	o.example = ex
	return o
}

func (o *NumberType) UnsetExample() raml.NumberType {
	o.example = nil
	return o
}

func (o *NumberType) Examples() raml.NumberExampleMap {
	return o.examples
}

func (o *NumberType) SetExamples(examples raml.NumberExampleMap) raml.NumberType {
	if examples == nil {
		return o.UnsetExamples()
	}

	o.examples = examples
	return o
}

func (o *NumberType) UnsetExamples() raml.NumberType {
	o.examples = NewNumberExampleMap()
	return o
}

func (o *NumberType) SetDisplayName(s string) raml.NumberType {
	o.displayName = &s
	return o
}

func (o *NumberType) UnsetDisplayName() raml.NumberType {
	o.displayName = nil
	return o
}

func (o *NumberType) SetDescription(s string) raml.NumberType {
	o.description = &s
	return o
}

func (o *NumberType) UnsetDescription() raml.NumberType {
	o.description = nil
	return o
}

func (o *NumberType) SetAnnotations(annotations raml.AnnotationMap) raml.NumberType {
	if annotations == nil {
		return o.UnsetAnnotations()
	}

	o.hasAnnotations.mp = annotations
	return o
}

func (o *NumberType) UnsetAnnotations() raml.NumberType {
	o.hasAnnotations.mp = NewAnnotationMap()
	return o
}

func (o *NumberType) SetFacetDefinitions(facets raml.FacetMap) raml.NumberType {
	if facets == nil {
		return o.UnsetFacetDefinitions()
	}

	o.facets = facets
	return o
}

func (o *NumberType) UnsetFacetDefinitions() raml.NumberType {
	o.facets = NewFacetMap()
	return o
}

func (o *NumberType) SetXML(x raml.XML) raml.NumberType {
	o.xml = x
	return o
}

func (o *NumberType) UnsetXML() raml.NumberType {
	o.xml = nil
	return o
}

func (o *NumberType) Enum() []float64 {
	return o.enum
}

func (o *NumberType) SetEnum(i []float64) raml.NumberType {
	o.enum = i
	return o
}

func (o *NumberType) UnsetEnum() raml.NumberType {
	o.enum = nil
	return o
}

func (o *NumberType) SetExtraFacets(facets raml.AnyMap) raml.NumberType {
	if facets == nil {
		return o.UnsetExtraFacets()
	}

	o.hasExtra.mp = facets
	return o
}

func (o *NumberType) UnsetExtraFacets() raml.NumberType {
	o.hasExtra.mp = NewAnyMap()
	return o
}

func (o *NumberType) SetRequired(b bool) raml.NumberType {
	o.required = b
	return o
}

func (o *NumberType) Minimum() option.Float64 {
	return option.NewMaybeFloat64(o.minimum)
}

func (o *NumberType) SetMinimum(pat float64) raml.NumberType {
	o.minimum = &pat
	return o
}

func (o *NumberType) UnsetMinimum() raml.NumberType {
	o.minimum = nil
	return o
}

func (o *NumberType) Maximum() option.Float64 {
	return option.NewMaybeFloat64(o.maximum)
}

func (o *NumberType) SetMaximum(pat float64) raml.NumberType {
	o.maximum = &pat
	return o
}

func (o *NumberType) UnsetMaximum() raml.NumberType {
	o.maximum = nil
	return o
}

func (o *NumberType) Format() raml.NumberFormat {
	return o.format
}

func (o *NumberType) SetFormat(f raml.NumberFormat) raml.NumberType {
	o.format = f
	return o
}

func (o *NumberType) UnsetFormat() raml.NumberType {
	o.format = nil
	return o
}

func (o *NumberType) MultipleOf() option.Float64 {
	return option.NewMaybeFloat64(o.multipleOf)
}

func (o *NumberType) SetMultipleOf(pat float64) raml.NumberType {
	o.multipleOf = &pat
	return o
}

func (o *NumberType) UnsetMultipleOf() raml.NumberType {
	o.multipleOf = nil
	return o
}

func (o *NumberType) marshal(out raml.AnyMap) error {
	logrus.Trace("internal.NumberType.marshal")
	out.PutNonNil(rmeta.KeyDefault, o.def)

	if err := o.ExtendedDataType.marshal(out); err != nil {
		return err
	}
	out.PutNonNil(rmeta.KeyFormat, o.format).
		PutNonNil(rmeta.KeyMinimum, o.minimum).
		PutNonNil(rmeta.KeyMaximum, o.maximum).
		PutNonNil(rmeta.KeyMultipleOf, o.multipleOf)
	out.PutNonNil(rmeta.KeyEnum, o.enum).
		PutNonNil(rmeta.KeyExample, o.example)

	if o.examples.Len() > 0 {
		out.PutNonNil(rmeta.KeyExamples, o.examples)
	}

	return nil
}

func (o *NumberType) assign(key, val *yaml.Node) error {
	switch key.Value {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind); err != nil {
			return err
		} else if err := ex.UnmarshalRAML(val); err != nil {
			return err
		} else {
			o.example = ex.(raml.NumberExample)
		}

		return nil
	case rmeta.KeyExamples:
		return o.examples.UnmarshalRAML(val)
	case rmeta.KeyEnum:
		return xyml.ForEachList(val, func(cur *yaml.Node) error {
			if val, err := xyml.ToFloat64(cur); err != nil {
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

	switch key.Value {
	case rmeta.KeyMinimum:
		return assign.AsFloat64Ptr(val, &o.minimum)
	case rmeta.KeyMaximum:
		return assign.AsFloat64Ptr(val, &o.maximum)
	case rmeta.KeyFormat:
		if val, err := NumberFormatSortingHat(val); err != nil {
			return err
		} else {
			o.format = val
			return nil
		}
	case rmeta.KeyMultipleOf:
		return assign.AsFloat64Ptr(val, &o.multipleOf)
	}

	return o.ExtendedDataType.assign(key, val)
}
