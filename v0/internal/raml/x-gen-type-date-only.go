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

// NewDateOnlyType returns a new internal implementation
// of the raml.DateOnlyType interface.
//
// Generated @ 2020-07-06T13:52:18.671712454-04:00
func NewDateOnlyType() *DateOnlyType {
	out := &DateOnlyType{
		examples: raml.NewDateOnlyExampleMap(0),
	}

	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeDateOnly, out)

	return out
}

// DateOnlyType is a default generated implementation of
// the raml.DateOnlyType interface
//
// Generated @ 2020-07-06T13:52:18.671712454-04:00
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
	o.examples = raml.NewDateOnlyExampleMap(0)
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
	o.hasAnnotations.mp = raml.NewAnnotationMap(0)
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
	o.facets = raml.NewFacetMap(0)
	return o
}

func (o *DateOnlyType) SetXML(x raml.XML) raml.DateOnlyType {
	o.xml = x
	return o
}

func (o *DateOnlyType) UnsetXML() raml.DateOnlyType {
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
	o.hasExtra.mp = raml.NewAnyMap(0)
	return o
}

func (o *DateOnlyType) SetRequired(b bool) raml.DateOnlyType {
	o.required = b
	return o
}

func (o *DateOnlyType) marshal(out raml.AnyMap) error {
	logrus.Trace("internal.DateOnlyType.marshal")
	out.PutIfNotNil(rmeta.KeyDefault, o.def)

	if err := o.ExtendedDataType.marshal(out); err != nil {
		return err
	}

	out.PutIfNotNil(rmeta.KeyEnum, o.enum).
		PutIfNotNil(rmeta.KeyExample, o.example)

	if o.examples.Len() > 0 {
		out.PutIfNotNil(rmeta.KeyExamples, o.examples)
	}

	return nil
}

func (o *DateOnlyType) assign(key, val *yaml.Node) error {
	switch key.Value {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind); err != nil {
			return err
		} else if err := ex.UnmarshalRAML(val); err != nil {
			return err
		} else {
			o.example = ex.(raml.DateOnlyExample)
		}

		return nil
	case rmeta.KeyExamples:
		return UnmarshalDateOnlyExampleMapRAML(o.examples, val)
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

	return o.ExtendedDataType.assign(key, val)
}
