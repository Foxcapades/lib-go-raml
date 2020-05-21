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

// NewBoolType returns a new internal implementation
// of the raml.BoolType interface.
//
// Generated @ 2020-05-20T21:46:01.015916886-04:00
func NewBoolType() *BoolType {
	out := &BoolType{
		examples: NewBoolExampleMap(),
	}

	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeBool, out)

	return out
}

// BoolType is a default generated implementation of
// the raml.BoolType interface
//
// Generated @ 2020-05-20T21:46:01.015916886-04:00
type BoolType struct {
	*ExtendedDataType

	def      *bool
	example  raml.BoolExample
	examples raml.BoolExampleMap
	enum     []bool
}

func (o *BoolType) SetType(s string) raml.BoolType {
	o.schema = s
	return o
}

func (o *BoolType) Default() option.Bool {
	return option.NewMaybeBool(o.def)
}

func (o *BoolType) SetDefault(i bool) raml.BoolType {
	o.def = &i
	return o
}

func (o *BoolType) UnsetDefault() raml.BoolType {
	o.def = nil
	return o
}

func (o *BoolType) Example() raml.BoolExample {
	return o.example
}

func (o *BoolType) SetExample(ex raml.BoolExample) raml.BoolType {
	o.example = ex
	return o
}

func (o *BoolType) UnsetExample() raml.BoolType {
	o.example = nil
	return o
}

func (o *BoolType) Examples() raml.BoolExampleMap {
	return o.examples
}

func (o *BoolType) SetExamples(examples raml.BoolExampleMap) raml.BoolType {
	if examples == nil {
		return o.UnsetExamples()
	}

	o.examples = examples
	return o
}

func (o *BoolType) UnsetExamples() raml.BoolType {
	o.examples = NewBoolExampleMap()
	return o
}

func (o *BoolType) SetDisplayName(s string) raml.BoolType {
	o.displayName = &s
	return o
}

func (o *BoolType) UnsetDisplayName() raml.BoolType {
	o.displayName = nil
	return o
}

func (o *BoolType) SetDescription(s string) raml.BoolType {
	o.description = &s
	return o
}

func (o *BoolType) UnsetDescription() raml.BoolType {
	o.description = nil
	return o
}

func (o *BoolType) SetAnnotations(annotations raml.AnnotationMap) raml.BoolType {
	if annotations == nil {
		return o.UnsetAnnotations()
	}

	o.hasAnnotations.mp = annotations
	return o
}

func (o *BoolType) UnsetAnnotations() raml.BoolType {
	o.hasAnnotations.mp = NewAnnotationMap()
	return o
}

func (o *BoolType) SetFacetDefinitions(facets raml.FacetMap) raml.BoolType {
	if facets == nil {
		return o.UnsetFacetDefinitions()
	}

	o.facets = facets
	return o
}

func (o *BoolType) UnsetFacetDefinitions() raml.BoolType {
	o.facets = NewFacetMap()
	return o
}

func (o *BoolType) SetXML(x raml.XML) raml.BoolType {
	o.xml = x
	return o
}

func (o *BoolType) UnsetXML() raml.BoolType {
	o.xml = nil
	return o
}

func (o *BoolType) Enum() []bool {
	return o.enum
}

func (o *BoolType) SetEnum(i []bool) raml.BoolType {
	o.enum = i
	return o
}

func (o *BoolType) UnsetEnum() raml.BoolType {
	o.enum = nil
	return o
}

func (o *BoolType) SetExtraFacets(facets raml.AnyMap) raml.BoolType {
	if facets == nil {
		return o.UnsetExtraFacets()
	}

	o.hasExtra.mp = facets
	return o
}

func (o *BoolType) UnsetExtraFacets() raml.BoolType {
	o.hasExtra.mp = NewAnyMap()
	return o
}

func (o *BoolType) SetRequired(b bool) raml.BoolType {
	o.required = b
	return o
}

func (o *BoolType) marshal(out raml.AnyMap) error {
	logrus.Trace("internal.BoolType.marshal")
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

func (o *BoolType) assign(key, val *yaml.Node) error {
	switch key.Value {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind); err != nil {
			return err
		} else if err := ex.UnmarshalRAML(val); err != nil {
			return err
		} else {
			o.example = ex.(raml.BoolExample)
		}

		return nil
	case rmeta.KeyExamples:
		return o.examples.UnmarshalRAML(val)
	case rmeta.KeyEnum:
		return xyml.ForEachList(val, func(cur *yaml.Node) error {
			if val, err := xyml.ToBool(cur); err != nil {
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

	return o.ExtendedDataType.assign(key, val)
}
