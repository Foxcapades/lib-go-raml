package raml

import (
	"github.com/Foxcapades/lib-go-raml/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// NewArrayType returns a new internal implementation
// of the raml.ArrayType interface.
//
// Generated @ 2020-05-25T19:07:00.757913962-04:00
func NewArrayType() *ArrayType {
	out := &ArrayType{
		examples: raml.NewArrayExampleMap(0),
	}

	out.minItems = rmeta.ArrayDefaultMinItems
	out.maxItems = rmeta.ArrayDefaultMaxItems
	out.uniqueItems = rmeta.ArrayDefaultUniqueItems

	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeArray, out)

	return out
}

// ArrayType is a default generated implementation of
// the raml.ArrayType interface
//
// Generated @ 2020-05-25T19:07:00.757913962-04:00
type ArrayType struct {
	*ExtendedDataType

	def         []interface{}
	example     raml.ArrayExample
	examples    raml.ArrayExampleMap
	enum        []interface{}
	uniqueItems bool
	minItems    uint
	maxItems    uint
	items       raml.DataType
}

func (o *ArrayType) SetType(s string) raml.ArrayType {
	o.schema = s
	return o
}

func (o *ArrayType) Default() []interface{} {
	return o.def
}

func (o *ArrayType) SetDefault(i []interface{}) raml.ArrayType {
	o.def = i
	return o
}

func (o *ArrayType) UnsetDefault() raml.ArrayType {
	o.def = nil
	return o
}

func (o *ArrayType) Example() raml.ArrayExample {
	return o.example
}

func (o *ArrayType) SetExample(ex raml.ArrayExample) raml.ArrayType {
	o.example = ex
	return o
}

func (o *ArrayType) UnsetExample() raml.ArrayType {
	o.example = nil
	return o
}

func (o *ArrayType) Examples() raml.ArrayExampleMap {
	return o.examples
}

func (o *ArrayType) SetExamples(examples raml.ArrayExampleMap) raml.ArrayType {
	if examples == nil {
		return o.UnsetExamples()
	}

	o.examples = examples
	return o
}

func (o *ArrayType) UnsetExamples() raml.ArrayType {
	o.examples = raml.NewArrayExampleMap(0)
	return o
}

func (o *ArrayType) SetDisplayName(s string) raml.ArrayType {
	o.displayName = &s
	return o
}

func (o *ArrayType) UnsetDisplayName() raml.ArrayType {
	o.displayName = nil
	return o
}

func (o *ArrayType) SetDescription(s string) raml.ArrayType {
	o.description = &s
	return o
}

func (o *ArrayType) UnsetDescription() raml.ArrayType {
	o.description = nil
	return o
}

func (o *ArrayType) SetAnnotations(annotations raml.AnnotationMap) raml.ArrayType {
	if annotations == nil {
		return o.UnsetAnnotations()
	}

	o.hasAnnotations.mp = annotations
	return o
}

func (o *ArrayType) UnsetAnnotations() raml.ArrayType {
	o.hasAnnotations.mp = raml.NewAnnotationMap(0)
	return o
}

func (o *ArrayType) SetFacetDefinitions(facets raml.FacetMap) raml.ArrayType {
	if facets == nil {
		return o.UnsetFacetDefinitions()
	}

	o.facets = facets
	return o
}

func (o *ArrayType) UnsetFacetDefinitions() raml.ArrayType {
	o.facets = raml.NewFacetMap(0)
	return o
}

func (o *ArrayType) SetXML(x raml.XML) raml.ArrayType {
	o.xml = x
	return o
}

func (o *ArrayType) UnsetXML() raml.ArrayType {
	o.xml = nil
	return o
}

func (o *ArrayType) Enum() []interface{} {
	return o.enum
}

func (o *ArrayType) SetEnum(i []interface{}) raml.ArrayType {
	o.enum = i
	return o
}

func (o *ArrayType) UnsetEnum() raml.ArrayType {
	o.enum = nil
	return o
}

func (o *ArrayType) SetExtraFacets(facets raml.AnyMap) raml.ArrayType {
	if facets == nil {
		return o.UnsetExtraFacets()
	}

	o.hasExtra.mp = facets
	return o
}

func (o *ArrayType) UnsetExtraFacets() raml.ArrayType {
	o.hasExtra.mp = raml.NewAnyMap(0)
	return o
}

func (o *ArrayType) SetRequired(b bool) raml.ArrayType {
	o.required = b
	return o
}

func (o *ArrayType) UniqueItems() bool {
	return o.uniqueItems
}

func (o *ArrayType) SetUniqueItems(val bool) raml.ArrayType {
	o.uniqueItems = val
	return o
}

func (o *ArrayType) MinItems() uint {
	return o.minItems
}

func (o *ArrayType) SetMinItems(min uint) raml.ArrayType {
	o.minItems = min
	return o
}

func (o *ArrayType) MaxItems() uint {
	return o.maxItems
}

func (o *ArrayType) SetMaxItems(u uint) raml.ArrayType {
	o.maxItems = u
	return o
}

func (o *ArrayType) Items() raml.DataType {
	return o.items
}

func (o *ArrayType) SetItems(val raml.DataType) raml.ArrayType {
	o.items = val
	return o
}

func (o *ArrayType) UnsetItems() raml.ArrayType {
	o.items = nil
	return o
}

func (o *ArrayType) marshal(out raml.AnyMap) error {
	logrus.Trace("internal.ArrayType.marshal")
	out.PutIfNotNil(rmeta.KeyDefault, o.def)

	if err := o.ExtendedDataType.marshal(out); err != nil {
		return err
	}

	if o.uniqueItems != rmeta.ArrayDefaultUniqueItems {
		out.Put(rmeta.KeyUniqueItems, o.uniqueItems)
	}

	if o.minItems != rmeta.ArrayDefaultMinItems {
		out.Put(rmeta.KeyMinItems, o.minItems)
	}

	if o.maxItems != rmeta.ArrayDefaultMaxItems {
		out.Put(rmeta.KeyMaxItems, o.maxItems)
	}

	out.PutIfNotNil(rmeta.KeyItems, o.items)

	out.PutIfNotNil(rmeta.KeyEnum, o.enum).
		PutIfNotNil(rmeta.KeyExample, o.example)

	if o.examples.Len() > 0 {
		out.PutIfNotNil(rmeta.KeyExamples, o.examples)
	}

	return nil
}

func (o *ArrayType) assign(key, val *yaml.Node) error {
	switch key.Value {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind); err != nil {
			return err
		} else if err := ex.UnmarshalRAML(val); err != nil {
			return err
		} else {
			o.example = ex.(raml.ArrayExample)
		}

		return nil
	case rmeta.KeyExamples:
		return UnmarshalArrayExampleMapRAML(o.examples, val)
	case rmeta.KeyEnum:
		return xyml.SequenceForEach(val, func(cur *yaml.Node) error {
			o.enum = append(o.enum, val)

			return nil
		})
		return nil
	case rmeta.KeyRequired:
		return assign.AsBool(val, &o.required)
	}

	switch key.Value {
	case rmeta.KeyUniqueItems:
		return assign.AsBool(val, &o.uniqueItems)
	case rmeta.KeyMinItems:
		return assign.ToUint(val, &o.minItems)
	case rmeta.KeyMaxItems:
		return assign.ToUint(val, &o.maxItems)
	case rmeta.KeyItems:
		if val, err := TypeSortingHat(val); err == nil {
			o.items = val
			return nil
		} else {
			return err
		}
	}

	return o.ExtendedDataType.assign(key, val)
}
