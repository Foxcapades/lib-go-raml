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

// NewDateOnlyType returns a new internal implementation
// of the raml.DateOnlyType interface.
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
func NewDateOnlyType(log *logrus.Entry) *DateOnlyType {
	log = xlog.WithType(log, "internal.DateOnlyType")

	out := &DateOnlyType{
		examples: NewDateOnlyExampleMap(log),
	}
	
	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeDateOnly, log, out)

	return out
}

// DateOnlyType is a default generated implementation of
// the raml.DateOnlyType interface
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
type DateOnlyType struct {
	*ExtendedDataType

	def      *string
	example  raml.DateOnlyExample
	examples raml.DateOnlyExampleMap
	enum     []string
	
}

func (o *DateOnlyType) SetType(s string) raml.DateOnlyType {
	o.schema = s
	return o
}

func (o *DateOnlyType) Default() option.String {
	return option.NewMaybeString(o.def)
}

func (o *DateOnlyType) SetDefault(i string) raml.DateOnlyType {
	o.def = &i
	return o
}

func (o *DateOnlyType) UnsetDefault() raml.DateOnlyType {
	o.def = nil
	return o
}

func (o *DateOnlyType) Example() raml.DateOnlyExample {
	return o.example
}

func (o *DateOnlyType) SetExample(ex raml.DateOnlyExample) raml.DateOnlyType {
	o.example = ex
	return o
}

func (o *DateOnlyType) UnsetExample() raml.DateOnlyType {
	o.example = nil
	return o
}

func (o *DateOnlyType) Examples() raml.DateOnlyExampleMap {
	return o.examples
}

func (o *DateOnlyType) SetExamples(examples raml.DateOnlyExampleMap) raml.DateOnlyType {
	if examples == nil {
		return o.UnsetExamples()
	}
	o.examples = examples
	return o
}

func (o *DateOnlyType) UnsetExamples() raml.DateOnlyType {
	o.examples = NewDateOnlyExampleMap(o.DataType.log)
	return o
}

func (o *DateOnlyType) SetDisplayName(s string) raml.DateOnlyType {
	o.displayName = &s
	return o
}

func (o *DateOnlyType) UnsetDisplayName() raml.DateOnlyType {
	o.displayName = nil
	return o
}

func (o *DateOnlyType) SetDescription(s string) raml.DateOnlyType {
	o.description = &s
	return o
}

func (o *DateOnlyType) UnsetDescription() raml.DateOnlyType {
	o.description = nil
	return o
}

func (o *DateOnlyType) SetAnnotations(annotations raml.AnnotationMap) raml.DateOnlyType {
	if annotations == nil {
		return o.UnsetAnnotations()
	}
	o.hasAnnotations.mp = annotations
	return o
}

func (o *DateOnlyType) UnsetAnnotations() raml.DateOnlyType {
	o.hasAnnotations.mp = NewAnnotationMap(o.DataType.log)
	return o
}

func (o *DateOnlyType) SetFacetDefinitions(facets raml.FacetMap) raml.DateOnlyType {
	if facets == nil {
		return o.UnsetFacetDefinitions()
	}
	o.facets = facets
	return o
}

func (o *DateOnlyType) UnsetFacetDefinitions() raml.DateOnlyType {
	o.facets = NewFacetMap(o.DataType.log)
	return o
}

func (o *DateOnlyType) SetXml(x raml.Xml) raml.DateOnlyType {
	o.xml = x
	return o
}

func (o *DateOnlyType) UnsetXml() raml.DateOnlyType {
	o.xml = nil
	return o
}

func (o *DateOnlyType) Enum() []string {
	return o.enum
}

func (o *DateOnlyType) SetEnum(i []string) raml.DateOnlyType {
	o.enum = i
	return o
}

func (o *DateOnlyType) UnsetEnum() raml.DateOnlyType {
	o.enum = nil
	return o
}

func (o *DateOnlyType) SetExtraFacets(facets raml.AnyMap) raml.DateOnlyType {
	if facets == nil {
		return o.UnsetExtraFacets()
	}
	o.hasExtra.mp = facets
	return o
}

func (o *DateOnlyType) UnsetExtraFacets() raml.DateOnlyType {
	o.hasExtra.mp = NewAnyMap(o.DataType.log)
	return o
}

func (o *DateOnlyType) SetRequired(b bool) raml.DateOnlyType {
	o.required = b
	return o
}

func (o *DateOnlyType) marshal(out raml.AnyMap) error {
	o.DataType.log.Trace("internal.DateOnlyType.marshal")
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

func (o *DateOnlyType) assign(key, val interface{}, log *logrus.Entry) error {
	log.Trace("internal.DateOnlyType.assign")
	switch key {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind, log); err != nil {
			return xlog.Error(log, err)
		} else if err := ex.UnmarshalRAML(val, log); err != nil {
			return err
		} else {
			o.example = ex.(raml.DateOnlyExample)
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
					"enum entries for a(n) date-only datatype must be of type " +
						"date-only.  expected string, got %s",
					reflect.TypeOf(arr[i]))
			}
			
		}
		return nil
	case rmeta.KeyRequired:
		return assign.AsBool(val, &o.required, log)
	}
	
	return o.ExtendedDataType.assign(key, val, log)
}

