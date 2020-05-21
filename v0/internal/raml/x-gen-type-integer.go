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

// NewIntegerType returns a new internal implementation
// of the raml.IntegerType interface.
//
// Generated @ 2020-05-20T21:46:01.015916886-04:00
func NewIntegerType() *IntegerType {
	out := &IntegerType{
		examples: NewIntegerExampleMap(),
	}

	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeInteger, out)

	return out
}

// IntegerType is a default generated implementation of
// the raml.IntegerType interface
//
// Generated @ 2020-05-20T21:46:01.015916886-04:00
type IntegerType struct {
	*ExtendedDataType

	def        *int64
	example    raml.IntegerExample
	examples   raml.IntegerExampleMap
	enum       []int64
	minimum    *int64
	maximum    *int64
	format     raml.NumberFormat
	multipleOf *float64
}

func (o *IntegerType) SetType(s string) raml.IntegerType {
	o.schema = s
	return o
}

func (o *IntegerType) Default() option.Int64 {
	return option.NewMaybeInt64(o.def)
}

func (o *IntegerType) SetDefault(i int64) raml.IntegerType {
	o.def = &i
	return o
}

func (o *IntegerType) UnsetDefault() raml.IntegerType {
	o.def = nil
	return o
}

func (o *IntegerType) Example() raml.IntegerExample {
	return o.example
}

func (o *IntegerType) SetExample(ex raml.IntegerExample) raml.IntegerType {
	o.example = ex
	return o
}

func (o *IntegerType) UnsetExample() raml.IntegerType {
	o.example = nil
	return o
}

func (o *IntegerType) Examples() raml.IntegerExampleMap {
	return o.examples
}

func (o *IntegerType) SetExamples(examples raml.IntegerExampleMap) raml.IntegerType {
	if examples == nil {
		return o.UnsetExamples()
	}

	o.examples = examples
	return o
}

func (o *IntegerType) UnsetExamples() raml.IntegerType {
	o.examples = NewIntegerExampleMap()
	return o
}

func (o *IntegerType) SetDisplayName(s string) raml.IntegerType {
	o.displayName = &s
	return o
}

func (o *IntegerType) UnsetDisplayName() raml.IntegerType {
	o.displayName = nil
	return o
}

func (o *IntegerType) SetDescription(s string) raml.IntegerType {
	o.description = &s
	return o
}

func (o *IntegerType) UnsetDescription() raml.IntegerType {
	o.description = nil
	return o
}

func (o *IntegerType) SetAnnotations(annotations raml.AnnotationMap) raml.IntegerType {
	if annotations == nil {
		return o.UnsetAnnotations()
	}

	o.hasAnnotations.mp = annotations
	return o
}

func (o *IntegerType) UnsetAnnotations() raml.IntegerType {
	o.hasAnnotations.mp = NewAnnotationMap()
	return o
}

func (o *IntegerType) SetFacetDefinitions(facets raml.FacetMap) raml.IntegerType {
	if facets == nil {
		return o.UnsetFacetDefinitions()
	}

	o.facets = facets
	return o
}

func (o *IntegerType) UnsetFacetDefinitions() raml.IntegerType {
	o.facets = NewFacetMap()
	return o
}

func (o *IntegerType) SetXML(x raml.XML) raml.IntegerType {
	o.xml = x
	return o
}

func (o *IntegerType) UnsetXML() raml.IntegerType {
	o.xml = nil
	return o
}

func (o *IntegerType) Enum() []int64 {
	return o.enum
}

func (o *IntegerType) SetEnum(i []int64) raml.IntegerType {
	o.enum = i
	return o
}

func (o *IntegerType) UnsetEnum() raml.IntegerType {
	o.enum = nil
	return o
}

func (o *IntegerType) SetExtraFacets(facets raml.AnyMap) raml.IntegerType {
	if facets == nil {
		return o.UnsetExtraFacets()
	}

	o.hasExtra.mp = facets
	return o
}

func (o *IntegerType) UnsetExtraFacets() raml.IntegerType {
	o.hasExtra.mp = NewAnyMap()
	return o
}

func (o *IntegerType) SetRequired(b bool) raml.IntegerType {
	o.required = b
	return o
}

func (o *IntegerType) Minimum() option.Int64 {
	return option.NewMaybeInt64(o.minimum)
}

func (o *IntegerType) SetMinimum(pat int64) raml.IntegerType {
	o.minimum = &pat
	return o
}

func (o *IntegerType) UnsetMinimum() raml.IntegerType {
	o.minimum = nil
	return o
}

func (o *IntegerType) Maximum() option.Int64 {
	return option.NewMaybeInt64(o.maximum)
}

func (o *IntegerType) SetMaximum(pat int64) raml.IntegerType {
	o.maximum = &pat
	return o
}

func (o *IntegerType) UnsetMaximum() raml.IntegerType {
	o.maximum = nil
	return o
}

func (o *IntegerType) Format() raml.IntegerFormat {
	return o.format
}

func (o *IntegerType) SetFormat(f raml.IntegerFormat) raml.IntegerType {
	o.format = f
	return o
}

func (o *IntegerType) UnsetFormat() raml.IntegerType {
	o.format = nil
	return o
}

func (o *IntegerType) MultipleOf() option.Float64 {
	return option.NewMaybeFloat64(o.multipleOf)
}

func (o *IntegerType) SetMultipleOf(pat float64) raml.IntegerType {
	o.multipleOf = &pat
	return o
}

func (o *IntegerType) UnsetMultipleOf() raml.IntegerType {
	o.multipleOf = nil
	return o
}

func (o *IntegerType) marshal(out raml.AnyMap) error {
	logrus.Trace("internal.IntegerType.marshal")
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

func (o *IntegerType) assign(key, val *yaml.Node) error {
	switch key.Value {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind); err != nil {
			return err
		} else if err := ex.UnmarshalRAML(val); err != nil {
			return err
		} else {
			o.example = ex.(raml.IntegerExample)
		}

		return nil
	case rmeta.KeyExamples:
		return o.examples.UnmarshalRAML(val)
	case rmeta.KeyEnum:
		return xyml.ForEachList(val, func(cur *yaml.Node) error {
			if val, err := xyml.ToInt64(cur); err != nil {
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
		return assign.AsInt64Ptr(val, &o.minimum)
	case rmeta.KeyMaximum:
		return assign.AsInt64Ptr(val, &o.maximum)
	case rmeta.KeyFormat:
		if val, err := IntegerFormatSortingHat(val); err != nil {
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
