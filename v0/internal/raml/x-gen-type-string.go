package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// NewStringType returns a new internal implementation
// of the raml.StringType interface.
//
// Generated @ 2020-05-25T19:07:00.757913962-04:00
func NewStringType() *StringType {
	out := &StringType{
		examples: raml.NewStringExampleMap(0),
	}

	out.minLength = rmeta.StringDefaultMinLength

	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeString, out)

	return out
}

// StringType is a default generated implementation of
// the raml.StringType interface
//
// Generated @ 2020-05-25T19:07:00.757913962-04:00
type StringType struct {
	*ExtendedDataType

	def       *string
	example   raml.StringExample
	examples  raml.StringExampleMap
	enum      []string
	pattern   *string
	minLength uint
	maxLength *uint
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
	o.examples = raml.NewStringExampleMap(0)
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
	o.hasAnnotations.mp = raml.NewAnnotationMap(0)
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
	o.facets = raml.NewFacetMap(0)
	return o
}

func (o *StringType) SetXML(x raml.XML) raml.StringType {
	o.xml = x
	return o
}

func (o *StringType) UnsetXML() raml.StringType {
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
	o.hasExtra.mp = raml.NewAnyMap(0)
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

func (o *StringType) marshal(out raml.AnyMap) error {
	logrus.Trace("internal.StringType.marshal")
	out.PutIfNotNil(rmeta.KeyDefault, o.def)

	if err := o.ExtendedDataType.marshal(out); err != nil {
		return err
	}
	out.PutIfNotNil(rmeta.KeyPattern, o.pattern)
	if o.minLength > 0 {
		out.Put(rmeta.KeyMinLength, o.minLength)
	}
	out.PutIfNotNil(rmeta.KeyMaxLength, o.maxLength)
	out.PutIfNotNil(rmeta.KeyEnum, o.enum).
		PutIfNotNil(rmeta.KeyExample, o.example)

	if o.examples.Len() > 0 {
		out.PutIfNotNil(rmeta.KeyExamples, o.examples)
	}

	return nil
}

func (o *StringType) assign(key, val *yaml.Node) error {
	switch key.Value {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind); err != nil {
			return err
		} else if err := ex.UnmarshalRAML(val); err != nil {
			return err
		} else {
			o.example = ex.(raml.StringExample)
		}

		return nil
	case rmeta.KeyExamples:
		return UnmarshalStringExampleMapRAML(o.examples, val)
	case rmeta.KeyEnum:
		return xyml.SequenceForEach(val, func(cur *yaml.Node) error {
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

	switch key.Value {
	case rmeta.KeyPattern:
		return assign.AsStringPtr(val, &o.pattern)
	case rmeta.KeyMinLength:
		return assign.ToUint(val, &o.minLength)
	case rmeta.KeyMaxLength:
		return assign.AsUintPtr(val, &o.maxLength)
	}

	return o.ExtendedDataType.assign(key, val)
}
