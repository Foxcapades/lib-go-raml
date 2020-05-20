package raml

import "github.com/Foxcapades/goop/v1/pkg/option"
// FileType generated @ 2020-05-20T00:33:46.349824232-04:00
type FileType interface {
	Unmarshaler
	Marshaler
	ExtendedDataType

	// SetType sets the parent type name, this does not change
	// the underlying kind of the DataType.
	SetType(string) FileType

	// Default returns an option of the current default value
	// which will be empty if the value is not set.
	Default() option.Untyped

	// SetDefault sets the default value for the current
	// DataType definition.
	SetDefault(interface{}) FileType

	// UnsetDefault removes the default value definition from
	// the current DataType definition.
	//
	// If no default was previously set, this method does
	// nothing.
	UnsetDefault() FileType

	// Example returns the singular example value assigned to
	// the current DataType definition.
	Example() FileExample

	// SetExample sets the singular example value assigned to
	// the current DataType definition.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExample.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExample(ex FileExample) FileType

	// UnsetExample removes the singular example value from
	// the current DataType definition.
	UnsetExample() FileType

	// Examples returns a mutable map of examples for the
	// current DataType definition.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	Examples() FileExampleMap

	// SetExamples replaces the current DataType definition's
	// example map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// ass calling UnsetExamples.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExamples(examples FileExampleMap) FileType

	// UnsetExamples clears the current DataType definition's
	// example map.
	UnsetExamples() FileType

	// SetDisplayName sets the current DataType definition's
	// displayName value.
	SetDisplayName(string) FileType

	// UnsetDisplayName removes the current DataType
	// definition's displayName value.
	UnsetDisplayName() FileType

	// SetDescription sets the current DataType definition's
	// description value.
	SetDescription(string) FileType

	// UnsetDescription removes the current DataType
	// definition's description value.
	UnsetDescription() FileType

	// SetAnnotations replaces the current DataType
	// definition's annotation map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) FileType

	// UnsetAnnotations clears the current DataType
	// definition's applied annotation map.
	UnsetAnnotations() FileType

	// SetFacetDefinitions replaces the current DataType
	// definition's custom facet definitions with the given
	// value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetFacets.
	SetFacetDefinitions(facets FacetMap) FileType

	// UnsetFacetDefinitions removes all custom facets defined
	// on the current DataType definition.
	UnsetFacetDefinitions() FileType

	// SetXml sets the current DataType definition's xml
	// serialization settings object to the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetXml.
	SetXml(Xml) FileType

	// Removes the xml serialization settings object from the
	// current DataType definition.
	UnsetXml() FileType

	// Enum returns a slice of the enum values assigned to the
	// current DataType definition.
	Enum() []interface{}

	// SetEnum replaces the current DataType definition's enum
	// slice with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetEnum.
	SetEnum([]interface{}) FileType

	// UnsetEnum clears the enum values from the current
	// DataType definition.
	UnsetEnum() FileType

	// SetExtraFacets replaces the facets applied to the
	// current DataType definition with the given values.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExtraFacets.
	SetExtraFacets(facets AnyMap) FileType

	// UnsetExtraFacets removes all extra facets from the
	// current DataType definition.
	UnsetExtraFacets() FileType

	SetRequired(bool) FileType

	// FileTypes returns a slice of the "fileTypes" facet set
	// on the current file type definition.
	FileTypes() []string

	// SetFileTypes replaces the values of the "fileTypes"
	// facet on the current file type definition with the
	// given values.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetFileTypes.
	SetFileTypes([]string) FileType

	// UnsetFileTypes clears the "fileTypes" facet for the
	// current file type definition.
	UnsetFileTypes() FileType

	// MinLength returns the value for the "minLength" facet
	// on the current file type definition.
	//
	// The default value is 0.  If this facet is set to its
	// default value when rendered, it will not appear in the
	// output RAML.
	MinLength() uint

	// SetMinLength sets the value of the "minLength" facet on
	// the current file type definition.
	//
	// The default value is 0.  If this facet is set to its
	// default value when rendered, it will not appear in the
	// output RAML.
	SetMinLength(min uint) FileType

	// MaxLength returns the value for the "maxLength" facet
	// on the current file type definition.
	//
	// The default value is 2_147_483_647.  If this facet is
	// set to its default value when rendered, it will not
	// appear in the output RAML.
	MaxLength() uint

	// SetMaxLength sets the value of the "maxLength" facet on
	// the current file type definition.
	//
	// The default value is 2_147_483_647.  If this facet is
	// set to its default value when rendered, it will not
	// appear in the output RAML.
	SetMaxLength(max uint) FileType

}
