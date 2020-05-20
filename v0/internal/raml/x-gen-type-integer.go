package raml

import (
	"reflect"

	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
)

// NewIntegerType returns a new internal implementation
// of the raml.IntegerType interface.
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
func NewIntegerType(log *logrus.Entry) *IntegerType {
	log = xlog.WithType(log, "internal.IntegerType")

	out := &IntegerType{
		examples: NewIntegerExampleMap(log),
	}
	
	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeInteger, log, out)

	return out
}

// IntegerType is a default generated implementation of
// the raml.IntegerType interface
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
type IntegerType struct {
	*ExtendedDataType

	def      *int64
	example  raml.IntegerExample
	examples raml.IntegerExampleMap
	enum     []int64
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
	o.examples = NewIntegerExampleMap(o.DataType.log)
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
	o.hasAnnotations.mp = NewAnnotationMap(o.DataType.log)
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
	o.facets = NewFacetMap(o.DataType.log)
	return o
}

func (o *IntegerType) SetXml(x raml.Xml) raml.IntegerType {
	o.xml = x
	return o
}

func (o *IntegerType) UnsetXml() raml.IntegerType {
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
	o.hasExtra.mp = NewAnyMap(o.DataType.log)
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

func (o IntegerType) render() bool {
	return true
}
func (o *IntegerType) marshal(out raml.AnyMap) error {
	o.DataType.log.Trace("internal.IntegerType.marshal")
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

func (o *IntegerType) assign(key, val interface{}, log *logrus.Entry) error {
	log.Trace("internal.IntegerType.assign")
	switch key {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind, log); err != nil {
			return xlog.Error(log, err)
		} else if err := ex.UnmarshalRAML(val, log); err != nil {
			return err
		} else {
			o.example = ex.(raml.IntegerExample)
		}
		return nil
	case rmeta.KeyExamples:
		return o.examples.UnmarshalRAML(val, log)
	case rmeta.KeyEnum:
		arr, err := assign.AsAnyList(val, log)
		if err != nil {
			return xlog.Error(log, "the enum facet must be an array. " + err.Error())
		}
		for i := range arr {
			
			l2 := xlog.AddPath(log, i)
			if tmp, ok := arr[i].(int64); ok{
				o.enum = append(o.enum, tmp)
			} else {
				return xlog.Errorf(l2,
					"enum entries for a(n) integer datatype must be of type " +
						"integer.  expected int64, got %s",
					reflect.TypeOf(arr[i]))
			}
			
		}
		return nil
	case rmeta.KeyRequired:
		return assign.AsBool(val, &o.required, log)
	}
	
	switch key {
	case rmeta.KeyMinimum:
		return assign.AsInt64Ptr(val, &o.minimum, log)
	case rmeta.KeyMaximum:
		return assign.AsInt64Ptr(val, &o.maximum, log)
	case rmeta.KeyFormat:
		if val, err := IntegerFormatSortingHat(val, log); err != nil {
			return err
		} else {
			o.format = val
			return nil
		}
	case rmeta.KeyMultipleOf:
		return assign.AsFloat64Ptr(val, &o.multipleOf, log)
	}

	return o.ExtendedDataType.assign(key, val, log)
}

