package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// CustomType generated @ 2023-01-17T10:02:54.294844187-05:00
type CustomType interface {
	Unmarshaler
	Marshaler
	ExtendedDataType

	// SetType sets the parent type name, this does not change
	// the underlying kind of the DataType.
	SetType(string) CustomType

	// Default returns an option of the current default value
	// which will be empty if the value is not set.
	Default() option.Untyped

	// SetDefault sets the default value for the current
	// DataType definition.
	SetDefault(interface{}) CustomType

	// UnsetDefault removes the default value definition from
	// the current DataType definition.
	//
	// If no default was previously set, this method does
	// nothing.
	UnsetDefault() CustomType

	// Example returns the singular example value assigned to
	// the current DataType definition.
	Example() CustomExample

	// SetExample sets the singular example value assigned to
	// the current DataType definition.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExample.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExample(ex CustomExample) CustomType

	// UnsetExample removes the singular example value from
	// the current DataType definition.
	UnsetExample() CustomType

	// Examples returns a mutable map of examples for the
	// current DataType definition.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	Examples() CustomExampleMap

	// SetExamples replaces the current DataType definition's
	// example map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// ass calling UnsetExamples.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExamples(examples CustomExampleMap) CustomType

	// UnsetExamples clears the current DataType definition's
	// example map.
	UnsetExamples() CustomType

	// SetDisplayName sets the current DataType definition's
	// displayName value.
	SetDisplayName(string) CustomType

	// UnsetDisplayName removes the current DataType
	// definition's displayName value.
	UnsetDisplayName() CustomType

	// SetDescription sets the current DataType definition's
	// description value.
	SetDescription(string) CustomType

	// UnsetDescription removes the current DataType
	// definition's description value.
	UnsetDescription() CustomType

	// SetAnnotations replaces the current DataType
	// definition's annotation map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) CustomType

	// UnsetAnnotations clears the current DataType
	// definition's applied annotation map.
	UnsetAnnotations() CustomType

	// SetFacetDefinitions replaces the current DataType
	// definition's custom facet definitions with the given
	// value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetFacets.
	SetFacetDefinitions(facets FacetMap) CustomType

	// UnsetFacetDefinitions removes all custom facets defined
	// on the current DataType definition.
	UnsetFacetDefinitions() CustomType

	// SetXML sets the current DataType definition's xml serialization settings
	// object to the given value.
	//
	// Passing this method a nil value is effectively the same as calling
	// UnsetXML.
	SetXML(XML) CustomType

	// UnsetXML removes the xml serialization settings object from the current
	// DataType definition.
	UnsetXML() CustomType

	// Enum returns a slice of the enum values assigned to the
	// current DataType definition.
	Enum() []interface{}

	// SetEnum replaces the current DataType definition's enum
	// slice with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetEnum.
	SetEnum([]interface{}) CustomType

	// UnsetEnum clears the enum values from the current
	// DataType definition.
	UnsetEnum() CustomType

	// SetExtraFacets replaces the facets applied to the
	// current DataType definition with the given values.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExtraFacets.
	SetExtraFacets(facets AnyMap) CustomType

	// UnsetExtraFacets removes all extra facets from the
	// current DataType definition.
	UnsetExtraFacets() CustomType

	SetRequired(bool) CustomType
}
