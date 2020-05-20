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

// NewDatetimeType returns a new internal implementation
// of the raml.DatetimeType interface.
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
func NewDatetimeType(log *logrus.Entry) *DatetimeType {
	log = xlog.WithType(log, "internal.DatetimeType")

	out := &DatetimeType{
		examples: NewDatetimeExampleMap(log),
	}
	
	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeDatetime, log, out)

	return out
}

// DatetimeType is a default generated implementation of
// the raml.DatetimeType interface
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
type DatetimeType struct {
	*ExtendedDataType

	def      *string
	example  raml.DatetimeExample
	examples raml.DatetimeExampleMap
	enum     []string
	format     raml.DateFormat
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
	o.examples = NewDatetimeExampleMap(o.DataType.log)
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
	o.hasAnnotations.mp = NewAnnotationMap(o.DataType.log)
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
	o.facets = NewFacetMap(o.DataType.log)
	return o
}

func (o *DatetimeType) SetXml(x raml.Xml) raml.DatetimeType {
	o.xml = x
	return o
}

func (o *DatetimeType) UnsetXml() raml.DatetimeType {
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
	o.hasExtra.mp = NewAnyMap(o.DataType.log)
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

func (o DatetimeType) render() bool {
	return true
}
func (o *DatetimeType) marshal(out raml.AnyMap) error {
	o.DataType.log.Trace("internal.DatetimeType.marshal")
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

func (o *DatetimeType) assign(key, val interface{}, log *logrus.Entry) error {
	log.Trace("internal.DatetimeType.assign")
	switch key {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind, log); err != nil {
			return xlog.Error(log, err)
		} else if err := ex.UnmarshalRAML(val, log); err != nil {
			return err
		} else {
			o.example = ex.(raml.DatetimeExample)
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
					"enum entries for a(n) datetime datatype must be of type " +
						"datetime.  expected string, got %s",
					reflect.TypeOf(arr[i]))
			}
			
		}
		return nil
	case rmeta.KeyRequired:
		return assign.AsBool(val, &o.required, log)
	}
	
	switch key {
	case rmeta.KeyFormat:
		if val, err := DateFormatSortingHat(val, log); err != nil {
			return err
		} else {
			o.format = val
			return nil
		}
	}

	return o.ExtendedDataType.assign(key, val, log)
}

