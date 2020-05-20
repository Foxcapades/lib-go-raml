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

// NewDatetimeOnlyType returns a new internal implementation
// of the raml.DatetimeOnlyType interface.
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
func NewDatetimeOnlyType(log *logrus.Entry) *DatetimeOnlyType {
	log = xlog.WithType(log, "internal.DatetimeOnlyType")

	out := &DatetimeOnlyType{
		examples: NewDatetimeOnlyExampleMap(log),
	}
	
	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeDatetimeOnly, log, out)

	return out
}

// DatetimeOnlyType is a default generated implementation of
// the raml.DatetimeOnlyType interface
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
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
	o.examples = NewDatetimeOnlyExampleMap(o.DataType.log)
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
	o.hasAnnotations.mp = NewAnnotationMap(o.DataType.log)
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
	o.facets = NewFacetMap(o.DataType.log)
	return o
}

func (o *DatetimeOnlyType) SetXml(x raml.Xml) raml.DatetimeOnlyType {
	o.xml = x
	return o
}

func (o *DatetimeOnlyType) UnsetXml() raml.DatetimeOnlyType {
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
	o.hasExtra.mp = NewAnyMap(o.DataType.log)
	return o
}

func (o *DatetimeOnlyType) SetRequired(b bool) raml.DatetimeOnlyType {
	o.required = b
	return o
}

func (o *DatetimeOnlyType) marshal(out raml.AnyMap) error {
	o.DataType.log.Trace("internal.DatetimeOnlyType.marshal")
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

func (o *DatetimeOnlyType) assign(key, val interface{}, log *logrus.Entry) error {
	log.Trace("internal.DatetimeOnlyType.assign")
	switch key {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind, log); err != nil {
			return xlog.Error(log, err)
		} else if err := ex.UnmarshalRAML(val, log); err != nil {
			return err
		} else {
			o.example = ex.(raml.DatetimeOnlyExample)
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
					"enum entries for a(n) datetime-only datatype must be of type " +
						"datetime-only.  expected string, got %s",
					reflect.TypeOf(arr[i]))
			}
			
		}
		return nil
	case rmeta.KeyRequired:
		return assign.AsBool(val, &o.required, log)
	}
	
	return o.ExtendedDataType.assign(key, val, log)
}

