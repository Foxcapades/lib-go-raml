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

// NewFileType returns a new internal implementation
// of the raml.FileType interface.
//
// Generated @ 2020-07-02T14:31:30.98374873-04:00
func NewFileType() *FileType {
	out := &FileType{
		examples: raml.NewFileExampleMap(0),
	}
	
	out.minLength = rmeta.FileDefaultMinLength
	out.maxLength = rmeta.FileDefaultMaxLength

	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeFile, out)

	return out
}

// FileType is a default generated implementation of
// the raml.FileType interface
//
// Generated @ 2020-07-02T14:31:30.98374873-04:00
type FileType struct {
	*ExtendedDataType

	def      *interface{}
	example  raml.FileExample
	examples raml.FileExampleMap
	enum     []interface{}
	fileTypes []string
	minLength uint
	maxLength uint
}

func (o *FileType) SetType(s string) raml.FileType {
	o.schema = s
	return o
}

func (o *FileType) Default() option.Untyped {
	return option.NewMaybeUntyped(o.def)
}

func (o *FileType) SetDefault(i interface{}) raml.FileType {
	o.def = &i
	return o
}

func (o *FileType) UnsetDefault() raml.FileType {
	o.def = nil
	return o
}

func (o *FileType) Example() raml.FileExample {
	return o.example
}

func (o *FileType) SetExample(ex raml.FileExample) raml.FileType {
	o.example = ex
	return o
}

func (o *FileType) UnsetExample() raml.FileType {
	o.example = nil
	return o
}

func (o *FileType) Examples() raml.FileExampleMap {
	return o.examples
}

func (o *FileType) SetExamples(examples raml.FileExampleMap) raml.FileType {
	if examples == nil {
		return o.UnsetExamples()
	}

	o.examples = examples
	return o
}

func (o *FileType) UnsetExamples() raml.FileType {
	o.examples = raml.NewFileExampleMap(0)
	return o
}

func (o *FileType) SetDisplayName(s string) raml.FileType {
	o.displayName = &s
	return o
}

func (o *FileType) UnsetDisplayName() raml.FileType {
	o.displayName = nil
	return o
}

func (o *FileType) SetDescription(s string) raml.FileType {
	o.description = &s
	return o
}

func (o *FileType) UnsetDescription() raml.FileType {
	o.description = nil
	return o
}

func (o *FileType) SetAnnotations(annotations raml.AnnotationMap) raml.FileType {
	if annotations == nil {
		return o.UnsetAnnotations()
	}

	o.hasAnnotations.mp = annotations
	return o
}

func (o *FileType) UnsetAnnotations() raml.FileType {
	o.hasAnnotations.mp = raml.NewAnnotationMap(0)
	return o
}

func (o *FileType) SetFacetDefinitions(facets raml.FacetMap) raml.FileType {
	if facets == nil {
		return o.UnsetFacetDefinitions()
	}

	o.facets = facets
	return o
}

func (o *FileType) UnsetFacetDefinitions() raml.FileType {
	o.facets = raml.NewFacetMap(0)
	return o
}

func (o *FileType) SetXML(x raml.XML) raml.FileType {
	o.xml = x
	return o
}

func (o *FileType) UnsetXML() raml.FileType {
	o.xml = nil
	return o
}

func (o *FileType) Enum() []interface{} {
	return o.enum
}

func (o *FileType) SetEnum(i []interface{}) raml.FileType {
	o.enum = i
	return o
}

func (o *FileType) UnsetEnum() raml.FileType {
	o.enum = nil
	return o
}

func (o *FileType) SetExtraFacets(facets raml.AnyMap) raml.FileType {
	if facets == nil {
		return o.UnsetExtraFacets()
	}

	o.hasExtra.mp = facets
	return o
}

func (o *FileType) UnsetExtraFacets() raml.FileType {
	o.hasExtra.mp = raml.NewAnyMap(0)
	return o
}

func (o *FileType) SetRequired(b bool) raml.FileType {
	o.required = b
	return o
}

func (o *FileType) FileTypes() []string {
	return o.fileTypes
}

func (o *FileType) SetFileTypes(val []string) raml.FileType {
	o.fileTypes = val
	return o
}

func (o *FileType) UnsetFileTypes() raml.FileType {
	o.fileTypes = nil
	return o
}

func (o *FileType) MinLength() uint {
	return o.minLength
}

func (o *FileType) SetMinLength(min uint) raml.FileType {
	o.minLength = min
	return o
}

func (o *FileType) MaxLength() uint {
	return o.maxLength
}

func (o *FileType) SetMaxLength(u uint) raml.FileType {
	o.maxLength = u
	return o
}

func (o *FileType) marshal(out raml.AnyMap) error {
	logrus.Trace("internal.FileType.marshal")
	out.PutIfNotNil(rmeta.KeyDefault, o.def)

	if err := o.ExtendedDataType.marshal(out); err != nil {
		return err
	}
	out.PutIfNotNil(rmeta.KeyFileTypes, o.fileTypes)

	if o.minLength != rmeta.FileDefaultMinLength {
		out.Put(rmeta.KeyMinLength, o.minLength)
	}

	if o.maxLength != rmeta.FileDefaultMaxLength {
		out.Put(rmeta.KeyMaxLength, o.maxLength)
	}
	out.PutIfNotNil(rmeta.KeyEnum, o.enum).
		PutIfNotNil(rmeta.KeyExample, o.example)

	if o.examples.Len() > 0 {
		out.PutIfNotNil(rmeta.KeyExamples, o.examples)
	}

	return nil
}

func (o *FileType) assign(key, val *yaml.Node) error {
	switch key.Value {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind); err != nil {
			return err
		} else if err := ex.UnmarshalRAML(val); err != nil {
			return err
		} else {
			o.example = ex.(raml.FileExample)
		}

		return nil
	case rmeta.KeyExamples:
		return UnmarshalFileExampleMapRAML(o.examples, val)
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
	case rmeta.KeyFileTypes:
		return assign.AsStringList(val, &o.fileTypes)
	case rmeta.KeyMinLength:
		return assign.ToUint(val, &o.minLength)
	case rmeta.KeyMaxLength:
		return assign.ToUint(val, &o.maxLength)
	}

	return o.ExtendedDataType.assign(key, val)
}

