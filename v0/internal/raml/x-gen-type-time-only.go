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

// NewTimeOnlyType returns a new internal implementation
// of the raml.TimeOnlyType interface.
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
func NewTimeOnlyType(log *logrus.Entry) *TimeOnlyType {
	log = xlog.WithType(log, "internal.TimeOnlyType")

	out := &TimeOnlyType{
		examples: NewTimeOnlyExampleMap(log),
	}
	
	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeTimeOnly, log, out)

	return out
}

// TimeOnlyType is a default generated implementation of
// the raml.TimeOnlyType interface
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
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
	o.examples = NewTimeOnlyExampleMap(o.DataType.log)
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
	o.hasAnnotations.mp = NewAnnotationMap(o.DataType.log)
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
	o.facets = NewFacetMap(o.DataType.log)
	return o
}

func (o *TimeOnlyType) SetXml(x raml.Xml) raml.TimeOnlyType {
	o.xml = x
	return o
}

func (o *TimeOnlyType) UnsetXml() raml.TimeOnlyType {
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
	o.hasExtra.mp = NewAnyMap(o.DataType.log)
	return o
}

func (o *TimeOnlyType) SetRequired(b bool) raml.TimeOnlyType {
	o.required = b
	return o
}

func (o *TimeOnlyType) marshal(out raml.AnyMap) error {
	o.DataType.log.Trace("internal.TimeOnlyType.marshal")
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

func (o *TimeOnlyType) assign(key, val interface{}, log *logrus.Entry) error {
	log.Trace("internal.TimeOnlyType.assign")
	switch key {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind, log); err != nil {
			return xlog.Error(log, err)
		} else if err := ex.UnmarshalRAML(val, log); err != nil {
			return err
		} else {
			o.example = ex.(raml.TimeOnlyExample)
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
			if tmp, ok := arr[i].(string); ok{
				o.enum = append(o.enum, tmp)
			} else {
				return xlog.Errorf(l2,
					"enum entries for a(n) time-only datatype must be of type " +
						"time-only.  expected string, got %s",
					reflect.TypeOf(arr[i]))
			}
			
		}
		return nil
	case rmeta.KeyRequired:
		return assign.AsBool(val, &o.required, log)
	}
	
	return o.ExtendedDataType.assign(key, val, log)
}

