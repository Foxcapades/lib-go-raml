package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// StringType generated @ 2020-10-19T13:48:24.9771134-04:00
type StringType interface {
	Unmarshaler
	Marshaler
	ExtendedDataType

	// SetType sets the parent type name, this does not change
	// the underlying kind of the DataType.
	SetType(string) StringType

	// Default returns an option of the current default value
	// which will be empty if the value is not set.
	Default() option.String

	// SetDefault sets the default value for the current
	// DataType definition.
	SetDefault(string) StringType

	// UnsetDefault removes the default value definition from
	// the current DataType definition.
	//
	// If no default was previously set, this method does
	// nothing.
	UnsetDefault() StringType

	// Example returns the singular example value assigned to
	// the current DataType definition.
	Example() StringExample

	// SetExample sets the singular example value assigned to
	// the current DataType definition.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExample.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExample(ex StringExample) StringType

	// UnsetExample removes the singular example value from
	// the current DataType definition.
	UnsetExample() StringType

	// Examples returns a mutable map of examples for the
	// current DataType definition.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	Examples() StringExampleMap

	// SetExamples replaces the current DataType definition's
	// example map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// ass calling UnsetExamples.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExamples(examples StringExampleMap) StringType

	// UnsetExamples clears the current DataType definition's
	// example map.
	UnsetExamples() StringType

	// SetDisplayName sets the current DataType definition's
	// displayName value.
	SetDisplayName(string) StringType

	// UnsetDisplayName removes the current DataType
	// definition's displayName value.
	UnsetDisplayName() StringType

	// SetDescription sets the current DataType definition's
	// description value.
	SetDescription(string) StringType

	// UnsetDescription removes the current DataType
	// definition's description value.
	UnsetDescription() StringType

	// SetAnnotations replaces the current DataType
	// definition's annotation map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) StringType

	// UnsetAnnotations clears the current DataType
	// definition's applied annotation map.
	UnsetAnnotations() StringType

	// SetFacetDefinitions replaces the current DataType
	// definition's custom facet definitions with the given
	// value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetFacets.
	SetFacetDefinitions(facets FacetMap) StringType

	// UnsetFacetDefinitions removes all custom facets defined
	// on the current DataType definition.
	UnsetFacetDefinitions() StringType

	// SetXML sets the current DataType definition's xml serialization settings
	// object to the given value.
	//
	// Passing this method a nil value is effectively the same as calling
	// UnsetXML.
	SetXML(XML) StringType

	// UnsetXML removes the xml serialization settings object from the current
	// DataType definition.
	UnsetXML() StringType

	// Enum returns a slice of the enum values assigned to the
	// current DataType definition.
	Enum() []string

	// SetEnum replaces the current DataType definition's enum
	// slice with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetEnum.
	SetEnum([]string) StringType

	// UnsetEnum clears the enum values from the current
	// DataType definition.
	UnsetEnum() StringType

	// SetExtraFacets replaces the facets applied to the
	// current DataType definition with the given values.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExtraFacets.
	SetExtraFacets(facets AnyMap) StringType

	// UnsetExtraFacets removes all extra facets from the
	// current DataType definition.
	UnsetExtraFacets() StringType

	SetRequired(bool) StringType
	// Pattern returns an option which will contain the regex
	// pattern for the current string type definition if it is
	// set.
	Pattern() option.String

	// SetPattern sets the regex pattern for the current
	// string type definition.
	SetPattern(string) StringType

	// UnsetPattern removes the regex pattern facet from the
	// current string type definition.
	UnsetPattern() StringType

	// MinLength returns the minLength facet value set on the
	// current string type definition.
	//
	// Default value is 0.  If this facet is set to its
	// default value when rendered, it will not appear in the
	// output RAML.
	MinLength() uint

	// SetMinLength sets the value for the minLength facet on
	// the current string type definition.
	//
	// Default value is 0.  If this facet is set to its
	// default value when rendered, it will not appear in the
	// output RAML.
	SetMinLength(min uint) StringType

	// MaxLength returns the maxLength facet value set on the
	// current string type definition.
	MaxLength() option.Uint

	// SetMaxLength sets the value for the maxLength facet on
	// the current string type definition.
	SetMaxLength(ln uint) StringType

	UnsetMaxLength() StringType
}
