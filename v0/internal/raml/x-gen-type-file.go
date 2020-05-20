package raml

import (
	

	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
)

// NewFileType returns a new internal implementation
// of the raml.FileType interface.
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
func NewFileType(log *logrus.Entry) *FileType {
	log = xlog.WithType(log, "internal.FileType")

	out := &FileType{
		examples: NewFileExampleMap(log),
	}
	
	out.minLength = rmeta.FileDefaultMinLength
	out.maxLength = rmeta.FileDefaultMaxLength

	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeFile, log, out)

	return out
}

// FileType is a default generated implementation of
// the raml.FileType interface
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
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
	o.examples = NewFileExampleMap(o.DataType.log)
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
	o.hasAnnotations.mp = NewAnnotationMap(o.DataType.log)
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
	o.facets = NewFacetMap(o.DataType.log)
	return o
}

func (o *FileType) SetXml(x raml.Xml) raml.FileType {
	o.xml = x
	return o
}

func (o *FileType) UnsetXml() raml.FileType {
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
	o.hasExtra.mp = NewAnyMap(o.DataType.log)
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

func (o FileType) render() bool {
	return true
}
func (o *FileType) marshal(out raml.AnyMap) error {
	o.DataType.log.Trace("internal.FileType.marshal")
	out.PutNonNil(rmeta.KeyDefault, o.def)

	if err := o.ExtendedDataType.marshal(out); err != nil {
		return err
	}
	out.PutNonNil(rmeta.KeyFileTypes, o.fileTypes)
	if o.minLength != rmeta.FileDefaultMinLength {
		out.Put(rmeta.KeyMinLength, o.minLength)
	}
	if o.maxLength != rmeta.FileDefaultMaxLength {
		out.Put(rmeta.KeyMaxLength, o.maxLength)
	}
	out.PutNonNil(rmeta.KeyEnum, o.enum).
		PutNonNil(rmeta.KeyExample, o.example)

	if o.examples.Len() > 0 {
		out.PutNonNil(rmeta.KeyExamples, o.examples)
	}

	return nil
}

func (o *FileType) assign(key, val interface{}, log *logrus.Entry) error {
	log.Trace("internal.FileType.assign")
	switch key {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind, log); err != nil {
			return xlog.Error(log, err)
		} else if err := ex.UnmarshalRAML(val, log); err != nil {
			return err
		} else {
			o.example = ex.(raml.FileExample)
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
			
			o.enum = append(o.enum, arr[i])
			
		}
		return nil
	case rmeta.KeyRequired:
		return assign.AsBool(val, &o.required, log)
	}
	
	switch key {
	case rmeta.KeyFileTypes:
		return assign.AsStringList(val, &o.fileTypes, log)
	case rmeta.KeyMinLength:
		return assign.ToUint(val, &o.minLength)
	case rmeta.KeyMaxLength:
		return assign.ToUint(val, &o.maxLength)
	}

	return o.ExtendedDataType.assign(key, val, log)
}

