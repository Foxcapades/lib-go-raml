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

// NewStringType returns a new internal implementation
// of the raml.StringType interface.
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
func NewStringType(log *logrus.Entry) *StringType {
	log = xlog.WithType(log, "internal.StringType")

	out := &StringType{
		examples: NewStringExampleMap(log),
	}
	
	out.minLength = rmeta.StringDefaultMinLength

	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeString, log, out)

	return out
}

// StringType is a default generated implementation of
// the raml.StringType interface
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
type StringType struct {
	*ExtendedDataType

	def      *string
	example  raml.StringExample
	examples raml.StringExampleMap
	enum     []string
	pattern      *string
	minLength    uint
	maxLength    *uint
}

func (o *StringType) SetType(s string) raml.StringType {
	o.schema = s
	return o
}

func (o *StringType) Default() option.String {
	return option.NewMaybeString(o.def)
}

func (o *StringType) SetDefault(i string) raml.StringType {
	o.def = &i
	return o
}

func (o *StringType) UnsetDefault() raml.StringType {
	o.def = nil
	return o
}

func (o *StringType) Example() raml.StringExample {
	return o.example
}

func (o *StringType) SetExample(ex raml.StringExample) raml.StringType {
	o.example = ex
	return o
}

func (o *StringType) UnsetExample() raml.StringType {
	o.example = nil
	return o
}

func (o *StringType) Examples() raml.StringExampleMap {
	return o.examples
}

func (o *StringType) SetExamples(examples raml.StringExampleMap) raml.StringType {
	if examples == nil {
		return o.UnsetExamples()
	}
	o.examples = examples
	return o
}

func (o *StringType) UnsetExamples() raml.StringType {
	o.examples = NewStringExampleMap(o.DataType.log)
	return o
}

func (o *StringType) SetDisplayName(s string) raml.StringType {
	o.displayName = &s
	return o
}

func (o *StringType) UnsetDisplayName() raml.StringType {
	o.displayName = nil
	return o
}

func (o *StringType) SetDescription(s string) raml.StringType {
	o.description = &s
	return o
}

func (o *StringType) UnsetDescription() raml.StringType {
	o.description = nil
	return o
}

func (o *StringType) SetAnnotations(annotations raml.AnnotationMap) raml.StringType {
	if annotations == nil {
		return o.UnsetAnnotations()
	}
	o.hasAnnotations.mp = annotations
	return o
}

func (o *StringType) UnsetAnnotations() raml.StringType {
	o.hasAnnotations.mp = NewAnnotationMap(o.DataType.log)
	return o
}

func (o *StringType) SetFacetDefinitions(facets raml.FacetMap) raml.StringType {
	if facets == nil {
		return o.UnsetFacetDefinitions()
	}
	o.facets = facets
	return o
}

func (o *StringType) UnsetFacetDefinitions() raml.StringType {
	o.facets = NewFacetMap(o.DataType.log)
	return o
}

func (o *StringType) SetXml(x raml.Xml) raml.StringType {
	o.xml = x
	return o
}

func (o *StringType) UnsetXml() raml.StringType {
	o.xml = nil
	return o
}

func (o *StringType) Enum() []string {
	return o.enum
}

func (o *StringType) SetEnum(i []string) raml.StringType {
	o.enum = i
	return o
}

func (o *StringType) UnsetEnum() raml.StringType {
	o.enum = nil
	return o
}

func (o *StringType) SetExtraFacets(facets raml.AnyMap) raml.StringType {
	if facets == nil {
		return o.UnsetExtraFacets()
	}
	o.hasExtra.mp = facets
	return o
}

func (o *StringType) UnsetExtraFacets() raml.StringType {
	o.hasExtra.mp = NewAnyMap(o.DataType.log)
	return o
}

func (o *StringType) SetRequired(b bool) raml.StringType {
	o.required = b
	return o
}

func (o *StringType) Pattern() option.String {
	return option.NewMaybeString(o.pattern)
}

func (o *StringType) SetPattern(pat string) raml.StringType {
	o.pattern = &pat
	return o
}

func (o *StringType) UnsetPattern() raml.StringType {
	o.pattern = nil
	return o
}

func (o *StringType) MinLength() uint {
	return o.minLength
}

func (o *StringType) SetMinLength(min uint) raml.StringType {
	o.minLength = min
	return o
}

func (o *StringType) MaxLength() option.Uint {
	return option.NewMaybeUint(o.maxLength)
}

func (o *StringType) SetMaxLength(u uint) raml.StringType {
	o.maxLength = &u
	return o
}

func (o *StringType) UnsetMaxLength() raml.StringType {
	o.maxLength = nil
	return o
}

func (o StringType) render() bool {
	return true
}
func (o *StringType) marshal(out raml.AnyMap) error {
	o.DataType.log.Trace("internal.StringType.marshal")
	out.PutNonNil(rmeta.KeyDefault, o.def)

	if err := o.ExtendedDataType.marshal(out); err != nil {
		return err
	}
	out.PutNonNil(rmeta.KeyPattern, o.pattern)
	if o.minLength > 0 {
		out.Put(rmeta.KeyMinLength, o.minLength)
	}
	out.PutNonNil(rmeta.KeyMaxLength, o.maxLength)
	out.PutNonNil(rmeta.KeyEnum, o.enum).
		PutNonNil(rmeta.KeyExample, o.example)

	if o.examples.Len() > 0 {
		out.PutNonNil(rmeta.KeyExamples, o.examples)
	}

	return nil
}

func (o *StringType) assign(key, val interface{}, log *logrus.Entry) error {
	log.Trace("internal.StringType.assign")
	switch key {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind, log); err != nil {
			return xlog.Error(log, err)
		} else if err := ex.UnmarshalRAML(val, log); err != nil {
			return err
		} else {
			o.example = ex.(raml.StringExample)
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
					"enum entries for a(n) string datatype must be of type " +
						"string.  expected string, got %s",
					reflect.TypeOf(arr[i]))
			}
			
		}
		return nil
	case rmeta.KeyRequired:
		return assign.AsBool(val, &o.required, log)
	}
	
	switch key {
	case rmeta.KeyPattern:
		return assign.AsStringPtr(val, &o.pattern, log)
	case rmeta.KeyMinLength:
		return assign.ToUint(val, &o.minLength)
	case rmeta.KeyMaxLength:
		return assign.AsUintPtr(val, &o.maxLength)
	}

	return o.ExtendedDataType.assign(key, val, log)
}

